package users

import (
	"fiber-pg-blog/entity"
	"fiber-pg-blog/repository"
	"fiber-pg-blog/usecase/auth"
	"fmt"
)

type UserInteractor struct {
	Repository repository.Repository
	Auth       auth.AuthUsecase
}

func (interactor *UserInteractor) SayHello(user entity.User) (message string, err error) {
	details, err := interactor.Repository.GetUser(user.Username)
	if err != nil {
		return message, err
	}
	message = fmt.Sprintf("Hello %s", details.Username)
	return message, nil
}

func (interactor *UserInteractor) SignIn(input entity.User) (authData auth.AuthData, err error) {
	details, err := interactor.Repository.GetUser(input.Username)
	if err != nil {
		return authData, err
	}
	if details.Username == "" {
		return authData, entity.ErrUserDoesNotExist
	}
	authData, err = interactor.Auth.Authenticate(details, input)
	return
}

func (interactor *UserInteractor) SignUp(input entity.User) error {
	hashed, err := interactor.Auth.HashAndSalt(input.Password)
	if err != nil {
		//log here
		return err
	}
	details, err := interactor.Repository.GetUser(input.Username)
	if err != nil {
		//log here
		return err
	}
	if details.Username == input.Username {
		return entity.ErrUserAlreadyExists
	}
	input.Password = hashed
	err = interactor.Repository.AddUser(input)
	if err != nil {
		//log here
		return err
	}
	return nil
}

func (interactor *UserInteractor) Get(username string) (user entity.User, err error) {
	user, err = interactor.Repository.GetUser(username)
	if err != nil {
		//log
		return user, err
	}
	return user, nil
}

func (interactor *UserInteractor) GetMany() (users []entity.User, err error) {
	users, err = interactor.Repository.GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}
