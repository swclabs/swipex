// Package router manager implementation
// File manager.go defines routes for APIs related to manager
// such as login, signup, logout, and user information.
package router

import (
	"swclabs/swix/internal/webapi/controller"
	"swclabs/swix/internal/webapi/middleware"

	"github.com/labstack/echo/v4"
)

// IManager interface for manager
type IManager interface {
	IRouter
}

// Manager struct	implementation of IManager
type Manager struct {
	controller controller.IManager
}

// NewManager creates a new Manager router object
func NewManager(controllers controller.IManager) IManager {
	return &Manager{
		controller: controllers,
	}
}

// Routers define route endpoint
func (account *Manager) Routers(e *echo.Echo) {
	// endpoint for users
	e.GET("/users", account.controller.GetMe, middleware.SessionProtected)
	e.PUT("/users", account.controller.UpdateUserInfo)
	e.PUT("/users/image", account.controller.UpdateUserImage, middleware.SessionProtected)

	// endpoint for authentication
	e.POST("/auth", account.controller.Auth)
	e.GET("/auth/email", account.controller.CheckLoginEmail)
	e.POST("/auth/signup", account.controller.SignUp)
	e.POST("/auth/login", account.controller.Login)
	e.GET("/auth/logout", account.controller.Logout)

	// endpoint for oauth2 service
	e.GET("/callback", controller.Auth0Callback)
	e.GET("/oauth2/login", controller.Auth0Login)
}
