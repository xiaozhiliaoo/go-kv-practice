package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var bucketName = []byte("bk1")

func main() {
	// Open the my-bolt.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my-bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := []byte("hello")
	value := []byte("world")

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			err := bucket.Put(key, value)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", bucketName)
		}
		val := bucket.Get(key)
		fmt.Printf(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
