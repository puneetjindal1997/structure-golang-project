package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloUser(c *gin.Context) {
	fmt.Println("Hello")
}
