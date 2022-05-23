package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	BaseController
}

func (controller Controller) List(ctx *gin.Context) {
	controller.success(ctx, "admin/list")
}

func (controller Controller) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "admin/create",
	})
}

func (controller Controller) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "admin/delete",
	})
}

func (controller Controller) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "admin/update",
	})
}
