package main

import (
  "database/sql"
  "log"

  "bank/api"
  db "bank/db/sqlc"
  _ "github.com/lib/pq"
)

const (
  dbDriver      = "postgres"
  dbSource      = "postgresql://root:123456@localhost:5432/postgres?sslmode=disable"
  serverAddress = "0.0.0.0:9090"
)

func main() {
  conn, err := sql.Open(dbDriver, dbSource)
  if err != nil {
    log.Fatal("cannot connect to db:", err)
  }

  store := db.NewStore(conn)
  server := api.NewServer(store)

  err = server.Start(serverAddress)

  if err != nil {
    log.Fatal("cannot start server:", err)
  }
}
