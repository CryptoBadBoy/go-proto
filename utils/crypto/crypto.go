package crypto

type ECDSAType int

const (
	Secp256k1 ECDSAType = iota
)

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}
type ECDSAManager interface {
	Generate() (KeyPair, error)
	GetPublicKey(priv []byte) []byte
	//Sign
	//Verify
}
