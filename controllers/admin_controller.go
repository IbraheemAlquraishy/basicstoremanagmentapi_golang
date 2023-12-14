package controllers

import (
	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var log configs.Login
	err := c.BindJSON(log)
	if err != nil {
		return
	}

}
