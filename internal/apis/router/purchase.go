// Package router implements the router interface
// File purchase.go defines routes for APIs related to purchasing, adding to cart,
// accessing invoices, order, and creating invoices, order, etc.
package router

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/apis/controller"

	"github.com/labstack/echo/v4"
)

var _ = boot.Router(NewPurchase)

// NewPurchase returns a new Purchase router object
func NewPurchase(controllers controller.IPurchase) IPurchase {
	return &Purchase{controllers: controllers}
}

// IPurchase extends the IRouter interface
type IPurchase interface {
	IRouter
}

// Purchase is the router implementation for IPurchase
type Purchase struct {
	controllers controller.IPurchase
}

// Routers define route endpoint
func (p *Purchase) Routers(e *echo.Echo) {
	e.POST("/purchase/carts", p.controllers.AddToCarts)
	e.GET("/purchase/carts", p.controllers.GetCarts)
	e.DELETE("/purchase/carts/:id", p.controllers.DeleteItem)

	e.GET("/purchase/orders", p.controllers.GetOrders)
	e.POST("/purchase/orders", p.controllers.CreateOrder)
}
