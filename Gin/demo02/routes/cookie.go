package routes

import (
	"demo02/controllers/cookie"

	"github.com/gin-gonic/gin"
)

func CookieRoutesInit(r *gin.Engine) {
	router := r.Group("/cookie")
	{
		router.GET("/set", cookie.Set)
		router.GET("/get", cookie.Get)
	}
}
