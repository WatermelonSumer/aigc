package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/hero/aigc/aigc/apps/assistant"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	version := api.Group("/v1")

	assistant.AsstRouters(version)
}
