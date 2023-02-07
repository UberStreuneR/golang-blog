package tests

import (
	"fiber-pg-blog/repository"
	repo "fiber-pg-blog/repository/gorm_repo"
	"fiber-pg-blog/usecase/users"
	"testing"
)

func TestRepoInterface(t *testing.T) {
	var _ repository.Repository = (*repo.GormRepository)(nil)
}

func TestUsecaseInterface(t *testing.T) {
	var _ users.UserUsecase = (*users.UserInteractor)(nil)
}
