package service

import (
	"context"
	"log"

	"github.com/caogonghui/memrizr/account/model"
	"github.com/caogonghui/memrizr/account/model/apperrors"
	"github.com/google/uuid"
)

type userService struct {
	UserRepository model.UserResponsitory
}

type USConfig struct {
	UserRepository model.UserResponsitory
}

func NewUserService(config *USConfig) model.UserService {
	return &userService{
		UserRepository: config.UserRepository,
	}
}

func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	return s.UserRepository.FindByID(ctx, uid)
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("Unable to signup user")
		return apperrors.NewInternal()
	}
	u.Password = pw
	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
