package repository

import (
	"context"
	"fmt"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	ArticleRepository interface {
		GetArticleByID(ctx context.Context, id models.Uuid) (*models.Article, error)
		CreateArticle(ctx context.Context, article models.Article) error
		UpdateArticle(ctx context.Context, id models.Uuid, article models.Article) error
		ListArticle(ctx context.Context) ([]models.Article, error)
		DeleteArticle(ctx context.Context, id models.Uuid) error
	}
)

func (r repository) GetArticleByID(ctx context.Context, id models.Uuid) (*models.Article, error) {
	article, err := r.memory.GetByID(id)
	if err != nil {
		return &models.Article{}, err
	}
	return &article, nil
}
func (r repository) CreateArticle(ctx context.Context, article models.Article) error {
	err := r.memory.Create(article)
	if err != nil {
		return fmt.Errorf("failed to create articel %v", err)
	}
	return nil
}
func (r repository) UpdateArticle(ctx context.Context, id models.Uuid, article models.Article) error {
	err := r.memory.Update(id, article)
	if err != nil {
		return err
	}
	return nil
}
func (r repository) ListArticle(ctx context.Context) ([]models.Article, error) {
	articles, err := r.memory.List()
	if err != nil {
		return []models.Article{}, err
	}
	return articles, nil
}
func (r repository) DeleteArticle(ctx context.Context, id models.Uuid) error {
	err := r.memory.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
