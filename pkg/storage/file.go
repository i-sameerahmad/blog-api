package storage

import (
	"encoding/gob"
	"errors"
	"io"
	"log"
	"os"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	FileStorage struct {
		fileHndler *os.File
	}

	categoryType map[models.Uuid]models.Category
)

var ()

func NewFileStorage(fileName string) (*FileStorage, error) {
	file, err := os.Create(fileName)
	// file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	fileStorage := &FileStorage{fileHndler: file}
	return fileStorage, nil
}

func (f *FileStorage) Create(category models.Category) error {
	//load existing data
	categories := make(categoryType)
	_, err := f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	e := gob.NewDecoder(f.fileHndler)

	if err := e.Decode(&categories); err != nil {
		if err == io.EOF {
			log.Println(err)
		} else {
			return err
		}
	}
	//append new data
	id := models.Uuid(len(categories)) + 1
	category.Id = id
	categories[id] = category
	// reset file pointer to overwrite the contents
	_, err = f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	err = f.fileHndler.Truncate(0)
	if err != nil {
		return err
	}
	// encode and write to file
	encoder := gob.NewEncoder(f.fileHndler)
	if err := encoder.Encode(categories); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (f *FileStorage) List() ([]models.Category, error) {

	var categories categoryType
	//reset file pointer to start of file
	_, err := f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return []models.Category{}, err
	}
	e := gob.NewDecoder(f.fileHndler)

	if err := e.Decode(&categories); err != nil {
		log.Println(err)
		return []models.Category{}, err
	}

	categoriesList := make([]models.Category, 0)
	for _, v := range categories {
		categoriesList = append(categoriesList, v)
	}
	return categoriesList, nil
}

func (f *FileStorage) GetByID(id models.Uuid) (models.Category, error) {
	var categories categoryType
	//reset file pointer to start of file
	_, err := f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return models.Category{}, err
	}
	e := gob.NewDecoder(f.fileHndler)

	if err := e.Decode(&categories); err != nil {
		log.Println(err)
		return models.Category{}, err
	}
	category := categories[id]
	return category, nil
}

func (f *FileStorage) Update(id models.Uuid, category models.Category) error {
	var categories categoryType
	//reset file pointer to start of file
	_, err := f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	e := gob.NewDecoder(f.fileHndler)

	if err := e.Decode(&categories); err != nil {
		log.Println(err)
		return err
	}
	if _, ok := categories[id]; !ok {
		return errors.New("category not found")
	}

	oldCategory := categories[id]
	oldCategory.Name = category.Name
	oldCategory.Description = category.Description
	categories[id] = oldCategory
	log.Println(categories)

	_, err = f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	err = f.fileHndler.Truncate(0)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(f.fileHndler)
	if err := encoder.Encode(categories); err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (f *FileStorage) Delete(id models.Uuid) error {
	var categories categoryType
	//reset file pointer to start of file
	_, err := f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	e := gob.NewDecoder(f.fileHndler)

	if err := e.Decode(&categories); err != nil {
		log.Println(err)
		return err
	}
	if _, ok := categories[id]; !ok {
		return errors.New("category not found")
	}
	delete(categories, id)

	_, err = f.fileHndler.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	err = f.fileHndler.Truncate(0)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(f.fileHndler)
	if err := encoder.Encode(categories); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
