package model

import "time"

type Position struct {
	ID        int
	Name      string
	Barcode   string
	Price     float32
	InStock   int
	CreatedAt time.Time
}

func NewPosition(name, barcode string, price float32, inStock int) *Position {
	return &Position{
		Name:      name,
		Barcode:   barcode,
		Price:     price,
		InStock:   inStock,
		CreatedAt: time.Now(),
	}
}
