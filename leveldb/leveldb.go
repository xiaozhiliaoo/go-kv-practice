package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	db, err := leveldb.OpenFile("my-level", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Put([]byte("key1"), []byte("value1"), nil)

	data, err := db.Get([]byte("key1"), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key1->", string(data))

}
