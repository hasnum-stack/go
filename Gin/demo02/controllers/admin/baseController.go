package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (base BaseController) success(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func (base BaseController) fail(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}
