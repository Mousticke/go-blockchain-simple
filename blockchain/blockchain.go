package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	DbPath = "./tmp/blocks"
	DbFile = "blockchain.db"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func NewBlockchain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions(DbPath)

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found. Creating a new one.")
			genesis := Genesis()
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.ValueCopy(nil)

			return err
		}
	})

	Handle(err)
	bc := Blockchain{
		LastHash: lastHash,
		Database: db,
	}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.ValueCopy(nil)

		return err
	})
	newBlock := NewBlock(data, lastHash)
	err = bc.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		bc.LastHash = newBlock.Hash

		return err
	})

	Handle(err)
	fmt.Printf("Added Block %x\n", newBlock.Hash)

}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	it := &BlockchainIterator{bc.LastHash, bc.Database}
	return it
}

func (it *BlockchainIterator) Next() *Block {
	var block *Block

	err := it.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(it.CurrentHash)
		Handle(err)
		val, err := item.ValueCopy(nil)
		block = Deserialize(val)

		return err
	})

	Handle(err)

	it.CurrentHash = block.PrevHash

	return block
}
