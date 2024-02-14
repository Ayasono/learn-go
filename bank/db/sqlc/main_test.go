package db

import (
  "database/sql"
  "log"
  "testing"

  _ "github.com/lib/pq"
)

const db = "postgres"
const dbSource = "postgresql://root:123456@localhost:5432/postgres?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
  conn, err := sql.Open(db, dbSource)
  if err != nil {
    log.Fatal("cannot connect to db:", err)
  }

  defer func(conn *sql.DB) {
    err := conn.Close()
    if err != nil {
      log.Fatal("conn has closed db:", err)
    }
  }(conn)

  testQueries = New(conn)

  m.Run()
}
