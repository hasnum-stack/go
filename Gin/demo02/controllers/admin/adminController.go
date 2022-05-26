package admin

import (
	"fmt"
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
	username, _ := ctx.Get("username")
	fmt.Printf("username: %v\n", username)
	assertUsername, ok := username.(string)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "admin/create",
			"username": "username is" + assertUsername,
		})
		return
	}
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
