package db

import "database/sql"

type Store struct {
  *Queries
  db *sql.DB
}

func NewStore(db *sql.DB) *Store {
  return &Store{
    db:      db,
    Queries: New(db),
  }
}

func (store *Store) execTx(fn func(*Queries) error) error {
  tx, err := store.db.Begin()
  if err != nil {
    return err
  }
  queries := New(tx)
  err = fn(queries)
  if err != nil {
    if rbErr := tx.Rollback(); rbErr != nil {
      return rbErr
    }
    return err
  }
  return tx.Commit()
}
