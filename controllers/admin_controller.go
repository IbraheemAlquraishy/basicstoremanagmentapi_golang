package controllers

import (
	"fmt"
	"net/http"

	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/configs"
	"github.com/IbraheemAlquraishy/basicstoremanagmentapi_golang/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var log configs.Login
	err := c.BindJSON(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad json form",
		})
		return
	}
	admin, err := services.Checklog(log)
	fmt.Println(err)
	if err == nil {
		token := configs.Gentoken(admin)
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", token, 3600*24, "", "", true, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "invalid name or password",
		})
	}
}

func Newadmin(c *gin.Context) {
	token, _ := c.Cookie("Authorization")

	_, err := configs.Verifytoken(token)
	fmt.Println(err)
	if err != nil {
		return
	}

	var log configs.Login
	err = c.BindJSON(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad json form",
		})
		return
	}
	err = services.Insertadmin(log)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "something",
		})

	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "done",
		})
	}
}
