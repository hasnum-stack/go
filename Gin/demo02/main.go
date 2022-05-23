package main

import (
	"fmt"
	"net/http"

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

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "zzh123456")
	})

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
		err := ctx.ShouldBind(user)
		if err == nil {
			fmt.Printf("user: %v\n", user)
			ctx.JSON(http.StatusOK, user)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	})

	// a := "123"
	// Pa := &a
	// *Pa = "456"
	// fmt.Printf("a: %v\n", Pa)

	// ATest := &test{
	// 	name: "zzh",
	// 	age:  19,
	// }
	// BTest := ATest
	// BTest.age = 20
	// fmt.Printf("ATest: %v\n", ATest)

	r.Run(":8081")
}
