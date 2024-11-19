package employee_usecase

import "my_project/internal/adapter/in_memory_db/employee_in_memory_db"

type UseCase struct {
	employeeRepo *employee_in_memory_db.Repo
}

func NewUseCase(employeeRepo *employee_in_memory_db.Repo) *UseCase {
	return &UseCase{
		employeeRepo: employeeRepo,
	}
}
