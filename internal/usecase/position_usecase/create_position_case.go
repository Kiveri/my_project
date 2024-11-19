package position_usecase

import (
	"fmt"
	"my_project/internal/domain/model"
	"time"
)

type AddPositionReq struct {
	NamePosition    string
	BarcodePosition string
	PricePosition   float32
	PositionType    model.PositionType
}

func (up *UseCasePosition) AddPosition(req AddPositionReq) (*model.Position, error) {
	now := time.Now()
	position := model.NewPosition(req.NamePosition, req.BarcodePosition, req.PricePosition, req.PositionType, now)
	position, err := up.positionsRepo.CreatePosition(position)
	if err != nil {
		return nil, fmt.Errorf("positionRepo.Create: %w", err)
	}
	return position, nil
}
