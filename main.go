package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hero/aigc/aigc/apps"
	"log"
)

func main() {
	r := gin.Default()
	apps.InitRouters(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("启动失败")
		return
	}
}
