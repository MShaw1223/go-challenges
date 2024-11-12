package dbstorage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
	Age   int
}

func InitDB() error{
	var err error
	db, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})

	if err != nil {
		return err
	}
	return db.AutoMigrate(&User{})
}


func CreateTable() {
	user := User{Name: "Miller"}
	result := db.Select("Name").Create(&user)
	if result.Error != nil {
		fmt.Println("Error adding name: ", result.Error)
	} else {
		fmt.Println("Added name: ", user.Name)
	}
}
