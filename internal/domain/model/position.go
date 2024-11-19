package model

import "time"

type PositionType int

const (
	BasicProductPositionType    PositionType = 1
	BasicConsumablePositionType PositionType = 2
	LiquidPositionType          PositionType = 3
	OversizePositionType        PositionType = 4
)

type Position struct {
	IdPosition        int
	NamePosition      string
	BarcodePosition   string
	PricePosition     float32
	InStockPosition   int
	CreatedAtPosition time.Time
	UpdatedAtPosition time.Time
	DeletedAtPosition *time.Time
	positionType      PositionType
}

func (p *Position) AddIdPosition(idPos int) {
	p.IdPosition = idPos
}

func (e *Employee) IsCanAddPosition() bool {

	return e.employeeRole == ManagerEmployeeRole
}

func NewPosition(name, barcode string, price float32, posType PositionType, now time.Time) *Position {
	return &Position{
		NamePosition:      name,
		BarcodePosition:   barcode,
		PricePosition:     price,
		positionType:      posType,
		CreatedAtPosition: now,
		UpdatedAtPosition: now,
	}
}
