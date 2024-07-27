package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId     uint64     `json:"orderId"`
	CustomerId  uuid.UUID  `json:"customerId"`
	LineItems   []LineItem `json:"lineItem"`
	CreatedAt   *time.Time `json:"createdAt"`
	ShippedAt   *time.Time `json:"shippedAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"itemId"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
