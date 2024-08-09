package service

import (
	"context"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	CategoryService interface {
		ListCategory(ctx context.Context) ([]models.Category, error)
		GetCategoryByID(ctx context.Context, id models.Uuid) (models.Category, error)
		CreateCategory(ctx context.Context, category models.Category) error
		UpdateCategory(ctx context.Context, id models.Uuid, category models.Category) error
		DeleteCategory(ctx context.Context, id models.Uuid) error
	}
)

func (s service) ListCategory(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repository.ListCategory(ctx)
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}
func (s service) GetCategoryByID(ctx context.Context, id models.Uuid) (models.Category, error) {
	category, err := s.repository.GetCategoryByID(ctx, id)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil

}
func (s service) DeleteCategory(ctx context.Context, id models.Uuid) error {
	err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s service) CreateCategory(ctx context.Context, category models.Category) error {
	if err := s.repository.CreateCategory(ctx, category); err != nil {
		return err
	}
	return nil
}
func (s service) UpdateCategory(ctx context.Context, id models.Uuid, category models.Category) error {
	if err := s.repository.UpdateCategory(ctx, id, category); err != nil {
		return err
	}

	return nil
}
