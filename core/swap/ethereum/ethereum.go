package ethereum

import (
	"context"
	"crypto/ecdsa"
	"go-proton/atomic/swap"
	"go-proton/utils/crypto"
	"go-proton/utils/slice"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	weiInEth = big.NewFloat(math.Pow10(18))
)

type Exchanger struct {
	contract *HashTimeLock
	client   *ethclient.Client
	ctx      context.Context
	privKey  *ecdsa.PrivateKey
}

func New(ctx context.Context, host string, contractAddress string, privBytes []byte) (*Exchanger, error) {
	client, err := ethclient.DialContext(ctx, host)
	if err != nil {
		return nil, err
	}

	secp256k1Manager := crypto.Secp256k1Manager{}
	address := common.HexToAddress(contractAddress)
	contract, err := NewHashTimeLock(address, client)
	if err != nil {
		return nil, err
	}

	return &Exchanger{
		contract: contract,
		client:   client,
		ctx:      ctx,
		privKey:  secp256k1Manager.GetEcdsaPrivateKey(privBytes),
	}, nil
}

func (ex *Exchanger) Send(recipient string, amount float64, hash []byte, blockTimeout int64) ([]byte, error) {
	lastBlock, err := ex.client.HeaderByNumber(ex.ctx, nil)
	if err != nil {
		return nil, err
	}

	transactOpts := bind.NewKeyedTransactor(ex.privKey)

	ethAmount := new(big.Float).SetFloat64(amount)
	weiAmountFloat := new(big.Float).Mul(ethAmount, weiInEth)
	weiAmountInt := new(big.Int)
	transactOpts.Value, _ = weiAmountFloat.Int(weiAmountInt)

	expireHeight := lastBlock.Number.Add(lastBlock.Number, big.NewInt(blockTimeout))
	tx, err := ex.contract.Lock(transactOpts, common.HexToAddress(recipient), slice.ConvertTo32(hash), expireHeight)
	if err != nil {
		return nil, err
	}

	return tx.Hash().Bytes(), err
}

func (ex *Exchanger) Receive(hash []byte, secret []byte) ([]byte, error) {
	tx, err := ex.contract.Unlock(bind.NewKeyedTransactor(ex.privKey), slice.ConvertTo32(hash), secret)
	if err != nil {
		return nil, err
	}
	return tx.Hash().Bytes(), err
}

func (ex *Exchanger) Return(hash []byte) ([]byte, error) {
	tx, err := ex.contract.ReturnToSender(bind.NewKeyedTransactor(ex.privKey), slice.ConvertTo32(hash))
	if err != nil {
		return nil, err
	}
	return tx.Hash().Bytes(), err
}

func (ex *Exchanger) ExtractSwap(hash []byte) (swap.AtomicSwapInfo, error) {
	swapRequest, err := ex.contract.SwapRequests(&bind.CallOpts{
		Context: ex.ctx,
		Pending: false,
	}, slice.ConvertTo32(hash))
	if err != nil {
		return swap.AtomicSwapInfo{}, err
	}

	ethAmount := new(big.Float).Quo(new(big.Float).SetInt(swapRequest.Amount), weiInEth)
	amount, _ := ethAmount.Float64()
	return swap.AtomicSwapInfo{
		SwapExpire: swapRequest.ExpireHeight.Int64(),
		Amount:     amount,
		SecretHash: swapRequest.SecretHash,
		Sender:     swapRequest.Sender.Bytes(),
		Recipient:  swapRequest.Recipient.Bytes(),
	}, nil
}

func (ex *Exchanger) ExtractSecret(hash []byte) ([]byte, error) {
	swapRequest, err := ex.contract.SwapRequests(&bind.CallOpts{
		Context: ex.ctx,
		Pending: false,
	}, slice.ConvertTo32(hash))
	if err != nil {
		return nil, err
	}
	return swapRequest.Secret, nil
}
