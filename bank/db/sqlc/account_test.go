package db

import (
  "context"
  "testing"

  "bank/util"
  "github.com/stretchr/testify/require"
)

func TestQueries_CreateAccount(t *testing.T) {
  arg := CreateAccountParams{
    Owner:    util.RandomOwner(),
    Balance:  util.RandomMoney(),
    Currency: util.RandomCurrency(),
  }

  account, err := testQueries.CreateAccount(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, account)

  require.Equal(t, arg.Owner, account.Owner)
  require.Equal(t, arg.Balance, account.Balance)
  require.Equal(t, arg.Currency, account.Currency)

  require.NotZero(t, account.ID)
  require.NotZero(t, account.CreatedAt)
}

func GetAccount(t *testing.T) Account {
  arg := CreateAccountParams{
    Owner:    util.RandomOwner(),
    Balance:  util.RandomMoney(),
    Currency: util.RandomCurrency(),
  }

  account, err := testQueries.CreateAccount(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, account)

  require.Equal(t, arg.Owner, account.Owner)
  require.Equal(t, arg.Balance, account.Balance)
  require.Equal(t, arg.Currency, account.Currency)

  require.NotZero(t, account.ID)
  require.NotZero(t, account.CreatedAt)

  return account
}

func TestCreateAccount(t *testing.T) {
  for i := 0; i < 5; i++ {
    GetAccount(t)
  }
}

func TestQueries_UpdateAccount(t *testing.T) {
  account := GetAccount(t)

  arg := UpdateAccountParams{
    ID:      account.ID,
    Balance: util.RandomMoney(),
  }

  updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, updatedAccount)

  require.Equal(t, account.ID, updatedAccount.ID)
  require.Equal(t, arg.Balance, updatedAccount.Balance)
  require.Equal(t, account.Owner, updatedAccount.Owner)
  require.Equal(t, account.Currency, updatedAccount.Currency)
  require.Equal(t, account.CreatedAt, updatedAccount.CreatedAt)
}

func TestQueries_DeleteAccount(t *testing.T) {
  // Create an account first
  createArg := CreateAccountParams{
    Owner:    util.RandomOwner(),
    Balance:  util.RandomMoney(),
    Currency: util.RandomCurrency(),
  }

  account, err := testQueries.CreateAccount(context.Background(), createArg)
  require.NoError(t, err)
  require.NotEmpty(t, account)

  // Delete the created account
  deletedAccount, err := testQueries.DeleteAccount(context.Background(), account.ID)
  require.NoError(t, err)
  require.NotEmpty(t, deletedAccount)

  // Try to get the deleted account
  _, err = testQueries.GetAccount(context.Background(), account.ID)

  // We expect an error here because the account was deleted
  require.Error(t, err)
}

func TestQueries_ListAccounts(t *testing.T) {
  for i := 0; i < 10; i++ {
    GetAccount(t)
  }

  arg := ListAccountsParams{
    Limit:  5,
    Offset: 5,
  }

  accounts, err := testQueries.ListAccounts(context.Background(), arg)
  require.NoError(t, err)
  require.Len(t, accounts, 5)

  for _, account := range accounts {
    require.NotEmpty(t, account)
  }
}
