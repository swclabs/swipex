package orders

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
)

// IOrders interface for orders repos
type IOrders interface {
	Create(ctx context.Context, order entity.Order) (int64, error)
	Get(ctx context.Context, userID int64, limit int) ([]entity.Order, error)
	GetByUUID(ctx context.Context, orderCode string) (*entity.Order, error)
	GetItemByCode(ctx context.Context, orderCode string) ([]model.Order, error)
	InsertProduct(ctx context.Context, product entity.ProductInOrder) error
	GetProductByOrderID(ctx context.Context, orderID int64) ([]entity.ProductInOrder, error)
}
