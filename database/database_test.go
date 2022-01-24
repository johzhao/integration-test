package database

import (
	"context"
	"database/sql"
	"integration-test/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDatabase(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}

type DatabaseTestSuite struct {
	suite.Suite

	ctx   context.Context
	sqlDB *sql.DB
	mock  sqlmock.Sqlmock
	db    *database
}

func (s *DatabaseTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(s.T(), err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))
	assert.NoError(s.T(), err)

	s.ctx = context.Background()
	s.sqlDB = db
	s.mock = mock
	s.db = &database{db: gormDB}
}

func (s *DatabaseTestSuite) TearDownTest() {
	err := s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)

	s.mock.ExpectClose()

	err = s.sqlDB.Close()
	assert.NoError(s.T(), err)
}

func (s *DatabaseTestSuite) TestInsertUser() {
	userName := "zhang san"
	userAge := 16
	createdUserID := int64(1)

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `users`").
		WithArgs(userName, userAge).
		WillReturnResult(sqlmock.NewResult(createdUserID, 1))
	s.mock.ExpectCommit()

	user := model.User{
		Name: userName,
		Age:  userAge,
	}
	userID, err := s.db.InsertUser(s.ctx, &user)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), createdUserID, userID)
}

func (s *DatabaseTestSuite) TestDeleteUser() {
	userID := int64(1)

	s.mock.ExpectBegin()
	s.mock.ExpectExec("DELETE FROM `users`").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := s.db.DeleteUser(s.ctx, userID)
	assert.NoError(s.T(), err)
}
