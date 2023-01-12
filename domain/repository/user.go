package repository

import (
	"context"
	"simple-api/domain/dto"
	"simple-api/domain/entity"
	"simple-api/helpers"
	"simple-api/store"
)

type UserRepository struct{}

func (u *UserRepository) FindAll() ([]dto.UserDTO, error) {
	result := []entity.UserEntity{}
	dto := []dto.UserDTO{}
	if err := store.Database().
		WithContext(context.Background()).
		Find(&result).
		Error; err != nil {
		return dto, err
	}
	helpers.MapLoose(result, &dto)
	return dto, nil
}
