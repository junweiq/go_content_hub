package dao

import (
	"errors"
	"fmt"
	"go_content_hub/internal/modal"

	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserDao struct {
	db *sqlx.DB
}

func NewUserDao(db *sqlx.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) CheckExist(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM user WHERE username = ?) AS `exists`"
	err := u.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		fmt.Printf("UserDao CheckExist() error [%v]", err)
		return false, err
	}
	return exists, nil
}

func (u *UserDao) Create(user *modal.User) error {
	query := "INSERT INTO user (username, password, nickname, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	nowTime := time.Now()
	_, err := u.db.Exec(query, user.Username, user.Password, user.Nickname, nowTime, nowTime)
	if err != nil {
		fmt.Printf("UserDao Create() error [%v]", err)
		return err
	}
	return nil
}

func (u *UserDao) FirstByUsername(username string) (*modal.User, error) {
	var user modal.User
	query := "SELECT * FROM user WHERE username = ? LIMIT 1"
	err := u.db.Get(&user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		fmt.Printf("UserDao FirstByUsername() error [%v]", err)
		return nil, err
	}
	return &user, nil
}
