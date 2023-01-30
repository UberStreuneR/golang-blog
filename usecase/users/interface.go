package users

import (
	"fiber-pg-blog/entity"
	"fiber-pg-blog/usecase/auth"
)

type UserUsecase interface {
	SayHello(user entity.User) (string, error)
	SignIn(input entity.User) (auth.AuthData, error)
	SignUp(input entity.User) error
	Get(username string) (entity.User, error)
	GetMany() ([]entity.User, error)
}
