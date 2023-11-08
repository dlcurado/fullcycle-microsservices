package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (c *Client) Update(newName string, newEmail string) error {
	updated := false

	if newName != "" {
		c.Name = newName
		updated = true

	}

	if newEmail != "" {
		c.Email = newEmail
		updated = true
	}

	if updated {
		c.UpdatedAt = time.Now()
	}

	return c.Validate()
}

func (client *Client) AddAccount(account *Account) error {
	if account.Client.ID != client.ID {
		return errors.New("account does not belong to client")
	}

	client.Accounts = append(client.Accounts, account)
	return nil
}
