package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/satori/go.uuid"
)

func (u Notification) TableName() string {
	return "notifications"
}

type Notification struct {
	OrderID uuid.UUID `json:"orderId" gorm:"type:uuid; not null; unique"`
	Status  bool
}

func init() {

	mx := &gormigrate.Migration{
		ID:       "0003",
		Migrate:  mixture.CreateTableM(&Notification{}),
		Rollback: mixture.DropTableR(&Notification{}),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
