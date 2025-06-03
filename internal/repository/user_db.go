package repository

import (
	"github.com/guide-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
}
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r userRepo) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r userRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r userRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepo) GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
