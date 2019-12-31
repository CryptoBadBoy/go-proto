package commands

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-proton/cmd/utils"
	"go-proton/core/accounts"
	"go-proton/core/storage"
	"go-proton/core/swap"
	"go-proton/core/swap/bitcoin"
	"go-proton/core/swap/ethereum"
	"strconv"

	"github.com/btcsuite/btcd/rpcclient"

	"gopkg.in/urfave/cli.v1"
)

const (
	defaultTimeoutBitcoin  = 20
	defaultTimeoutEthereum = 2300
)

var (
	cfg         utils.Config
	SwapCommand = cli.Command{
		Name:        "swap",
		Usage:       "Atomic swap",
		Description: "",
		Subcommands: []cli.Command{
			{
				Name:        "sell",
				Usage:       "sell token",
				Description: "",
				Action:      sell,
				ArgsUsage:   "<tokenSell> <seller> <buyer> <amount>",
			},
			{
				Name:        "buy",
				Usage:       "buy token",
				Description: "",
				Action:      buy,
				ArgsUsage:   "<hash> <tokenSell> <tokenBuy> <buyer> <seller> <amount>",
			},
			{
				Name:        "unlock",
				Usage:       "unlock swap",
				Description: "",
				Action:      unlock,
				ArgsUsage:   "<token> <hash> <seller/buyer> <secret>",
			},
			{
				Name:        "return",
				Usage:       "return swap",
				Description: "",
				Action:      returnAmount,
				ArgsUsage:   "<token> <hash> <seller/buyer>",
			},
		},
	}
)

func SetConfig(newCfg utils.Config) {
	cfg = newCfg
}

func sell(ctx *cli.Context) error {
	tokenSell := accounts.ConvertToType(ctx.Args().First())
	seller := ctx.Args()[1]
	buyer := ctx.Args()[2]
	amount, err := strconv.ParseFloat(ctx.Args()[3], 64)
	if err != nil {
		panic(err)
	}

	hasher := sha256.New()
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		panic(err)
	}
	hasher.Write(secret)
	hash := hasher.Sum(nil)

	priv, err := storage.GetKeystore(tokenSell, seller)
	if err != nil {
		panic(err)
	}

	var exchanger swap.Exchanger
	var timeout int64

	switch tokenSell {
	case accounts.Bitcoin:
		timeout = defaultTimeoutBitcoin
		exchanger, err = bitcoin.New(&rpcclient.ConnConfig{
			Host:         cfg.Bitcoin.Host,
			User:         cfg.Bitcoin.User,
			Pass:         cfg.Bitcoin.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, priv)
		if err != nil {
			panic(err)
		}
	case accounts.Ethereum:
		timeout = defaultTimeoutEthereum
		exchanger, err = ethereum.New(context.Background(), cfg.Ethereum.Host, cfg.Ethereum.ContractAddress, priv)
		if err != nil {
			panic(err)
		}
	}
	txHash, err := exchanger.Send(buyer, amount, hash, timeout)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Swap sell info:\n")
	fmt.Printf("Hash: %v\n", hex.EncodeToString(txHash))
	fmt.Printf("Secret: %v\n", hex.EncodeToString(secret))
	fmt.Printf("SecretHash: %v\n", hex.EncodeToString(hash))
	return nil
}

func buy(ctx *cli.Context) error {
	hashString := ctx.Args()[0]
	tokenSell := accounts.ConvertToType(ctx.Args()[1])
	tokenBuy := accounts.ConvertToType(ctx.Args()[2])
	buyer := ctx.Args()[3]
	seller := ctx.Args()[4]
	amount, err := strconv.ParseFloat(ctx.Args()[5], 64)
	if err != nil {
		panic(err)
	}

	hash, err := hex.DecodeString(hashString)
	if err != nil {
		panic(err)
	}

	priv, err := storage.GetKeystore(tokenBuy, buyer)
	if err != nil {
		panic(err)
	}

	var sellerExchanger swap.Exchanger
	var buyerExchanger swap.Exchanger
	var timeout int64
	switch tokenSell {
	case accounts.Bitcoin:
		timeout = defaultTimeoutEthereum * 2
		sellerExchanger, err = bitcoin.New(&rpcclient.ConnConfig{
			Host:         cfg.Bitcoin.Host,
			User:         cfg.Bitcoin.User,
			Pass:         cfg.Bitcoin.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, priv)
		if err != nil {
			panic(err)
		}
		buyerExchanger, err = ethereum.New(context.Background(), cfg.Ethereum.Host, cfg.Ethereum.ContractAddress, priv)
		if err != nil {
			panic(err)
		}
	case accounts.Ethereum:
		timeout = defaultTimeoutBitcoin * 2
		sellerExchanger, err = ethereum.New(context.Background(), cfg.Ethereum.Host, cfg.Ethereum.ContractAddress, priv)
		if err != nil {
			panic(err)
		}
		buyerExchanger, err = bitcoin.New(&rpcclient.ConnConfig{
			Host:         cfg.Bitcoin.Host,
			User:         cfg.Bitcoin.User,
			Pass:         cfg.Bitcoin.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, priv)
		if err != nil {
			panic(err)
		}
	}

	atomicInfo, err := sellerExchanger.ExtractSwap(hash)
	if err != nil {
		panic(err)
	}

	txHash, err := buyerExchanger.Send(seller, amount, atomicInfo.SecretHash[:], timeout)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Swap buy info:\n")
	fmt.Printf("Hash: %v\n", hex.EncodeToString(txHash))
	fmt.Printf("SecretHash: %v\n", hex.EncodeToString(hash))
	return nil
}

func unlock(ctx *cli.Context) error {
	token := accounts.ConvertToType(ctx.Args()[0])
	hash, err := hex.DecodeString(ctx.Args()[1])
	if err != nil {
		panic(err)
	}
	owner := ctx.Args()[2]
	secret, err := hex.DecodeString(ctx.Args()[3])
	if err != nil {
		panic(err)
	}

	priv, err := storage.GetKeystore(token, owner)
	if err != nil {
		panic(err)
	}

	var exchanger swap.Exchanger
	switch token {
	case accounts.Bitcoin:
		exchanger, err = bitcoin.New(&rpcclient.ConnConfig{
			Host:         cfg.Bitcoin.Host,
			User:         cfg.Bitcoin.User,
			Pass:         cfg.Bitcoin.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, priv)
		if err != nil {
			panic(err)
		}
	case accounts.Ethereum:
		exchanger, err = ethereum.New(context.Background(), cfg.Ethereum.Host, cfg.Ethereum.ContractAddress, priv)
		if err != nil {
			panic(err)
		}
	}

	txHash, err := exchanger.Receive(hash, secret)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Swap unlock info:\n")
	fmt.Printf("Unlock hash: %v\n", hex.EncodeToString(txHash))
	return nil
}

func returnAmount(ctx *cli.Context) error {
	token := accounts.ConvertToType(ctx.Args()[0])
	hash, err := hex.DecodeString(ctx.Args()[1])
	if err != nil {
		panic(err)
	}
	owner := ctx.Args()[2]

	priv, err := storage.GetKeystore(token, owner)
	if err != nil {
		panic(err)
	}

	var exchanger swap.Exchanger
	switch token {
	case accounts.Bitcoin:
		exchanger, err = bitcoin.New(&rpcclient.ConnConfig{
			Host:         cfg.Bitcoin.Host,
			User:         cfg.Bitcoin.User,
			Pass:         cfg.Bitcoin.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, priv)
		if err != nil {
			panic(err)
		}
	case accounts.Ethereum:
		exchanger, err = ethereum.New(context.Background(), cfg.Ethereum.Host, cfg.Ethereum.ContractAddress, priv)
		if err != nil {
			panic(err)
		}
	}

	txHash, err := exchanger.Return(hash)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Swap return info:\n")
	fmt.Printf("Return hash: %v\n", hex.EncodeToString(txHash))
	return nil
}
