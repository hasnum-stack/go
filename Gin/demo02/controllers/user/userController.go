package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "user/list",
	})
}
