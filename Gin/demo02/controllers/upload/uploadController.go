package upload

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

//单文件上传
func Single(ctx *gin.Context) {
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

//多文件上传
func Multiple(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["multiple"]
	//定义允许的文件
	allowExtMap := map[string]bool{
		".jpg": true,
		".png": true,
		".gif": true,
		"jpeg": true,
	}

	//校验是否合法
	for _, file := range files {
		extName := path.Ext(file.Filename)
		_, ok := allowExtMap[extName]
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    501,
				"message": strings.Join([]string{"文件扩展不合法", extName}, ""),
				"files":   file,
			})
			return
		}
	}

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"files": files,
	})
}
