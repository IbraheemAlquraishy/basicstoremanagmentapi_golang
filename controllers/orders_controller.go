package controllers

import (
	"fmt"
	"net/http"

	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"
	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/services"
	"github.com/gin-gonic/gin"
)

func Post_order(c *gin.Context) {
	var order configs.Postoder
	var neworder configs.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not valid request body",
		})

	}
	fmt.Println(order)
	p := services.Getordersproduct(order)
	if p.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no such product",
		})

	} else {
		if services.Checkquantity(order, p) {
			neworder.Productid = order.Productid
			neworder.Quantity = order.Quantity
			neworder.Location = order.Location
			neworder.Price = services.Gettotalprice(order, p)
			services.Createorder(neworder)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})

		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "we dont have that quantity of this product",
			})
		}
	}
}
