package routes

import (
	"demo02/controllers/upload"

	"github.com/gin-gonic/gin"
)

func UploadRoutesInit(r *gin.Engine) {
	userRouter := r.Group("/upload")
	{
		userRouter.POST("/image", upload.Image)
		userRouter.POST("/multiple", upload.Multiple)

	}
}
