package database

import (
	"database/sql"

	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
)

type AccountDb struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDb {
	return &AccountDb{
		DB: db,
	}
}

func (c *AccountDb) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := c.DB.Prepare("SELECT a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID, &account.ClientID, &account.Balance, &account.CreatedAt,
		&client.ID, &client.Name, &client.Email, &client.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *AccountDb) Save(account *entity.Account) error {

	stmt, err := c.DB.Prepare("INSERT INTO accounts VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (c *AccountDb) UpdateBalance(account *entity.Account) error {

	stmt, err := c.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}

	return nil
}
