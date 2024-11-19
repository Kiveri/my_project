package employee_in_memory_db

import "my_project/internal/domain/model"

func (re *RepoEmployees) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.AddIdEmployee(re.getNextEmployeeID())
	re.employees[employee.IdEmployee] = employee

	return employee, nil
}
