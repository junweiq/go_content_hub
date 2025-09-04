package dao

import (
	"errors"
	"fmt"
	"go_content_hub/internal/modal"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) CheckExist(username string) (bool, error) {
	var user modal.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		fmt.Printf("UserDao CheckExist() error [%v]", err)
		return false, err
	}
	return true, nil
}

func (u *UserDao) Create(user *modal.User) error {
	if err := u.db.Create(user).Error; err != nil {
		fmt.Printf("UserDao Create() error [%v]", err)
		return err
	}
	return nil
}

func (u *UserDao) FirstByUsername(username string) (*modal.User, error) {
	var user modal.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		fmt.Printf("UserDao FirstByUsername() error [%v]", err)
		return nil, err
	}
	return &user, nil
}
