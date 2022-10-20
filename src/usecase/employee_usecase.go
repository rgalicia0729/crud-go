package usecase

import (
	"github.com/rgalicia0729/crud-go/src/domain/entity"
	"github.com/rgalicia0729/crud-go/src/domain/values"
	"github.com/rgalicia0729/crud-go/src/infrastructure/repository"
)

type Employee struct {
	employeeRepository repository.Employee
}

func NewEmployee(employeeRepository repository.Employee) *Employee {
	return &Employee{
		employeeRepository: employeeRepository,
	}
}

func (e *Employee) FindAllEmployees() ([]*entity.Employee, error) {
	return e.employeeRepository.FindAllEmployees()
}

func (e *Employee) CreateEmployee(values *values.Employee) (*entity.Employee, error) {
	employeeEntity := entity.NewEmployee()
	employeeEntity.FirstName = values.FirstName
	employeeEntity.LastName = values.LastName
	employeeEntity.Email = values.Email

	return e.employeeRepository.CreateEmployee(employeeEntity)
}

func (e *Employee) UpdateEmployee(values *values.Employee) (*entity.Employee, error) {
	employeeEntity := entity.NewEmployee()
	employeeEntity.Id = values.Id
	employeeEntity.FirstName = values.FirstName
	employeeEntity.LastName = values.LastName
	employeeEntity.Email = values.Email

	return e.employeeRepository.UpdateEmployee(employeeEntity)
}
