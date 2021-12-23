package service

import (
	"context"

	"github.com/caogonghui/memrizr/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserResponsitory
}

type USConfig struct {
	UserRepository model.UserResponsitory
}

func NewUserService(config *USConfig) model.UserService {
	return &UserService{
		UserRepository: config.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	return s.UserRepository.FindByID(ctx, uid)
}
