package employee_usecase

import (
	"fmt"
	"my_project/internal/domain/model"
	"time"
)

type AddEmployeeReq struct {
	NameEmployee    string
	SurnameEmployee string
	PhoneEmployee   string
	EmailEmployee   string
	EmployeeRole    model.EmployeeRole
}

func (ue *UseCaseEmployee) AddEmployee(req AddEmployeeReq) (*model.Employee, error) {
	now := time.Now()
	employee := model.NewEmployee(req.NameEmployee, req.SurnameEmployee, req.PhoneEmployee, req.EmailEmployee, req.EmployeeRole, now)
	employee, err := ue.employeesRepo.CreateEmployee(employee)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.Create: %w", err)
	}
	return employee, nil
}
