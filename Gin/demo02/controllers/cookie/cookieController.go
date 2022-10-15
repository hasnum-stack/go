package cookie

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取
func Get(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	fmt.Printf("err: %v\n", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, username)
}

//设置
func Set(ctx *gin.Context) {
	host := ctx.Request.Host

	fmt.Printf("host: %v\n", host)
	ctx.SetCookie("username", "hasnum", 0, "", host, false, false)
	ctx.JSON(http.StatusOK, "ok")
}
