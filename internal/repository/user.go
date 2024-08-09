package repository

import (
	"context"
	"fmt"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	UserRepository interface {
		GetUserByID(ctx context.Context, id models.Uuid) (*models.User, error)
		CreateUser(ctx context.Context, user models.User) error
		UpdateUser(ctx context.Context, id models.Uuid, user models.User) error
		ListUsers(ctx context.Context) ([]models.User, error)
		DeleteUser(ctx context.Context, id models.Uuid) error
	}
)

func (r repository) GetUserByID(ctx context.Context, id models.Uuid) (*models.User, error) {
	user, err := r.db.GetByID(id)
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (r repository) CreateUser(ctx context.Context, user models.User) error {
	err := r.db.Create(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r repository) UpdateUser(ctx context.Context, id models.Uuid, user models.User) error {
	err := r.db.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := r.db.List()
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (r repository) DeleteUser(ctx context.Context, id models.Uuid) error {
	err := r.db.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
