package gateway

import "github.com.br/dlcurado/fc-ws-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
