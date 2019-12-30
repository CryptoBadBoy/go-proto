package bitcoin

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"go-proton/atomic/swap"
	"go-proton/constants"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

func spendP2Tx(client *rpcclient.Client, amount float64, changeScript []byte, payoutScript []byte, privKey *btcec.PrivateKey) (*wire.MsgTx, error) {
	unspendTxs, err := client.ListUnspent()
	if err != nil {
		return nil, err
	}

	p2kAddress, err := btcutil.NewAddressPubKey(crypto.FromECDSAPub(&privKey.PublicKey), constants.BtcChainParams)
	if err != nil {
		return nil, err
	}

	satoshiAmount, err := btcutil.NewAmount(amount)
	if err != nil {
		return nil, err
	}

	satoshiFeeB, err := btcutil.NewAmount(estimateSmartFee(client) / 1024)
	if err != nil {
		return nil, err
	}

	address := p2kAddress.EncodeAddress()
	var inTxs []*wire.TxIn
	var totalInputAmount float64
	var scripts [][]byte
	for _, v := range unspendTxs {
		if v.Address != address {
			continue
		}

		byteScript, err := hex.DecodeString(v.ScriptPubKey)
		if err != nil {
			return nil, err
		}

		scriptType := txscript.GetScriptClass(byteScript)
		if scriptType != txscript.PubKeyTy && scriptType != txscript.PubKeyHashTy {
			continue
		}

		txHash, err := chainhash.NewHashFromStr(v.TxID)
		if err != nil {
			return nil, err
		}
		totalInputAmount += v.Amount

		inTxs = append(inTxs, wire.NewTxIn(wire.NewOutPoint(txHash, v.Vout), nil, nil))
		scripts = append(scripts, byteScript)

		if totalInputAmount > amount {
			var outTxs []*wire.TxOut
			outTxs = append(outTxs, wire.NewTxOut(int64(satoshiAmount), payoutScript))
			freeOutWithoutFee := totalInputAmount - amount

			if freeOutWithoutFee > 0 {
				satoshiFreeOut, err := btcutil.NewAmount(freeOutWithoutFee)
				if err != nil {
					return nil, err
				}
				outTxs = append(outTxs, wire.NewTxOut(int64(satoshiFreeOut), changeScript))
			} else if freeOutWithoutFee <= 0 {
				continue
			}
			tx, err := formatTx(outTxs, inTxs, scripts, privKey)
			fee := int64(tx.SerializeSize()) * int64(satoshiFeeB)

			if outTxs[1].Value < fee {
				continue
			} else if outTxs[1].Value == fee {
				outTxs = outTxs[:0]
			} else {
				outTxs[1].Value -= fee
			}

			tx, err = formatTx(outTxs, inTxs, scripts, privKey)
			if err != nil {
				return nil, err
			}

			return tx, nil
		}
	}

	if totalInputAmount <= amount {
		return nil, errors.New("insufficient funds")
	}

	return nil, errors.New("unknown error")
}

func spendHTLCTx(client *rpcclient.Client, hash []byte, payOutScript []byte, privKey *btcec.PrivateKey, endScript []byte, timeoutEnded bool) (*wire.MsgTx, error) {
	txHash, err := chainhash.NewHash(hash)
	if err != nil {
		return nil, err
	}

	satoshiFeeB, err := btcutil.NewAmount(estimateSmartFee(client) / 1024)
	if err != nil {
		return nil, err
	}

	htlcTx, err := client.GetRawTransaction(txHash)
	if err != nil {
		return nil, err
	}

	var htlcOut *wire.TxOut
	var atomicData *txscript.AtomicSwapDataPushes
	inputIndex := 0
	for i, tx := range htlcTx.MsgTx().TxOut {
		value, err := txscript.ExtractAtomicSwapDataPushes(0, tx.PkScript)
		if err != nil {
			return nil, err
		}
		if value == nil {
			continue
		}

		inputIndex = i
		atomicData = value
		htlcOut = tx
		break
	}

	unlockTx := wire.NewMsgTx(2)
	// format tx
	if timeoutEnded {
		unlockTx.LockTime = uint32(atomicData.LockTime)
	}
	unlockTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txHash, uint32(inputIndex)), nil, nil))
	unlockTx.TxIn[0].Sequence = 0
	unlockTx.AddTxOut(wire.NewTxOut(htlcOut.Value, payOutScript))

	// sign tx
	// TODO add normal commission calculation
	for i := 0; i < 2; i++ {
		sign, err := txscript.RawTxInSignature(unlockTx, inputIndex, htlcOut.PkScript, txscript.SigHashAll, privKey)
		if err != nil {
			return nil, err
		}
		scriptBuilder := txscript.NewScriptBuilder()
		scriptBuilder.AddData(sign)
		scriptBuilder.AddData(crypto.FromECDSAPub(&privKey.PublicKey))
		scriptBuilder.AddOps(endScript)
		script, err := scriptBuilder.Script()
		if err != nil {
			return nil, err
		}

		unlockTx.TxIn[0].SignatureScript = script
		if unlockTx.TxOut[0].Value == htlcOut.Value {
			unlockTx.TxOut[0].Value -= int64(unlockTx.SerializeSize()) * int64(satoshiFeeB)
		}
	}
	return unlockTx, nil
}

func formatTx(outTxs []*wire.TxOut, inTxs []*wire.TxIn, scripts [][]byte, privKey *btcec.PrivateKey) (*wire.MsgTx, error) {
	unsignedContract := wire.NewMsgTx(2)
	unsignedContract.TxOut = outTxs
	unsignedContract.TxIn = inTxs
	for i, v := range unsignedContract.TxIn {
		scriptType := txscript.GetScriptClass(scripts[i])
		sign, err := txscript.RawTxInSignature(unsignedContract, i, scripts[i], txscript.SigHashAll, privKey)
		scriptBuilder := txscript.NewScriptBuilder()
		switch scriptType {
		case txscript.PubKeyTy:
			scriptBuilder.AddData(sign)
		case txscript.PubKeyHashTy:
			scriptBuilder.AddData(sign)
			scriptBuilder.AddData(crypto.FromECDSAPub(&privKey.PublicKey))
		}

		script, err := scriptBuilder.Script()
		if err != nil {
			return nil, err
		}
		v.SignatureScript = script
	}
	return unsignedContract, nil
}

func atomicSwapContract(sender *[ripemd160.Size]byte, recipient []byte, heightTimeout int64, hash []byte) ([]byte, error) {
	b := txscript.NewScriptBuilder()

	b.AddOp(txscript.OP_IF)
	{
		b.AddOp(txscript.OP_SIZE)
		b.AddInt64(swap.SecretSize)
		b.AddOp(txscript.OP_EQUALVERIFY)
		b.AddOp(txscript.OP_SHA256)
		b.AddData(hash)
		b.AddOp(txscript.OP_EQUALVERIFY)
		b.AddOp(txscript.OP_DUP)
		b.AddOp(txscript.OP_HASH160)
		b.AddData(recipient)
	}
	b.AddOp(txscript.OP_ELSE)
	{
		b.AddInt64(heightTimeout)
		b.AddOp(txscript.OP_CHECKLOCKTIMEVERIFY)
		b.AddOp(txscript.OP_DROP)
		b.AddOp(txscript.OP_DUP)
		b.AddOp(txscript.OP_HASH160)
		b.AddData(sender[:])
	}
	b.AddOp(txscript.OP_ENDIF)
	b.AddOp(txscript.OP_EQUALVERIFY)
	b.AddOp(txscript.OP_CHECKSIG)

	return b.Script()
}

func estimateSmartFee(client *rpcclient.Client) float64 {
	params := []json.RawMessage{[]byte("6")}
	estimateRawResp, err := client.RawRequest("estimatesmartfee", params)
	if err != nil {
		return defaultFeePerKb
	}
	var estimateResp struct {
		FeeRate float64  `json:"feerate"`
		Errors  []string `json:"errors"`
	}
	err = json.Unmarshal(estimateRawResp, &estimateResp)
	if err != nil {
		return defaultFeePerKb
	}
	if estimateResp.Errors != nil {
		return defaultFeePerKb
	}
	return estimateResp.FeeRate
}
