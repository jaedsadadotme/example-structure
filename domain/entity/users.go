package entity

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	Id        uuid.UUID `gorm:"id" json:"id"`
	Firstname string    `gorm:"firstname" json:"firstname"`
	Lastname  string    `gorm:"lastname" json:"lastname"`
}

func (UserEntity) TableName() string {
	return "users"
}
