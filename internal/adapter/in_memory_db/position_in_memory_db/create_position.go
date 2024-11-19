package position_in_memory_db

import "my_project/internal/domain/model"

func (rp *RepoPositions) CreatePosition(position *model.Position) (*model.Position, error) {
	position.AddIdPosition(rp.getNextPositionId())
	rp.positions[position.IdPosition] = position

	return position, nil
}
