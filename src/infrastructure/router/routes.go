package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rgalicia0729/crud-go/src/infrastructure/handler"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/api/employees", handler.FindAllEmployees)
	r.GET("/api/employees/:employeeId", handler.FindEmployeeById)
	r.POST("/api/employees", handler.CreateEmployee)
	r.PUT("/api/employees/:employeeId", handler.UpdateEmployee)
	r.DELETE("/api/employees/:employeeId", handler.DeleteEmployee)
}
