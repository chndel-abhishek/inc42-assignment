package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/api/users", func(c *gin.Context) {
		getUsers(c, db) // Pass the db variable to getUsers
	})
	r.Run(":8085")
}

func getUsers(c *gin.Context, db *gorm.DB) {
	var users []User
	db.Find(&users)
	c.JSON(200, users)
}
