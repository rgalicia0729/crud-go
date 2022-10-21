package repository

import (
	"database/sql"
	"github.com/rgalicia0729/crud-go/src/domain/entity"
	"github.com/rgalicia0729/crud-go/src/errormessages"
	"github.com/rgalicia0729/crud-go/src/infrastructure/db"
	"log"
)

type Employee interface {
	FindAllEmployees() ([]*entity.Employee, error)
	FindEmployeeById(employeeId int) (*entity.Employee, error)
	CreateEmployee(employee *entity.Employee) (*entity.Employee, error)
	UpdateEmployee(employee *entity.Employee) (*entity.Employee, error)
	DeleteEmployee(employeeId int) error
}

type employee struct {
	db *sql.DB
}

func NewEmployee() *employee {
	return &employee{
		db: db.PostgresPool(),
	}
}

func (e *employee) FindAllEmployees() ([]*entity.Employee, error) {
	const sqlQuery = `
		SELECT
			id,
			first_name,
			last_name,
			email,
			created_at,
			updated_at
		FROM employees
		WHERE status = true
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

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var employees []*entity.Employee
	for rows.Next() {
		var employee entity.Employee

		var updatedAtNull sql.NullTime
		err := rows.Scan(
			&employee.Id,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
			&employee.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		employee.UpdatedAt = updatedAtNull.Time

		employees = append(employees, &employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *employee) FindEmployeeById(employeeId int) (*entity.Employee, error) {
	const sqlQuery = `
		SELECT
			id,
			first_name,
			last_name,
			email,
			created_at,
			updated_at
		FROM employees
		WHERE status = true AND id = $1
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

	var employee entity.Employee
	var updatedAtNull sql.NullTime
	err = stmt.QueryRow(
		employeeId,
	).Scan(
		&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, errormessages.ErrEmployeeNotFound
	}

	employee.UpdatedAt = updatedAtNull.Time

	return &employee, nil
}

func (e *employee) CreateEmployee(employee *entity.Employee) (*entity.Employee, error) {
	const sqlQuery = `
		INSERT INTO employees(first_name, last_name, email)
		VALUES($1, $2, $3)
		RETURNING id, created_at, updated_at
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
		&employee.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}

	employee.UpdatedAt = updatedAtNull.Time

	return employee, nil
}

func (e *employee) UpdateEmployee(employee *entity.Employee) (*entity.Employee, error) {
	const sqlQuery = `
		UPDATE employees
		SET
		    first_name = $1,
		    last_name = $2,
		    email = $3,
		    updated_at = NOW()
		WHERE id = $4
		RETURNING created_at, updated_at
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
		employee.Id,
	).Scan(
		&employee.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, errormessages.ErrEmployeeNotFound
	}

	employee.UpdatedAt = updatedAtNull.Time

	return employee, nil
}

func (e *employee) DeleteEmployee(employeeId int) error {
	const sqlQuery = `
		UPDATE employees
		SET status = false, updated_at = NOW()
		WHERE status = true AND id = $1
	`

	stmt, err := e.db.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Println(err)
		}
	}()

	result, err := stmt.Exec(employeeId)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errormessages.ErrEmployeeNotFound
	}

	return nil
}
