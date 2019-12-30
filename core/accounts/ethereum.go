package accounts

import (
	"encoding/hex"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

type EthFormatter struct{}

func (EthFormatter) Address(pubKey []byte) (string, error) {
	return "0x" + hex.EncodeToString(ethCrypto.Keccak256(pubKey[1:])[12:]), nil
}

func (formatter EthFormatter) Type() BlockchainType {
	return Ethereum
}
