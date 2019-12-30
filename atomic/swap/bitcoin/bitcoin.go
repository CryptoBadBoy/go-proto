package bitcoin

import (
	"errors"
	"go-proton/atomic/swap"
	"go-proton/constants"
	"go-proton/core/accounts"

	"github.com/mr-tron/base58/base58"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"github.com/btcsuite/btcd/btcec"

	"github.com/btcsuite/btcutil"

	"github.com/btcsuite/btcd/txscript"

	"github.com/btcsuite/btcd/rpcclient"
)

const (
	defaultFeePerKb = 0.0001
)

type Exchanger struct {
	*btcutil.AddressPubKeyHash
	*rpcclient.ConnConfig
	*btcec.PrivateKey
}

func New(conf *rpcclient.ConnConfig, privBytes []byte) (*Exchanger, error) {
	account, err := accounts.FromPrivateKey(privBytes, accounts.Bitcoin)
	if err != nil {
		return nil, err
	}

	cp2Addr, err := btcutil.NewAddressPubKey(account.PublicKey, constants.BtcChainParams)
	if err != nil {
		return nil, err
	}
	cp2AddrP2PKH := cp2Addr.AddressPubKeyHash()

	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)
	return &Exchanger{
		AddressPubKeyHash: cp2AddrP2PKH,
		ConnConfig:        conf,
		PrivateKey:        priv,
	}, nil
}

func (ex *Exchanger) Send(recipient string, amount float64, hash []byte, blockTimeout int64) ([]byte, error) {
	client, err := rpcclient.New(ex.ConnConfig, nil)
	if err != nil {
		return nil, err
	}
	defer client.Shutdown()

	block, err := client.GetBlockCount()
	if err != nil {
		return nil, err
	}

	swapExpire := block + blockTimeout

	recipientAddress, err := base58.Decode(recipient)
	if err != nil {
		return nil, err
	}

	contract, err := atomicSwapContract(ex.AddressPubKeyHash.Hash160(), recipientAddress, swapExpire, hash)
	if err != nil {
		return nil, err
	}

	p2hScript, err := txscript.PayToAddrScript(ex.AddressPubKeyHash)
	if err != nil {
		return nil, err
	}

	htlcTx, err := spendP2Tx(client, amount, p2hScript, contract, ex.PrivateKey)
	if err != nil {
		return nil, err
	}

	signedTx, err := client.SendRawTransaction(htlcTx, true)
	if err != nil {
		return nil, err
	}
	return signedTx.CloneBytes(), nil
}

func (ex *Exchanger) Receive(hash []byte, secret []byte) ([]byte, error) {
	return ex.SpendHTLC(hash, false, secret)
}

func (ex *Exchanger) Return(hash []byte) ([]byte, error) {
	return ex.SpendHTLC(hash, true, nil)
}

func (ex *Exchanger) ExtractSwap(hash []byte) (swap.AtomicSwapInfo, error) {
	client, err := rpcclient.New(ex.ConnConfig, nil)
	if err != nil {
		return swap.AtomicSwapInfo{}, err
	}
	defer client.Shutdown()

	txHash, err := chainhash.NewHash(hash)
	if err != nil {
		return swap.AtomicSwapInfo{}, err
	}

	htlcTx, err := client.GetRawTransaction(txHash)
	if err != nil {
		return swap.AtomicSwapInfo{}, err
	}

	var atomicData *txscript.AtomicSwapDataPushes
	var amount int64
	for _, tx := range htlcTx.MsgTx().TxOut {
		value, err := txscript.ExtractAtomicSwapDataPushes(0, tx.PkScript)
		if err != nil {
			return swap.AtomicSwapInfo{}, err
		}
		if value == nil {
			continue
		}
		amount = tx.Value
		atomicData = value
		break
	}
	satoshiAmount := btcutil.Amount(amount)

	return swap.AtomicSwapInfo{
		SwapExpire: atomicData.LockTime,
		Amount:     satoshiAmount.ToBTC(),
		SecretHash: atomicData.SecretHash,
		Recipient:  atomicData.RecipientHash160[:],
		Sender:     atomicData.RefundHash160[:],
	}, nil
}

func (ex *Exchanger) ExtractSecret(hash []byte) ([]byte, error) {
	//TODO: scan blocks for htlc out
	return nil, errors.New("no implementation")
}

func (ex *Exchanger) SpendHTLC(hash []byte, timeoutEnded bool, secret []byte) ([]byte, error) {
	client, err := rpcclient.New(ex.ConnConfig, nil)
	if err != nil {
		return nil, err
	}
	defer client.Shutdown()

	p2hScript, err := txscript.PayToAddrScript(ex.AddressPubKeyHash)
	if err != nil {
		return nil, err
	}
	scriptBuilder := txscript.NewScriptBuilder()
	if timeoutEnded {
		scriptBuilder.AddOp(txscript.OP_FALSE)
	} else {
		scriptBuilder.AddData(secret)
		scriptBuilder.AddOp(txscript.OP_TRUE)
	}

	endScript, err := scriptBuilder.Script()
	if err != nil {
		return nil, err
	}

	unlockTx, err := spendHTLCTx(client, hash, p2hScript, ex.PrivateKey, endScript, timeoutEnded)
	if err != nil {
		return nil, err
	}

	signedTx, err := client.SendRawTransaction(unlockTx, true)
	if err != nil {
		return nil, err
	}
	return signedTx.CloneBytes(), nil
}
