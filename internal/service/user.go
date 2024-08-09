package service

import (
	"context"
	"fmt"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type UserService interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id models.Uuid) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, id models.Uuid, user models.User) error
	DeleteUser(ctx context.Context, id models.Uuid) error
}

func (s service) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.repository.ListUsers(ctx)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (s service) GetUserByID(ctx context.Context, id models.Uuid) (*models.User, error) {
	user, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (s service) DeleteUser(ctx context.Context, id models.Uuid) error {
	err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) CreateUser(ctx context.Context, user models.User) error {
	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user: %s", err.Error())
	}
	return err
}

func (s service) UpdateUser(ctx context.Context, id models.Uuid, user models.User) error {
	err := s.repository.UpdateUser(ctx, id, user)
	if err != nil {
		return err
	}
	return nil
}
