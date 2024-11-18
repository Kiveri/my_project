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
	EmployeeRole EmployeeRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func NewEmployee(name, surname, phone, email string, role EmployeeRole, now time.Time) *Employee {
	return &Employee{
		Name:         name,
		Surname:      surname,
		Phone:        phone,
		Email:        email,
		EmployeeRole: role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
