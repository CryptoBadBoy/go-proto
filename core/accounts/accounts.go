package accounts

import (
	"go-proton/constants"
	"go-proton/utils/crypto"
)

type BlockchainType int

const (
	None BlockchainType = iota
	Bitcoin
	Ethereum
)

type Account struct {
	Type BlockchainType
	crypto.KeyPair
	Address string
}

type AddressFormatter interface {
	Address(pubKey []byte) (string, error)
	Type() BlockchainType
}

func (bType BlockchainType) String() string {
	switch bType {
	case Bitcoin:
		return "bitcoin"
	case Ethereum:
		return "ethereum"
	}
	return ""
}

func ConvertToType(str string) BlockchainType {
	switch str {
	case "bitcoin":
		return Bitcoin
	case "ethereum":
		return Ethereum
	}
	return None
}

func ECDSAByType(blockchainType BlockchainType) crypto.ECDSAManager {
	var manager crypto.ECDSAManager
	switch blockchainType {
	case Bitcoin:
		manager = crypto.Secp256k1Manager{}
	case Ethereum:
		manager = crypto.Secp256k1Manager{}
	}
	return manager
}

func FormatterByType(blockchainType BlockchainType) AddressFormatter {
	var formatter AddressFormatter
	switch blockchainType {
	case Bitcoin:
		formatter = BtcFormatter{
			ChainCfg: constants.BtcChainParams,
		}
	case Ethereum:
		formatter = EthFormatter{}
	}
	return formatter
}

func Generate(blockchainType BlockchainType) (Account, error) {
	formatter := FormatterByType(blockchainType)
	manager := ECDSAByType(blockchainType)
	keyPair, err := manager.Generate()
	if err != nil {
		return Account{}, err
	}
	address, err := formatter.Address(keyPair.PublicKey)
	if err != nil {
		return Account{}, err
	}
	return Account{
		Type:    formatter.Type(),
		Address: address,
		KeyPair: keyPair,
	}, nil
}

func FromPrivateKey(privateKey []byte, blockchainType BlockchainType) (Account, error) {
	formatter := FormatterByType(blockchainType)
	manager := ECDSAByType(blockchainType)
	publicKey := manager.GetPublicKey(privateKey)
	address, err := formatter.Address(publicKey)
	if err != nil {
		return Account{}, err
	}
	return Account{
		Type:    formatter.Type(),
		Address: address,
		KeyPair: crypto.KeyPair{
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		},
	}, nil
}
