package repository

import (
	"database/sql"
	"github.com/rgalicia0729/crud-go/src/domain/entity"
	"github.com/rgalicia0729/crud-go/src/infrastructure/db"
	"log"
)

type Employee interface {
	CreateEmployee(employee *entity.Employee) (*entity.Employee, error)
}

type employee struct {
	db *sql.DB
}

func NewEmployee() *employee {
	return &employee{
		db: db.PostgresPool(),
	}
}

func (e *employee) CreateEmployee(employee *entity.Employee) (*entity.Employee, error) {
	const sqlQuery = `
		INSERT INTO employees(first_name, last_name, email)
		VALUES($1, $2, $3)
		RETURNING id, status, created_at, updated_at
	`

	stmt, err := e.db.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}()

	var updatedAtNull sql.NullTime

	err = stmt.QueryRow(
		employee.FirstName,
		employee.LastName,
		employee.Email,
	).Scan(
		&employee.Id,
		&employee.Status,
		&employee.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}

	employee.UpdatedAt = updatedAtNull.Time

	return employee, nil
}
