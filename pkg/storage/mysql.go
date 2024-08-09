package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ImranZahoor/blog-api/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLStorage struct {
	db *gorm.DB
}

var (
	UserAlreadyExists = fmt.Errorf("User Alreay Exists")
	UserNotFound      = errors.New("User not found")
)

func NewMySQLStorageInit() (*MySQLStorage, error) {
	dsn := "root:admin123@tcp(localhost:3306)/blog"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gormDB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return &MySQLStorage{db: gormDB}, nil
}

func (ms *MySQLStorage) Create(user models.User) error {
	err := ms.db.Create(&user).Error
	log.Println("MySQL::Creat::User")
	if err != nil {
		fmt.Print(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserNotFound
		}

		return err
	}

	return nil
}

func (ms *MySQLStorage) Update(id models.Uuid, user models.User) error {
	err := ms.db.Model(&models.User{Id: id}).Updates(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserNotFound
		}

		return err
	}

	return nil
}

func (ms *MySQLStorage) List() ([]models.User, error) {
	var users []models.User
	err := ms.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ms *MySQLStorage) GetByID(id models.Uuid) (models.User, error) {
	var user models.User
	err := ms.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, UserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}

func (ms *MySQLStorage) DeleteByID(id models.Uuid) error {
	err := ms.db.Delete(&models.User{Id: id}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserNotFound
		}

		return err
	}

	return nil
}
