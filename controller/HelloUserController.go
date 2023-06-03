package controller

import (
	"fmt"
	"net/http"
	"structure/database"
	"structure/models"

	"github.com/gin-gonic/gin"
)

func HelloUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("Can't bind data")
		return
	}

	// database
	err = database.InsertData(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error occured"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
