package employee_usecase

import (
	"fmt"
	"my_project/internal/domain/model"
	"time"
)

type AddEmployeeReq struct {
	Name         string
	Surname      string
	Phone        string
	Email        string
	EmployeeRole model.EmployeeRole
}

func (u *UseCase) AddEmployee(req AddEmployeeReq) (*model.Employee, error) {
	now := time.Now()
	employee := model.NewEmployee(req.Name, req.Surname, req.Phone, req.Email, req.EmployeeRole, now)
	employee, err := u.employeeRepo.Create(employee)
	if err != nil {
		return nil, fmt.Errorf("employeeRepo.Create: %w", err)
	}
	return employee, nil
}
