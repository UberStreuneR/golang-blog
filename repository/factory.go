package repository

import (
	repo "fiber-pg-blog/repository/gorm_repo"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateSqliteRepository() *repo.GormRepository {
	dialector := sqlite.Open("test.db")
	return repo.CreateGormRepository(&dialector, &gorm.Config{})
}

func CreatePostgresDialector() (dialector *gorm.Dialector, config gorm.Option) {
	return
}
