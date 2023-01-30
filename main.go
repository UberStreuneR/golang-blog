package main

// import (
// 	"fmt"

// 	repo "fiber-pg-blog/repository/sqlite"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// type DB struct {
// 	db *gorm.DB
// }

// func main() {
// 	gormdb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Could not connect to db")
// 	}
// 	repo := repo.CreateSqliteRepository(gormdb)
// 	// user := entity.User{Username: "Giovanni"}
// 	// err = repo.AddUser(user)
// 	fmt.Println(err == nil)
// 	user, err := repo.GetUser("Giovanni")
// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 	}
// 	fmt.Println(user.ID, user.Username, user.Password == "", "_")
// }
