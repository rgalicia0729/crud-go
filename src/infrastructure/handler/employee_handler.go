package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rgalicia0729/crud-go/src/domain/values"
	"github.com/rgalicia0729/crud-go/src/errormessages"
	"github.com/rgalicia0729/crud-go/src/infrastructure/repository"
	"github.com/rgalicia0729/crud-go/src/usecase"
	"log"
	"net/http"
	"strconv"
)

func FindAllEmployees(c *gin.Context) {
	employeeRepository := repository.NewEmployee()
	employeeUseCase := usecase.NewEmployee(employeeRepository)

	employeesResult, err := employeeUseCase.FindAllEmployees()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employees": employeesResult,
	})
}

func FindEmployeeById(c *gin.Context) {
	employeeId, err := strconv.Atoi(c.Param("employeeId"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	employeeRepository := repository.NewEmployee()
	employeeUseCase := usecase.NewEmployee(employeeRepository)

	employeeResult, err := employeeUseCase.FindEmployeeById(employeeId)
	if err != nil {
		log.Println(err)

		if errors.Is(err, errormessages.ErrEmployeeNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee": employeeResult,
	})
}

func CreateEmployee(c *gin.Context) {
	type RequestBody struct {
		FirstName string `json:"firstName" binding:"required"`
		LastName  string `json:"lastName" binding:"required"`
		Email     string `json:"email" binding:"required"`
	}

	requestBody := new(RequestBody)
	if err := c.ShouldBindJSON(requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func UpdateEmployee(c *gin.Context) {
	employeeId, err := strconv.Atoi(c.Param("employeeId"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	type RequestBody struct {
		FirstName string `json:"firstName" binding:"required"`
		LastName  string `json:"lastName" binding:"required"`
		Email     string `json:"email" binding:"required"`
	}

	requestBody := new(RequestBody)
	if err := c.ShouldBindJSON(requestBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeRepository := repository.NewEmployee()
	employeeUseCase := usecase.NewEmployee(employeeRepository)

	employeeValues := values.Employee{
		Id:        employeeId,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
	}

	employeeResult, err := employeeUseCase.UpdateEmployee(&employeeValues)
	if err != nil {
		log.Println(err)

		if errors.Is(err, errormessages.ErrEmployeeNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee": employeeResult,
	})
}

func DeleteEmployee(c *gin.Context) {
	employeeId, err := strconv.Atoi(c.Param("employeeId"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	employeeRepository := repository.NewEmployee()
	employeeUseCase := usecase.NewEmployee(employeeRepository)

	if err = employeeUseCase.DeleteEmployee(employeeId); err != nil {
		log.Println(err)

		if errors.Is(err, errormessages.ErrEmployeeNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted":    "Ok",
		"employeeId": employeeId,
	})
}
