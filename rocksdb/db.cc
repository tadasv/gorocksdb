#include "db.h"
#include "rocksdb/db.h"

extern "C" {

struct DBEngine {
    rocksdb::DB* handle;
};

int DBOpen(DBEngine **db) {
    rocksdb::Options options;
    options.create_if_missing = true;
    rocksdb::DB *handle;
    rocksdb::DB::Open(options, "/tmp/test_db", &handle);

    *db = new DBEngine;
    (*db)->handle = handle;
    return 0;
}

void DBClose(DBEngine *db) {
    delete db->handle;
    delete db;
}

void DBPut(DBEngine *db) {
    std::string value;
    rocksdb::Status s = db->handle->Put(rocksdb::WriteOptions(), "key", "value");
}

}
