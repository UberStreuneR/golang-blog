package tests

import (
	"testing"

	"fiber-pg-blog/entity"
	"fiber-pg-blog/repository"
	repo "fiber-pg-blog/repository/sqlite"

	"log"

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

// func (suite *RepoTestSuite) TearDownSuite() {
// 	suite.Repository.DeleteAllUsers()
// 	log.Println("Tear down code run")
// }

func (suite *RepoTestSuite) TearDownTest() {
	suite.Repository.DeleteAllUsers()
	log.Println("After test run")
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
	assert.Equal(suite.T(), nil, err)
	userObj, _ := suite.Repository.GetUser("user")
	log.Println(userObj)
	assert.Equal(suite.T(), uint(1), userObj.ID,)
	assert.Equal(suite.T(), "user", userObj.Username)
	assert.Equal(suite.T(), "password", userObj.Password)
}

func (suite *RepoTestSuite) TestDeleteUser() {
	user := entity.User{Username: "user_to_delete", Password: "password"}
	suite.Repository.AddUser(user)
	userObj, _ := suite.Repository.GetUser("user_to_delete")
	suite.Repository.DeleteUser(userObj.ID)
	users, _ := suite.Repository.GetAllUsers()
	assert.Empty(suite.T(), users)
}

func (suite *RepoTestSuite) TestDeleteAllUsers() {
	user1 := entity.User{Username:"first", Password: "password"}
	user2 := entity.User{Username:"second", Password: "password"}
	suite.Repository.AddUser(user1)
	suite.Repository.AddUser(user2)
	suite.Repository.DeleteAllUsers()
	users, _ := suite.Repository.GetAllUsers()
	assert.Empty(suite.T(), users)
}	

func TestRepos(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	sqliteRepo := repo.CreateSqliteRepository(db)
	suite.Run(t, &RepoTestSuite{Repository: sqliteRepo})
}
