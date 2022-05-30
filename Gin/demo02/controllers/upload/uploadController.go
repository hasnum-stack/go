package upload

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func Image(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	dst := path.Join("./static/upload", file.Filename)
	ctx.SaveUploadedFile(file, dst)

	fmt.Println(file.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func Multiple(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["multiple"]
	for _, file := range files {
		fmt.Printf("file: %v\n", file.Filename)
		dst := path.Join("./static/upload", file.Filename)
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"files": files,
	})
}
