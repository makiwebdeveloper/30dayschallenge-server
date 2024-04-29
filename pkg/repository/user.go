package repository

import (
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/domain"
	"gorm.io/gorm"
)

type User interface {
	Create(user *domain.User) error
	Save(user *domain.User) error
	Delete(userId int) error
	FindById(userId int) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Save(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(userId int) error {
	return r.db.Delete(&domain.User{}, userId).Error
}

func (r *UserRepository) FindById(userId int) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, userId).Error
	return &user, err
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username =?", username).First(&user).Error
	return &user, err
}
