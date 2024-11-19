package employee_in_memory_db

import "my_project/internal/domain/model"

type RepoEmployees struct {
	employees      map[int]*model.Employee
	nextIdEmployee int
}

func NewRepoEmployees() *RepoEmployees {
	return &RepoEmployees{
		nextIdEmployee: 1,
		employees:      make(map[int]*model.Employee),
	}
}

func (re *RepoEmployees) getNextEmployeeID() int {
	nextIdEmployee := re.nextIdEmployee
	re.nextIdEmployee++

	return nextIdEmployee
}
