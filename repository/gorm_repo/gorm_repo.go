package gorm_repo

import (
	"fiber-pg-blog/entity"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func CreateGormRepository(dialector *gorm.Dialector, config gorm.Option) *GormRepository {
	db, err := gorm.Open(*dialector, config)
	if err != nil {
		panic("Failed to connect to database with sqlite")
	}
	db.AutoMigrate(&entity.User{}, &entity.Post{})
	return &GormRepository{db: db}
}

func (repo *GormRepository) GetUser(username string) (*entity.User, error) {
	var userObj entity.User
	result := repo.db.First(&userObj, "username = ?", username)
	if result.Error != nil {
		return &userObj, entity.ErrUserDoesNotExist
	}
	return &userObj, nil
}

func (repo *GormRepository) GetAllUsers() (*[]entity.User, error) {
	var users []entity.User
	result := repo.db.Find(&users)
	return &users, result.Error
}

func (repo *GormRepository) AddUser(user *entity.User) error {
	result := repo.db.Create(user)
	return result.Error
}

func (repo *GormRepository) DeleteUser(ID uint) error {
	result := repo.db.Unscoped().Delete(&entity.User{}, ID)
	return result.Error
}

func (repo *GormRepository) DeleteAllUsers() error {
	result := repo.db.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.User{})
	return result.Error
}

func (repo *GormRepository) GetPost(ID uint) (*entity.Post, error) {
	var postObj entity.Post
	result := repo.db.First(&postObj, "ID = ?", ID)
	return &postObj, result.Error
}

func (repo *GormRepository) GetAllPosts() (*[]entity.Post, error) {
	var posts []entity.Post
	result := repo.db.Find(&posts)
	return &posts, result.Error
}

func (repo *GormRepository) AddPost(post *entity.Post) error {
	result := repo.db.Create(post)
	return result.Error
}

func (repo *GormRepository) DeletePost(ID uint) error {
	result := repo.db.Unscoped().Delete(&entity.Post{}, ID)
	return result.Error
}

func (repo *GormRepository) DeleteAllPosts() error {
	result := repo.db.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.Post{})
	return result.Error
}

func (repo *GormRepository) GetPostsByUser(ID uint) (*[]entity.Post, error) {
	var posts []entity.Post
	result := repo.db.Joins(("JOIN users ON users.ID == posts.UserID")).Find(&posts)
	return &posts, result.Error

}
