package dbstorage

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

type User struct {
	gorm.Model
	ID   int
	Name string
}

func CreateTable() {
	user := User{Name: "Miller"}
	result := db.Select("Name").Create(&user)
	fmt.Println("Added name: ", result)
}

func main() {
	CreateTable()
}
