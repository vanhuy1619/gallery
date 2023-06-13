package main

import (
	"awesomeProject2/datasource"
	"awesomeProject2/middleware"
	"awesomeProject2/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	var db = datasource.ConfigData()
	router := gin.Default()

	//router.Use()

	v1 := router.Group("/api")
	{
		v1.POST("/items", middleware.AuthMiddleware(), repositories.CreateItem(db))    // create item
		v1.GET("/items", middleware.AuthMiddleware(), repositories.GetListOfItems(db)) // list items
		v1.GET("/items/:id", repositories.ReadItemById(db))                            // get an item by ID
		v1.PUT("/items/:id", repositories.EditItemById(db))                            // edit an item by ID
		v1.DELETE("/items/:id", repositories.DeleteItemById(db))                       // delete an item by ID

		v1.POST("/signup", repositories.Regist(db))
		v1.POST("/login", repositories.Login(db))

	}

	router.Run()
}
