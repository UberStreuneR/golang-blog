package tests

import (
	"testing"

	"fiber-pg-blog/entity"
	"fiber-pg-blog/repository"

	"log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	suite.Repository.DeleteAllPosts()
	log.Println("After test run")
}

func (suite *RepoTestSuite) TestUserDoesNotExist() {
	assert := assert.New(suite.T())
	user, err := suite.Repository.GetUser("404username")
	assert.Equal(entity.User{}, *user)
	assert.EqualError(err, entity.ErrUserDoesNotExist.Error())
}

func (suite *RepoTestSuite) TestAddUser() {
	user := entity.User{Username: "user", Password: "password"}
	err := suite.Repository.AddUser(&user)
	assert.Equal(suite.T(), nil, err)
	assert.NotEqual(suite.T(), user.ID, uint(0))
	userObj, _ := suite.Repository.GetUser("user")
	suite.T().Log(user)
	assert.Equal(suite.T(), uint(1), userObj.ID)
	assert.Equal(suite.T(), "user", userObj.Username)
	assert.Equal(suite.T(), "password", userObj.Password)
}

func (suite *RepoTestSuite) TestDeleteUser() {
	user := entity.User{Username: "user_to_delete", Password: "password"}
	suite.Repository.AddUser(&user)
	userObj, _ := suite.Repository.GetUser("user_to_delete")
	suite.Repository.DeleteUser(userObj.ID)
	users, _ := suite.Repository.GetAllUsers()
	assert.Empty(suite.T(), users)
}

func (suite *RepoTestSuite) TestDeleteAllUsers() {
	user1 := entity.User{Username: "first", Password: "password"}
	user2 := entity.User{Username: "second", Password: "password"}
	suite.Repository.AddUser(&user1)
	suite.Repository.AddUser(&user2)
	suite.Repository.DeleteAllUsers()
	users, _ := suite.Repository.GetAllUsers()
	assert.Empty(suite.T(), users)
}

func (suite *RepoTestSuite) AddPost() {
	user := entity.User{Username: "Author", Password: "password"}
	suite.Repository.AddUser(&user)
	post := entity.Post{Title: "Test post", Body: "Some body once told me", UserID: user.ID}
	err := suite.Repository.AddPost(&post)
	assert.Nil(suite.T(), err)
	postObj, err := suite.Repository.GetPost(post.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), post.ID, postObj.ID)
	assert.Equal(suite.T(), post.Title, postObj.Title)
	assert.Equal(suite.T(), post.Body, postObj.Body)
	assert.Equal(suite.T(), post.UserID, postObj.UserID)
}

func (suite *RepoTestSuite) TestDeletePost() {
	user := entity.User{Username: "Author", Password: "password"}
	suite.Repository.AddUser(&user)
	post := entity.Post{Title: "Post to delete", Body: "Body to delete", UserID: user.ID}
	err := suite.Repository.AddPost(&post)
	assert.Nil(suite.T(), err)
	postObj, err := suite.Repository.GetPost(post.ID)
	err = suite.Repository.DeletePost(postObj.ID)
	suite.T().Log(err)
	assert.Nil(suite.T(), err)
	posts, _ := suite.Repository.GetAllPosts()
	assert.Empty(suite.T(), posts)
}

func (suite *RepoTestSuite) TestDeleteAllPosts() {
	user := entity.User{Username: "first", Password: "password"}
	suite.Repository.AddUser(&user)
	post1 := entity.Post{Title: "First to delete", Body: "Body 1", UserID: user.ID}
	post2 := entity.Post{Title: "Second to delete", Body: "Body 2", UserID: user.ID}
	suite.Repository.AddPost(&post1)
	suite.Repository.AddPost(&post2)
	posts, err := suite.Repository.GetAllPosts()
	assert.NotEmpty(suite.T(), *posts)
	assert.Nil(suite.T(), err)
	err = suite.Repository.DeleteAllPosts()
	assert.Nil(suite.T(), err)
	postsInDB, err := suite.Repository.GetAllPosts()
	assert.Empty(suite.T(), *postsInDB)
}

func (suite *RepoTestSuite) TestGetPostsByUser() {
	user := entity.User{Username: "first", Password: "password"}
	suite.Repository.AddUser(&user)
	post1 := entity.Post{Title: "First to delete", Body: "Body 1", UserID: user.ID}
	post2 := entity.Post{Title: "Second to delete", Body: "Body 2", UserID: user.ID}
	suite.Repository.AddPost(&post1)
	suite.Repository.AddPost(&post2)
	posts, err := suite.Repository.GetPostsByUser(user.ID)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), posts, 2)
	assert.Equal(suite.T(), (*posts)[0].UserID, user.ID)
	assert.Equal(suite.T(), (*posts)[1].UserID, user.ID)
}

func TestRepos(t *testing.T) {
	sqliteRepo := repository.CreateSqliteRepository()
	suite.Run(t, &RepoTestSuite{Repository: sqliteRepo})
}
