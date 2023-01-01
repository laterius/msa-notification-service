package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

// SaveStatus saving status for order
func (s *Service) SaveStatus(orderId uuid.UUID, status bool) error {
	err := s.db.Create(Notification{
		OrderID: orderId,
		Status:  status,
	}).Error
	return err
}

func (s *Service) GetStatus(orderId uuid.UUID) (notify *Notification, err error) {
	err = s.db.Model(notify).Where(orderId).First(&notify).Error
	return
}

type Notification struct {
	OrderID uuid.UUID `json:"orderId" gorm:"type:uuid; not null; unique"`
	Status  bool
}
