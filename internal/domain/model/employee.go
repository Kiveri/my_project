package model

import "time"

type EmployeeRole int

const (
	SpecialistEmployeeRole EmployeeRole = 1
	LeaderEmployeeRole     EmployeeRole = 2
	ManagerEmployeeRole    EmployeeRole = 3
)

type Employee struct {
	IdEmployee        int
	NameEmployee      string
	SurnameEmployee   string
	PhoneEmployee     string
	EmailEmployee     string
	employeeRole      EmployeeRole
	CreatedAtEmployee time.Time
	UpdatedAtEmployee time.Time
	DeletedAtEmployee *time.Time
}

func (e *Employee) AddIdEmployee(idEmp int) {
	e.IdEmployee = idEmp
}

func (e *Employee) IsCanBuildOrder() bool {
	return e.employeeRole == ManagerEmployeeRole
}

func NewEmployee(name, surname, phone, email string, role EmployeeRole, now time.Time) *Employee {
	return &Employee{
		NameEmployee:      name,
		SurnameEmployee:   surname,
		PhoneEmployee:     phone,
		EmailEmployee:     email,
		employeeRole:      role,
		CreatedAtEmployee: now,
		UpdatedAtEmployee: now,
	}
}
