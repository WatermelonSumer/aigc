package assistant

import "github.com/gin-gonic/gin"

func AsstRouters(r *gin.RouterGroup) {
	api := r.Group("/asst")

	services := &AsstServiceImpl{}

	api.POST("/list", services.List)
	api.POST("/detail", services.Detail)
	api.POST("/retrieve", services.Retrieve)
	api.POST("/delete", services.Delete)
	api.POST("/create", services.Create)
	api.POST("/update", services.Update)
}
