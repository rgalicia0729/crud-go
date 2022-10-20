package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rgalicia0729/crud-go/src/domain/values"
	"github.com/rgalicia0729/crud-go/src/infrastructure/repository"
	"github.com/rgalicia0729/crud-go/src/usecase"
	"log"
	"net/http"
)

func CreateEmployee(c *gin.Context) {
	type RequestBody struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}

	requestBody := new(RequestBody)
	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	employeeRepository := repository.NewEmployee()
	employeeUseCase := usecase.NewEmployee(employeeRepository)

	employeeValues := values.Employee{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
	}

	employeeResult, err := employeeUseCase.CreateEmployee(&employeeValues)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"employee": employeeResult,
	})
}
