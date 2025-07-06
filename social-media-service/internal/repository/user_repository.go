package repository

import (
	"social-media-service/internal/model"
	"social-media-service/pkg/db"
)

type UserRepository interface {
	Create(user *model.User) error
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	return db.DB.Create(user).Error
}
