package router

import (
	"net/http"
	"structure/constant"
	"structure/controller"
)

// this routes represent user
var user = Routes{
	Route{"Hello User", http.MethodPost, constant.HelloUserPath, controller.HelloUser},
}
