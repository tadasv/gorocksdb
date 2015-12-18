package main

import (
	"github.com/tadasv/gorocksdb/rocksdb"
)

func main() {
	db := rocksdb.NewRocksDB()
	db.Open()
	db.Close()
	println("hello")
}
