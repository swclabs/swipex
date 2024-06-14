package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICartRepository implement all method of Carts To access database
type ICartRepository interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, warehouseId int64) error
	// InsertMany insert many Products to database by list of productID
	InsertMany(ctx context.Context, warehouseIds []int64) error
	// GetCartByUserID is a method get CartSchema from database by userId
	GetCartByUserID(ctx context.Context, userId int64) (*domain.CartSchema, error)
	// RemoveItem delete Products in Cart
	RemoveItem(ctx context.Context, warehouseId int64, userId int64) error
}
