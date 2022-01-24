package test

import (
	"context"
	"database/sql"
	"integration-test/database"
	"integration-test/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	dialect    = "mysql"
	datasource = "root:root@tcp(localhost:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
)

func Test_Integration(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite

	ctx         context.Context
	sqlDB       *sql.DB
	db          database.Database
	userService service.User
}

func (s *IntegrationTestSuite) SetupSuite() {
	db, err := sql.Open(dialect, datasource)
	assert.NoError(s.T(), err)

	s.sqlDB = db
}

func (s *IntegrationTestSuite) TearDownSuite() {
	err := s.sqlDB.Close()
	assert.NoError(s.T(), err)
}

func (s *IntegrationTestSuite) SetupTest() {
	db := database.NewDatabase()
	err := db.Open(datasource)
	assert.NoError(s.T(), err)

	s.ctx = context.Background()
	s.db = db
	s.userService = service.NewUserService(s.db)
}

func (s *IntegrationTestSuite) TearDownTest() {
	err := s.db.Close()
	assert.NoError(s.T(), err)
}

func (s *IntegrationTestSuite) BeforeTest(suiteName string, testName string) {
	LoadDBTestFixtures(s.T(), s.sqlDB, dialect, suiteName, testName)
}

func (s *IntegrationTestSuite) TestCreateUser() {
	userName := "zhang san"
	userAge := 16
	userID, err := s.userService.CreateUser(s.ctx, userName, userAge)
	assert.NoError(s.T(), err)
	assert.Less(s.T(), int64(0), userID)
}

func (s *IntegrationTestSuite) TestDeleteUser() {
	userID := int64(10)
	err := s.userService.DeleteUser(s.ctx, userID)
	assert.NoError(s.T(), err)
}
