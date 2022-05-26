package routes

import (
	"demo02/controllers/admin"
	"demo02/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutesInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	adminRouter.Use(middlewares.InitMiddleware)
	{
		adminRouter.GET("/create", admin.Controller{}.Create)
		adminRouter.GET("/delete/:id", admin.Controller{}.Delete)
		adminRouter.GET("/update/:id", admin.Controller{}.Update)
		adminRouter.GET("/list", admin.Controller{}.List)
	}
}
