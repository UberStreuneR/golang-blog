package tests

import (
	"fmt"
	"testing"

	"fiber-pg-blog/entity"
	"fiber-pg-blog/repository"
	repo "fiber-pg-blog/repository/sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/drivers/sqlite"
)

type RepoTestSuite struct {
	suite.Suite
	Repository repository.Repository
}

func (suite *RepoTestSuite) TearDownSuite() {
	suite.Repository.DeleteAllUsers()
}

func (suite *RepoTestSuite) AfterTest() {
	suite.Repository.DeleteAllUsers()
	fmt.Println("After test run")
}

func (suite *RepoTestSuite) TestUserDoesNotExist() {
	assert := assert.New(suite.T())
	user, err := suite.Repository.GetUser("404username")
	assert.Equal(entity.User{}, user)
	assert.EqualError(err, entity.ErrUserDoesNotExist.Error())
}

func (suite *RepoTestSuite) TestAddUser() {
	user := entity.User{Username: "user", Password: "password"}
	err := suite.Repository.AddUser(user)
	assert.Equal(suite.T(), err, nil)
	userObj, _ := suite.Repository.GetUser("user")
	assert.Equal(suite.T(), userObj.ID, uint(1))
	assert.Equal(suite.T(), userObj.Username, "user")
	assert.Equal(suite.T(), userObj.Password, "password")
	suite.Repository.DeleteUser(userObj.ID)
}

func (suite *RepoTestSuite) TestDeleteUser() {
	user := entity.User{Username: "user_to_delete", Password: "password"}
	suite.Repository.AddUser(user)
	// assert.EqualError(suite.T(), err.Error, nil)
	userObj, _ := suite.Repository.GetUser("user_to_delete")
	suite.Repository.DeleteUser(userObj.ID)
	users, _ := suite.Repository.GetAllUsers()
	assert.Empty(suite.T(), users)
}

func TestRepos(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	sqliteRepo := repo.CreateSqliteRepository(db)
	suite.Run(t, &RepoTestSuite{Repository: sqliteRepo})
}
