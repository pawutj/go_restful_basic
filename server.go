package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	// r := gin.Default()

	// r.GET("/ping" , func(c *gin.Context){
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"message":"ping",
	// 	})
	// })

	r := setupRouter()


	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	h := CustomerHandler{}
	h.Initialize()
	r.GET("/customers",h.GetAllCustomer)
	r.POST("/customers"  ,h.SaveCustomer)
	return r
}

type CustomerHandler struct{
	DB *gorm.DB
}

func (h *CustomerHandler) Initialize(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/finalproject?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(err) 
	}

	db.AutoMigrate(&Customer{})
	h.DB =db

}

func (h *CustomerHandler) GetAllCustomer(c *gin.Context)  {
	customers := []Customer{}
	h.DB.Find(&customers)

	c.JSON(http.StatusOK,customers)
}

func (h *CustomerHandler) SaveCustomer(c *gin.Context){
	customer := Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&customer);err != nil{
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK,customer)

}

type Customer struct{
	Id uint `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstname"`
	Age int `json:"age"`
}