package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	AdminRoutesInit(router)
	UserRoutesInit(router)
	UploadRoutesInit(router)
	CookieRoutesInit(router)
}
