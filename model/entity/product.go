package entity

import "time"

type Product struct {
	ProductId     string
	Name          string
	Description   string
	Price         int
	StockQuantity int
	ProductImages string
	UserId        string
	CategoryId    string
	CreatedAt     string
	UpdatedAt     string
}
