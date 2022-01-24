package service

import (
	"context"
	"integration-test/api"
	"integration-test/database"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

type UserTestSuite struct {
	suite.Suite

	ctx     context.Context
	mockCtl *gomock.Controller
	db      *database.MockDatabase

	userService *user
}

func (s *UserTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockCtl = gomock.NewController(s.T())
	s.db = database.NewMockDatabase(s.mockCtl)
	s.userService = &user{db: s.db}
}

func (s *UserTestSuite) TearDownTest() {
	s.mockCtl.Finish()
}

func (s *UserTestSuite) TestCreateUser() {
	createdUserID := int64(1)
	s.db.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(createdUserID, nil)

	userID, err := s.userService.CreateUser(s.ctx, "zhang san", 16)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), createdUserID, userID)
}

func (s *UserTestSuite) TestCreateUserInvalidName() {
	_, err := s.userService.CreateUser(s.ctx, "", 16)
	assert.ErrorIs(s.T(), err, api.ErrInvalidUserName)
}

func (s *UserTestSuite) TestCreateUserInvalidAge() {
	_, err := s.userService.CreateUser(s.ctx, "zhang san", -1)
	assert.ErrorIs(s.T(), err, api.ErrInvalidUserAge)
}

func (s *UserTestSuite) TestDeleteUser() {
	s.db.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(nil)

	err := s.userService.DeleteUser(s.ctx, 1)
	assert.NoError(s.T(), err)
}

func (s *UserTestSuite) TestDeleteUserInvalidUserID() {
	err := s.userService.DeleteUser(s.ctx, 0)
	assert.ErrorIs(s.T(), err, api.ErrInvalidUserID)
}
