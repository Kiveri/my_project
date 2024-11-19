package Order_in_memory_db

import "my_project/internal/domain/model"

type RepoOrder struct {
	positions      map[int]model.Order
	nextIdPosition int
}
