package create_account

import (
	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	"github.com.br/dlcurado/fc-ws-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	ClientGateway  gateway.ClientGateway
	AccountGateway gateway.AccountGateway
}

func NewCreateAccountUseCase(clientGateway gateway.ClientGateway, accountGateway gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		ClientGateway:  clientGateway,
		AccountGateway: accountGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client, 0)

	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: client.ID,
	}, nil
}
