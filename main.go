package main

import (
	"PointOfSale/database"
	"PointOfSale/schema"
	"PointOfSale/seeds"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Response struct {
	Message string
}

func main() {
	r := gin.New()
	database.Connect()
	database.Database.AutoMigrate(&schema.GroupMenu{})
	database.Database.AutoMigrate(&schema.Menu{})
	database.Database.AutoMigrate(&schema.Transaction{})
	database.Database.AutoMigrate(&schema.TransactionItem{})
	database.Database.AutoMigrate(&schema.TableTransaction{})
	seeds.InsertGroupMenu()
	seeds.InsertMenu()
	seeds.InsertTable()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome Home",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/groupmenu", func(c *gin.Context) {
		var data []schema.GroupMenu
		err := database.Database.Preload(clause.Associations).Find(&data).Error
		if err != nil {
			c.JSON(http.StatusNotFound, Response{Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/menu", func(c *gin.Context) {
		var data []schema.Menu
		err := database.Database.Find(&data).Error
		if err != nil {
			c.JSON(http.StatusNotFound, Response{Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/table", func(c *gin.Context) {
		var data []schema.TableTransaction
		err := database.Database.Find(&data).Error
		if err != nil {
			c.JSON(http.StatusNotFound, Response{Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/transaction", func(c *gin.Context) {
		var data []schema.Transaction
		err := database.Database.Preload(clause.Associations).Find(&data).Error
		if err != nil {
			c.JSON(http.StatusNotFound, Response{Message: "Not Found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	var port string
	if os.Getenv("PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}
	fmt.Println("GIN_MODE:", os.Getenv("GIN_MODE"))
	r.Run(fmt.Sprintf(":%v", port))
}