package main

import (
	"crypto/rand"
	"crypto/sha256"
	"go-proton/atomic/swap/bitcoin"
	"go-proton/core/accounts"
	"log"

	"github.com/mr-tron/base58"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	// Select network
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:19011",
		User:         "admin2",
		Pass:         "123",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	priv, err := base58.Decode("GZKhz5NBokGgGGGcVox9HfBiukXvQipN68UKBuWPfi1i")
	if err != nil {
		panic(err)
	}
	println(priv)

	account, err := accounts.FromPrivateKey(priv, accounts.Bitcoin)
	if err != nil {
		panic(err)
	}
	println(account.Address)

	//	0x33729bd991affb47080bde938a6f4cdfcdd713c9
	btc, _ := bitcoin.New(connCfg, priv)
	hasher := sha256.New()
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		panic(err)
	}
	hasher.Write(secret)
	hash := hasher.Sum(nil)
	txHash, err := btc.Send(account.PublicKey, 0.1, hash, 10)
	if err != nil {
		panic(err)
	}
	info, err := btc.ExtractSwap(txHash)
	if err != nil {
		panic(err)
	}
	println("amount: ", info.Amount)
	println("block: ", info.SwapExpire)
	println("hash: ", txHash)
	rv, err := btc.Return(hash)
	if err != nil {
		panic(err)
	}
	print(rv)
}
