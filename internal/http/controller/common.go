package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/helper/oauth2"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/pkg/utils"
)

// HealthCheck .
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/healthcheck [GET]
func HealthCheck(c echo.Context) error {
	common := service.NewCommonService()
	return c.JSON(200, common.HealthCheck())
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /oauth2/login [GET]
func Auth0Login(c echo.Context) error {
	auth := oauth2.New()
	url := auth.AuthCodeURL(auth.State)
	if err := utils.SaveSession(c, utils.BaseSessions, "state", auth.State); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Auth0Callback(c echo.Context) error {
	auth := oauth2.New()
	return auth.OAuth2CallBack(c)
}

func WorkerCheck(c echo.Context) error {
	common := service.NewCommonService()
	if err := common.DelayWorkerCheck(); err != nil {
		return c.JSON(400, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(200, common.HealthCheck())
}

func Foo(ctx echo.Context) error {
	// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
	// in the middleware before, it will be sent through as well
	panic("y tho")
}