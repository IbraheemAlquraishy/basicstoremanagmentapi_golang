package main

import (
	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.GET("/products", controllers.Getallproducts)

	app.POST("/order", controllers.Post_order)

	//admin panel
	app.POST("/login", controllers.Login)
	app.POST("/newadmin", controllers.Newadmin)
	app.Run("localhost:8000")

}
