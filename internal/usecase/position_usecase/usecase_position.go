package position_usecase

import "my_project/internal/adapter/in_memory_db/position_in_memory_db"

type UseCasePosition struct {
	positionsRepo *position_in_memory_db.RepoPositions
}

func NewUseCasePosition(positionRepo *position_in_memory_db.RepoPositions) *UseCasePosition {
	return &UseCasePosition{
		positionsRepo: positionRepo,
	}
}
