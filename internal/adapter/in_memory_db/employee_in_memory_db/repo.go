package employee_in_memory_db

import "my_project/internal/domain/model"

type Repo struct {
	employees map[int]*model.Employee
	nextID    int
}

func NewRepo() *Repo {
	return &Repo{
		nextID:    1,
		employees: make(map[int]*model.Employee),
	}
}

func (r *Repo) getNextEmployeeID() int {
	nextID := r.nextID
	r.nextID++

	return nextID
}
