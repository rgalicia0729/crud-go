package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rgalicia0729/crud-go/src/infrastructure/handler"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/api/employees", handler.CreateEmployee)
}
