package auth

import (
	"fiber-pg-blog/entity"
)

type AuthData struct {
	Username string
	Token    string
}

type AuthUsecase interface {
	ComparePasswords(password1, password2 string) bool
	CreateAuthData(username string) (AuthData, error)
	Authenticate(details, input *entity.User) (AuthData, error)
	HashAndSalt(password string) (string, error)
}
