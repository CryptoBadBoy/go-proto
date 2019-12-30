package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type Secp256k1Manager struct{}

func (Secp256k1Manager) Generate() (KeyPair, error) {
	priv, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return KeyPair{}, err
	}
	return KeyPair{
		PrivateKey: priv.D.Bytes(),
		PublicKey:  crypto.FromECDSAPub(&priv.PublicKey),
	}, nil
}

func (secp Secp256k1Manager) GetPublicKey(privBytes []byte) []byte {
	privateKey := secp.GetEcdsaPrivateKey(privBytes)
	return crypto.FromECDSAPub(&privateKey.PublicKey)
}

func (Secp256k1Manager) GetEcdsaPrivateKey(privBytes []byte) *ecdsa.PrivateKey {
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
		},
		D: new(big.Int),
	}
	privateKey.D.SetBytes(privBytes)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privBytes)
	return privateKey
}

/*


func newAccountSecp256k1(k []byte) *ecdsa.PrivateKey {
	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
		},
		D: new(big.Int),
	}
	priv.D.SetBytes(k)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(k)
	return priv
}

*/
