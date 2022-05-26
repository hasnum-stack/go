package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LogStatistics(ctx *gin.Context) {

	localCtx := ctx.Copy()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine记录信息" + localCtx.Request.URL.Path)
	}()
}
