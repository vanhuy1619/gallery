package main

import (
	"awesomeProject2/activity"
	"awesomeProject2/config"
	"awesomeProject2/datasource"
	"awesomeProject2/middleware"
	"awesomeProject2/repositories"
	"awesomeProject2/workflow"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"os"
)

func init() {
	log.Println("Main initialization, load config file")
	config.LoadConfig()
}
func RunTemporalWorker() {
	temporal, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatal("Unable create temporal: error", err)
	}

	//create worker
	w := worker.New(temporal, workflow.GalerryQueueName, worker.Options{})
	w.RegisterActivity(activity.Login)
	w.RegisterActivity(activity.PostImage)
	w.RegisterActivity(activity.SharePost)

	//regist workflow
	w.RegisterWorkflow(workflow.GalleryWorkFlow)

	//start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("Unable to start Temporal worker:", err)
	}
}

func main() {

	//start temporal workflow
	//RunTemporalWorker()

	db, err := datasource.ConfigData()
	if err != nil {
		log.Fatal("Lỗi khi kết nối đến cơ sở dữ liệu:", err)
		return
	}

	// Kiểm tra kết nối
	dbConfig, err := db.DB()
	if err != nil {
		log.Fatal("Lỗi khi kiểm tra kết nối:", err)
	}

	fmt.Println("Kết nối thành công đến cơ sở dữ liệu Yugabyte DB!", dbConfig.Ping())

	router := gin.Default()

	router.Use()

	v1 := router.Group("/api")
	{
		v1.POST("/items", middleware.AuthMiddleware(), repositories.CreateItem(db))    // create item
		v1.GET("/items", middleware.AuthMiddleware(), repositories.GetListOfItems(db)) // list items
		v1.GET("/items/:id", repositories.ReadItemById(db))                            // get an item by ID
		v1.PUT("/items/:id", repositories.EditItemById(db))                            // edit an item by ID
		v1.DELETE("/items/:id", repositories.DeleteItemById(db))                       // delete an item by ID

		v1.POST("/signup", repositories.Regist(db))
		v1.POST("/login", repositories.Login(db))

		v1.POST("/user/upload/images", middleware.AuthMiddleware(), repositories.UploadImages(db))

	}

	router.Run(":" + os.Getenv("CLIENT_POST"))
}
