package database

import (
	"database/sql"
	"testing"

	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDb
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance varchar(255), created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")

	clientFrom, err := entity.NewClient("Client 1", "client1@email.com")
	s.Nil(err)
	s.clientFrom = clientFrom

	clientTo, err := entity.NewClient("Client 2", "client2@email.com")
	s.Nil(err)
	s.clientTo = clientTo

	accountFrom := entity.NewAccount(s.clientFrom, 1000)
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.clientFrom, 1000)
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}

// func (s *AccountDBTestSuite) TestSave() {
// 	account := entity.NewAccount(s.client, 1000)
// 	err := s.accountDB.Save(account)
// 	s.Nil(err)
// }

// func (s *AccountDBTestSuite) TestFindByID() {
// 	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
// 		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
// 	)

// 	account := entity.NewAccount(s.client, 1000)
// 	err := s.accountDB.Save(account)
// 	s.Nil(err)

// 	accountDB, err := s.accountDB.FindByID(account.ID)
// 	s.Nil(err)
// 	s.Equal(account.ID, accountDB.ID)
// 	s.Equal(account.ClientID, accountDB.ClientID)
// 	s.Equal(account.Balance, accountDB.Balance)
// 	s.Equal(account.Client.ID, accountDB.Client.ID)
// 	s.Equal(account.Client.Name, accountDB.Client.Name)
// 	s.Equal(account.Client.Email, accountDB.Client.Email)
// }
