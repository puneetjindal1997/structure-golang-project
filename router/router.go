package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name       string
	Method     string
	Path       string // localhost:8000/helloworld
	HandleFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) User(g *gin.RouterGroup) {
	users := g.Group("/" + os.Getenv("UserRoute"))
	for _, singleRoute := range user {
		switch singleRoute.Method {
		case http.MethodPost:
			users.POST(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodGet:
			users.GET(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodPut:
			users.PUT(singleRoute.Path, singleRoute.HandleFunc)
		case http.MethodDelete:
			users.DELETE(singleRoute.Path, singleRoute.HandleFunc)
		}
	}
}

func Routing() {
	r := routes{
		router: gin.Default(),
	}
	grouping := r.router.Group(os.Getenv("ApiVersion"))
	r.User(grouping)
	r.router.Run(":" + os.Getenv("Port"))
}
