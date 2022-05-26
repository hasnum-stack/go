package routes

import (
	"demo02/controllers/user"
	"demo02/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(r *gin.Engine) {
	userRouter := r.Group("/user", middlewares.InitMiddleware)
	userRouter.GET("/list", user.List)
}
