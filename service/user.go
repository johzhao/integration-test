package service

import (
	"context"
	"fmt"
	"integration-test/api"
	"integration-test/database"
	"integration-test/model"
)

type User interface {
	CreateUser(ctx context.Context, name string, age int) (int64, error)
	DeleteUser(ctx context.Context, userID int64) error
}

func NewUserService(db database.Database) User {
	return &user{
		db: db,
	}
}

type user struct {
	db database.Database
}

func (u *user) CreateUser(ctx context.Context, name string, age int) (int64, error) {
	if len(name) == 0 {
		return -1, api.ErrInvalidUserName
	}
	if age < 0 {
		return -1, api.ErrInvalidUserAge
	}

	userID, err := u.db.InsertUser(ctx, &model.User{
		Name: name,
		Age:  age,
	})
	if err != nil {
		return -1, fmt.Errorf("db insert user: %w", err)
	}

	return userID, nil
}

func (u *user) DeleteUser(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return api.ErrInvalidUserID
	}

	if err := u.db.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("db delete user: %w", err)
	}

	return nil
}
