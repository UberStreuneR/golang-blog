package repository

import (
	"fiber-pg-blog/entity"
)

type Repository interface {
	GetUser(username string) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	AddUser(user entity.User) error
	DeleteUser(ID uint) error
	DeleteAllUsers() error

	GetPost(ID uint) (entity.Post, error)
	GetAllPosts() ([]entity.Post, error)
	AddPost(post entity.Post) error
}
