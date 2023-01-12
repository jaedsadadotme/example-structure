package dto

import (
	"github.com/google/uuid"
)

var (
	UserId string
)

type UserDTO struct {
	Id        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}
