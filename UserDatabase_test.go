package userdatabase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_UserDatabase "github.com/pienaahj/UserDatabase/mocks"
	"github.com/pienaahj/UserDatabase/models"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

// set up the TestSuite
type UserDatabaseSuite struct {
	suite.Suite
}

func TestUserDatabaseSuite(t *testing.T) {
	suite.Run(t, new(UserDatabaseSuite))
}

// setup function for gomockcontroller
var ctrl *gomock.Controller
var u *mock_UserDatabase.MockUserDatabase

// setup the Test function
func (s *UserDatabaseSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())
	u = mock_UserDatabase.NewMockUserDatabase(ctrl)

}

// clean up the TestSuite
func (u *UserDatabaseSuite) TearDown() {
	ctrl.Finish()
}

func (s *UserDatabaseSuite) TestUserDatabase() {
	ctx := context.Background()
	username := "mocked-db-user"
	usr := &models.User{Username: username}
	filter := bson.M{"error": true}
	filterNoError := bson.M{"error": false}
	u.EXPECT().Create(ctx, usr).Return(nil)
	u.EXPECT().FindOne(ctx, filter).Return(nil, errors.New("mocked-user-db-error"))
	u.EXPECT().FindOne(ctx, filterNoError).Return(usr, nil)
	u.EXPECT().DeleteByUsername(ctx, username).Return(nil)
}
