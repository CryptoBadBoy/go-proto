package main

import (
	"crypto/rand"
	"crypto/sha256"
	"go-proton/atomic/swap"
	"go-proton/atomic/swap/bitcoin"
	"go-proton/core/accounts"
	"go-proton/core/storage"
	"strconv"

	"gopkg.in/urfave/cli.v1"
)

const (
	defaultTimeoutBitcoin  = 20
	defaultTimeoutEthereum = 2400
)

var (
	swapCommand = cli.Command{
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
				Aliases:     []string{"all"},
				Description: "",
				Action:      listAccounts,
				ArgsUsage:   "<hash> <tokenSell> <tokenBuy> <buyer> <seller> <amount>",
			},
			{
				Name:        "unlock",
				Usage:       "unlock swap",
				Description: "",
				Action:      listAccounts,
				ArgsUsage:   "<token> <hash> <secret>",
			},
		},
	}
)

func sell(ctx *cli.Context) error {
	tokenSell := accounts.ConvertToType(ctx.Args().First())
	seller := ctx.Args()[1]
	buyer := ctx.Args()[2]
	amount, err := strconv.ParseFloat(ctx.Args()[3], 64)
	if err != nil {
		return err
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
		return err
	}

	var exchanger swap.Exchanger
	var timeout int64

	switch tokenSell {
	case accounts.Bitcoin:
		timeout = defaultTimeoutBitcoin
		exchanger, err = bitcoin.New(nil, priv)
		if err != nil {
			return err
		}
	case accounts.Ethereum:
		timeout = defaultTimeoutEthereum
	}

	exchanger.Send(buyer, amount, hash, timeout)
	return nil
}