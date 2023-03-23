package main

import (
	"fmt"
	bitcask "github.com/rhainlee/bitcask-kv-go"
)

func main() {
	opts := bitcask.DefaultOptions
	opts.DirPath = "D:\\work\\temp"
	db, err := bitcask.Open(opts)
	if err != nil {
		panic(err)
	}

	db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	fmt.Println("val = ", string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}

}
