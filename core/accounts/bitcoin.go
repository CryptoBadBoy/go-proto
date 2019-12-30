package accounts

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type BtcFormatter struct {
	ChainCfg *chaincfg.Params
}

func (formatter BtcFormatter) Address(pubKey []byte) (string, error) {
	address, err := btcutil.NewAddressPubKey(pubKey, formatter.ChainCfg)
	if err != nil {
		return "", err
	}
	return address.EncodeAddress(), nil
}

func (formatter BtcFormatter) Type() BlockchainType {
	return Bitcoin
}
