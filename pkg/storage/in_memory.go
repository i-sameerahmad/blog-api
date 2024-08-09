package storage

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	InMemoryStorage struct {
		articles map[models.Uuid]models.Article
		mutex    sync.Mutex
	}
)

var (
	ArticleAlreadyExists = fmt.Errorf("Article Alreay Exists")
	articleNotFound      = errors.New("Article not found")
)

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{articles: make(map[models.Uuid]models.Article)}
}

func (im *InMemoryStorage) Create(article models.Article) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	id := im.nextKey()
	article.Id = id
	im.articles[id] = article
	return nil
}
func (im *InMemoryStorage) Update(id models.Uuid, article models.Article) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	if _, exist := im.articles[id]; exist {
		art := im.articles[id]
		art.Title = article.Title
		art.Description = article.Description
		im.articles[id] = art
		return nil
	}
	return articleNotFound
}

func (im *InMemoryStorage) List() ([]models.Article, error) {
	var articles []models.Article
	for _, v := range im.articles {
		articles = append(articles, v)
	}
	return articles, nil
}
func (im *InMemoryStorage) GetByID(id models.Uuid) (models.Article, error) {
	article, exist := im.articles[id]
	if exist {
		return article, nil
	}
	return models.Article{}, articleNotFound
}

func (im *InMemoryStorage) DeleteByID(id models.Uuid) error {
	if _, exist := im.articles[id]; exist {

		delete(im.articles, id)
		return nil
	}
	return articleNotFound
}
func (im *InMemoryStorage) nextKey() models.Uuid {

	return models.Uuid(len(im.articles) + 1)
}
