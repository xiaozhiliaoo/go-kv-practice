package main

import (
	"fmt"
	"os"

	badger "github.com/dgraph-io/badger/v3"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("start run!")

	dir := "bdata"
	defer os.RemoveAll(dir)

	opt := badger.DefaultOptions(dir)

	db, err := badger.Open(opt)

	check(err)

	defer db.Close()

	check(err)

	key := func(i int) []byte {
		return []byte(fmt.Sprintf("%d", i))
	}

	val := func(i int) []byte {
		return []byte(fmt.Sprintf("%0128d", i))
	}

	txn1 := db.NewTransaction(true)

	defer txn1.Discard()

	check(txn1.Set([]byte("bKey"), []byte("bVal")))
	check(txn1.Commit())
	fmt.Printf("Insert Key '%s' using txn.Set\n", "bKey")

	txn2 := db.NewTransaction(false)
	entry, err := txn2.Get([]byte("bKey"))
	check(err)
	fmt.Printf("Read Key '%s' using txn.Get\n", string(entry.Key()))

	N, M := 50000, 1000

	wb := db.NewWriteBatch()

	defer wb.Cancel()

	for i := 0; i < N; i++ {
		check(wb.Set(key(i), val(i)))
	}

	check(wb.Flush())

	fmt.Println("Insert", N, "Delete", M)

	db.View(func(txn *badger.Txn) error {
		io := badger.DefaultIteratorOptions
		itr := txn.NewIterator(io)
		defer itr.Close()

		i := 0
		for itr.Rewind(); itr.Valid(); itr.Next() {
			i++
		}
		fmt.Println("Read", i, "keys")
		return nil
	})
	check(err)

}
