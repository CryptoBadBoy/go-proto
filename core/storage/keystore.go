package storage

import (
	"go-proton/core/accounts"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	path = "db/keystore"
)

func PutKeystore(bType accounts.BlockchainType, address string, priv []byte) error {
	db, err := leveldb.OpenFile(path+"/"+bType.String(), nil)
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Put([]byte(address), priv, nil); err != nil {
		return err
	}

	return nil
}

func Keystore(bType accounts.BlockchainType) ([][]byte, error) {
	accounts := make([][]byte, 0)
	db, err := leveldb.OpenFile(path+"/"+bType.String(), nil)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		value := make([]byte, len(iter.Value()))
		copy(value, iter.Value())
		accounts = append(accounts, value)
	}
	iter.Release()

	if err = iter.Error(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func DeleteKeystore(bType accounts.BlockchainType, address string) error {
	db, err := leveldb.OpenFile(path+"/"+bType.String(), nil)
	defer db.Close()

	if err != nil {
		return err
	} else if err := db.Delete([]byte(address), nil); err != nil {
		return err
	}

	return nil
}

func GetKeystore(bType accounts.BlockchainType, address string) ([]byte, error) {
	db, err := leveldb.OpenFile(path+"/"+bType.String(), nil)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	return db.Get([]byte(address), nil)
}
