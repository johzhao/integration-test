package main

import (
	"fmt"
	"integration-test/database"
	"integration-test/service"
	"log"
)

const (
	datasource = "root:root@tcp(localhost:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
)

func start() error {
	db := database.NewDatabase()
	if err := db.Open(datasource); err != nil {
		return fmt.Errorf("db open: %w", err)
	}
	defer func() { _ = db.Close() }()

	userService := service.NewUserService(db)
	_ = userService

	// xxx: start the web service

	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatalf("start failed with error: %v", err)
	}
}
