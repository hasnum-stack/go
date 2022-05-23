package routes

import (
	"demo02/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.GET("/list", user.List)
}
