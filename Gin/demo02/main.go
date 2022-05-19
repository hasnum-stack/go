package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type test struct {
	name string
	age  int
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "zzh123456")
	})

	a := "123"
	Pa := &a
	*Pa = "456"
	fmt.Printf("a: %v\n", Pa)

	ATest := &test{
		name: "zzh",
		age:  19,
	}
	BTest := ATest
	BTest.age = 20
	fmt.Printf("ATest: %v\n", ATest)

	r.Run(":8081")
}
