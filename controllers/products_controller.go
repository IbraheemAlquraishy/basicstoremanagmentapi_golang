package controllers

import (
	"net/http"

	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/services"
	"github.com/gin-gonic/gin"
)

func Getallproducts(c *gin.Context) {
	b := services.Queryallproducts()
	print(b)
	c.JSON(http.StatusOK, b)
}
