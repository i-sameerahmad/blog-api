package repository

import (
	"context"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	CategoryRepository interface {
		GetCategoryByID(ctx context.Context, id models.Uuid) (models.Category, error)
		CreateCategory(ctx context.Context, category models.Category) error
		UpdateCategory(ctx context.Context, id models.Uuid, category models.Category) error
		ListCategory(ctx context.Context) ([]models.Category, error)
		DeleteCategory(ctx context.Context, id models.Uuid) error
	}
)

func (r *repository) GetCategoryByID(ctx context.Context, id models.Uuid) (models.Category, error) {
	category, err := r.file.GetByID(id)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *repository) CreateCategory(ctx context.Context, category models.Category) error {
	if err := r.file.Create(category); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateCategory(ctx context.Context, id models.Uuid, category models.Category) error {
	err := r.file.Update(id, category)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ListCategory(ctx context.Context) ([]models.Category, error) {
	categories, err := r.file.List()
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}

func (r *repository) DeleteCategory(ctx context.Context, id models.Uuid) error {
	err := r.file.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
