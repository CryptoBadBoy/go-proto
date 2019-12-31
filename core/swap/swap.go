package swap

const (
	SecretSize = 32
)

type Exchanger interface {
	Send(recipient string, amount float64, hash []byte, blockTimeout int64) ([]byte, error)
	Receive(hash []byte, secret []byte) ([]byte, error)
	Return(hash []byte) ([]byte, error)
	ExtractSwap(hash []byte) (AtomicSwapInfo, error)
	ExtractSecret(hash []byte) ([]byte, error)
}

type AtomicSwapInfo struct {
	SwapExpire int64
	Amount     float64
	SecretHash [32]byte
	Sender     []byte
	Recipient  []byte
}
