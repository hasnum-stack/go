package main

import (
	"demo02/routes"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// type test struct {
// 	name string
// 	age  int
// }

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type Article struct {
	Name string `xml:"name" json:"name"`
	Desc string `xml:"desc" json:"desc"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "zzh123456")
	})

	//手动获取表单
	r.POST("/doAddUser", func(ctx *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		ctx.BindJSON(&json)
		fmt.Printf("json: %v\n", json)

		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "20")
		// ctx.SetCookie("token", "token-123", 0, ctx.Request.Host)
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.POST("/getUser", func(ctx *gin.Context) {
		user := &UserInfo{}
		if err := ctx.ShouldBind(user); err == nil {
			fmt.Printf("user: %v\n", user)
			ctx.JSON(http.StatusOK, user)
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	// 获取xml数据
	r.POST("/xml", func(ctx *gin.Context) {
		article := &Article{}
		xmlData, _ := ctx.GetRawData()

		if err := xml.Unmarshal(xmlData, article); err == nil {
			fmt.Printf("article: %v\n", article)
			ctx.JSON(http.StatusOK, article)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	//动态路由参数获取请求数据
	r.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		fmt.Printf("id: %v\n", id)
		if strings.TrimSpace(id) == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": id,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"response": gin.H{
				"id": id,
			},
		})
	})

	//路由分组
	v1 := r.Group("/v1")
	{
		v1.GET("/tmp", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, http.StateActive)
		})
	}

	routes.Init(r)
	r.Run(":8081")
}
