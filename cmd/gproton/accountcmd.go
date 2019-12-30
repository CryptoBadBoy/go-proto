package main

import (
	"fmt"
	"go-proton/core/storage"

	"github.com/mr-tron/base58"

	"go-proton/core/accounts"

	"gopkg.in/urfave/cli.v1"
)

var (
	accountCommand = cli.Command{
		Name:        "account",
		Usage:       "Blockchain accounts",
		Description: "",
		Subcommands: []cli.Command{
			{
				Name:        "new",
				Usage:       "Generate new account",
				Aliases:     []string{"generate", "create"},
				Description: "",
				Action:      accountCreate,
				ArgsUsage:   "<blockchain type>",
			},
			{
				Name:        "list",
				Usage:       "All accounts",
				Aliases:     []string{"all"},
				Description: "",
				Action:      listAccounts,
				ArgsUsage:   "<blockchain type>",
			},
			{
				Name:        "get",
				Usage:       "Get account info",
				Aliases:     []string{"info"},
				Description: "",
				Action:      getAccount,
				ArgsUsage:   "<blockchain type> <address>",
			},
			/*{
				Name:        "import",
				Usage:       "Imoprt account",
				Description: "",
				Action:      ,
				ArgsUsage:   "<blockchain type>",
			},*/
			{
				Name:        "delete",
				Usage:       "Delete account",
				Aliases:     []string{"remove", "drop"},
				Description: "",
				Action:      deleteAccount,
				ArgsUsage:   "<blockchain type> <address>",
			},
		},
	}
)

func accountCreate(ctx *cli.Context) error {
	blockchainType := accounts.ConvertToType(ctx.Args().First())

	account, err := accounts.Generate(blockchainType)
	if err != nil {
		return err
	}

	if err := storage.PutKeystore(blockchainType, account.Address, account.PrivateKey); err != nil {
		return err
	}

	fmt.Printf("New account generated:\n")
	fmt.Printf("Public Key: %v\n", base58.Encode(account.PublicKey))
	fmt.Printf("Private Key: %v\n", base58.Encode(account.PrivateKey))
	fmt.Printf("Address: %v\n", account.Address)

	return nil
}

func listAccounts(ctx *cli.Context) error {
	blockchainType := accounts.ConvertToType(ctx.Args().First())
	privKeys, err := storage.Keystore(blockchainType)
	if err != nil {
		return err
	}

	fmt.Println("All accounts:")
	for _, privKey := range privKeys {
		account, err := accounts.FromPrivateKey(privKey, blockchainType)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", account.Address)
	}

	return nil
}

func deleteAccount(ctx *cli.Context) error {
	if err := storage.DeleteKeystore(accounts.ConvertToType(ctx.Args().First()), ctx.Args().Get(1)); err != nil {
		return err
	}

	fmt.Println("Success!")

	return nil
}

func getAccount(ctx *cli.Context) error {
	blockchainType := accounts.ConvertToType(ctx.Args().First())

	priv, err := storage.GetKeystore(blockchainType, ctx.Args().Get(1))
	if err != nil {
		return err
	}
	println(priv)

	account, err := accounts.FromPrivateKey(priv, blockchainType)
	if err != nil {
		return err
	}
	fmt.Printf("New account generated:\n")
	fmt.Printf("Public Key: %v\n", base58.Encode(account.PublicKey))
	fmt.Printf("Private Key: %v\n", base58.Encode(account.PrivateKey))
	fmt.Printf("Address: %v\n", account.Address)

	return nil
}
