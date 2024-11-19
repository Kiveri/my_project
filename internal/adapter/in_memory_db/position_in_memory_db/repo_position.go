package position_in_memory_db

import "my_project/internal/domain/model"

type RepoPositions struct {
	positions      map[int]model.Position
	nextIdPosition int
}

func NewRepoPosition() *RepoPositions {
	return &RepoPositions{
		nextIdPosition: 1,
		positions:      make(map[int]model.Position),
	}
}

func (rp *RepoPositions) getNextPositionId() int {
	nextIdPosition := rp.nextIdPosition
	rp.nextIdPosition++

	return nextIdPosition
}
