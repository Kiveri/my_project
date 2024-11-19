package model

import "time"

type EmployeeRole int

const (
	SpecialistEmployeeRole EmployeeRole = 1
	LeaderEmployeeRole     EmployeeRole = 2
	ManagerEmployeeRole    EmployeeRole = 3
)

type Employee struct {
	ID           int
	Name         string
	Surname      string
	Phone        string
	Email        string
	employeeRole EmployeeRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (e *Employee) AddIdentifier(id int) {
	e.ID = id
}

func (e *Employee) IsCanBuildOrder() bool {
	return e.employeeRole == ManagerEmployeeRole
}

func NewEmployee(name, surname, phone, email string, role EmployeeRole, now time.Time) *Employee {
	return &Employee{
		Name:         name,
		Surname:      surname,
		Phone:        phone,
		Email:        email,
		employeeRole: role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
