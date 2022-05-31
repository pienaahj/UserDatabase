package userdatabase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
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
