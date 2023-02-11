package gateway

import "github.com.br/tulio/dd-ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
