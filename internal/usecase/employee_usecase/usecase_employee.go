package employee_usecase

import "my_project/internal/adapter/in_memory_db/employee_in_memory_db"

type UseCaseEmployee struct {
	employeesRepo *employee_in_memory_db.RepoEmployees
}

func NewUseCaseEmployee(employeeRepo *employee_in_memory_db.RepoEmployees) *UseCaseEmployee {
	return &UseCaseEmployee{
		employeesRepo: employeeRepo,
	}
}
