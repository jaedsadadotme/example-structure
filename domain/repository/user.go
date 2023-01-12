package repository

import (
	"simple-api/domain/dto"
)

type UserRepository struct{}

func (u *UserRepository) FindAll() (dto.UserDTO, error) {
	result := dto.UserDTO{}
	return result, nil
}
