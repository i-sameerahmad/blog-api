package repository

import "github.com/ImranZahoor/blog-api/pkg/storage"

type (
	Repository interface {
		ArticleRepository
		UserRepository
		CategoryRepository
	}
	repository struct {
		memory *storage.InMemoryStorage
		db     *storage.MySQLStorage
		file   *storage.FileStorage
	}
)

func NewRepository(memStorage *storage.InMemoryStorage, dbStorage *storage.MySQLStorage, fileStorage *storage.FileStorage) Repository {
	return &repository{
		memory: memStorage,
		db:     dbStorage,
		file:   fileStorage,
	}

}
