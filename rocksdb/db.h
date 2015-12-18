#ifndef DB_H
#define DB_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct DBEngine DBEngine;

int DBOpen(DBEngine **db);
void DBClose(DBEngine *db);

#ifdef __cplusplus
}
#endif

#endif
