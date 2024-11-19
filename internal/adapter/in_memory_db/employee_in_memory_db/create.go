package employee_in_memory_db

import "my_project/internal/domain/model"

func (r *Repo) Create(employee *model.Employee) (*model.Employee, error) {
	employee.AddIdentifier(r.getNextEmployeeID())
	r.employees[employee.ID] = employee

	return employee, nil
}
