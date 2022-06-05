package cookie

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取
func Get(ctx *gin.Context) {

}

//设置
func Set(ctx *gin.Context) {
	host := ctx.Request.Host

	fmt.Printf("host: %v\n", host)
	ctx.SetCookie("test", "1234", 0, "/", host, false, false)
	ctx.JSON(http.StatusOK, "ok")
}
