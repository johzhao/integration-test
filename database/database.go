package database

import (
	"context"
	"fmt"
	"integration-test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockgen -source=database.go -destination=./mock_database.go -package=database

const (
	invalidUserID = int64(-1)
)

type Database interface {
	Open(dataSourceName string) error
	Close() error

	InsertUser(ctx context.Context, user *model.User) (int64, error)
	DeleteUser(ctx context.Context, userID int64) error
	FindUserByID(ctx context.Context, userID int64) (*model.User, error)
	SearchUsersByName(ctx context.Context, userName string) ([]*model.User, error)
}

func NewDatabase() Database {
	return &database{}
}

type database struct {
	db *gorm.DB
}

func (d *database) Open(dataSourceName string) error {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm open: %w", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("auto migrate user: %w", err)
	}

	d.db = db

	return nil
}

func (d *database) Close() error {
	return nil
}

func (d *database) InsertUser(ctx context.Context, user *model.User) (int64, error) {
	err := d.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return invalidUserID, fmt.Errorf("db create: %w", err)
	}

	return user.ID, nil
}

func (d *database) DeleteUser(ctx context.Context, userID int64) error {
	err := d.db.WithContext(ctx).Delete(&model.User{ID: userID}).Error
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (d *database) FindUserByID(ctx context.Context, userID int64) (*model.User, error) {
	var result model.User
	err := d.db.WithContext(ctx).Where(&model.User{ID: userID}).First(&result).Error
	if err != nil {
		return nil, fmt.Errorf("db first: %w", err)
	}

	return &result, nil
}

func (d *database) SearchUsersByName(ctx context.Context, userName string) ([]*model.User, error) {
	var result []*model.User
	err := d.db.WithContext(ctx).Where(&model.User{Name: userName}).Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("db find: %w", err)
	}

	return result, nil
}
