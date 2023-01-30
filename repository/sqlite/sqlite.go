package sqlite

import (
	"fiber-pg-blog/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteRepository struct {
	db *gorm.DB
}

func CreateSqliteRepository(db *gorm.DB) *SqliteRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database with sqlite")
	}
	db.AutoMigrate(&entity.User{}, &entity.Post{})
	return &SqliteRepository{db: db}
}

func (repo *SqliteRepository) GetUser(username string) (entity.User, error) {
	var userObj entity.User
	result := repo.db.First(&userObj, "username = ?", username)
	if result.Error != nil {
		return userObj, entity.ErrUserDoesNotExist
	}
	return userObj, nil
}

func (repo *SqliteRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	repo.db.Find(&users)
	return users, nil
}

func (repo *SqliteRepository) AddUser(user entity.User) error {
	result := repo.db.Create(&user)
	return result.Error
}

func (repo *SqliteRepository) DeleteUser(ID uint) error {
	result := repo.db.Delete(&entity.User{}, ID)
	return result.Error
}

func (repo *SqliteRepository) DeleteAllUsers() error {
	repo.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.User{})
	return nil
}

func (repo *SqliteRepository) GetPost(ID uint) (entity.Post, error) {
	var postObj entity.Post
	repo.db.First(&postObj, "ID = ?", ID)
	return postObj, nil
}

func (repo *SqliteRepository) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post
	repo.db.Find(&posts)
	return posts, nil
}

func (repo *SqliteRepository) AddPost(post entity.Post) error {
	repo.db.Create(&post)
	return nil
}
