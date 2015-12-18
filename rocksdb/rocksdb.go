package rocksdb

import (
	_ "github.com/cockroachdb/c-rocksdb"

	// #cgo CXXFLAGS: -std=c++11
	// #cgo CPPFLAGS: -I ../../../cockroachdb/c-rocksdb/internal/include
	// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
	// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
	// #include "db.h"
	"C"

	"errors"
)

type RocksDB struct {
	db *C.DBEngine
}

func NewRocksDB() *RocksDB {
	return &RocksDB{}
}

func (r *RocksDB) Open() error {
	if r.db != nil {
		return nil
	}

	status := C.DBOpen(&r.db)
	if status != 0 {
		return errors.New("failed to open db")
	}

	return nil
}

func (r *RocksDB) Close() error {
	if r.db == nil {
		return nil
	}

	C.DBClose(r.db)
	return nil
}

func (r *RocksDB) Put() error {
	if r.db == nil {
		return nil
	}

	C.DBPut(r.db)
	return nil
}
