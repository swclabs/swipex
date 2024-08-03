// Package controller implements the controller interface
package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/service/article"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IArticle interface for article controller
type IArticle interface {
	UploadArticle(c echo.Context) error
	UpdateCollectionsImage(c echo.Context) error
	GetArticleData(c echo.Context) error

	UploadHeadlineBanner(c echo.Context) error
	GetSlicesOfHeadlineBanner(c echo.Context) error
}

// Article struct implementation of IArticle
type Article struct {
	Services article.IArticle
}

// NewArticle creates a new Article object
func NewArticle(service article.IArticle) IArticle {
	return &Article{
		Services: service,
	}
}

// GetSlicesOfHeadlineBanner .
// @Description get list of headline banner
// @Tags collections
// @Accept json
// @Produce json
// @Param position query string true "position of collections"
// @Param limit query int true "limit headline of collections"
// @Success 200 {object} dtos.HeadlineBanners
// @Router /collections/headline [GET]
func (p *Article) GetSlicesOfHeadlineBanner(c echo.Context) error {
	var (
		pos    = c.QueryParam("position")
		sLimit = c.QueryParam("limit")
	)
	limit, err := strconv.Atoi(sLimit)
	if pos == "" || err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "position and limit are required. limit must be a number",
		})
	}
	headlines, err := p.Services.SliceOfHeadlineBanner(c.Request().Context(), pos, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, headlines)
}

// UploadHeadlineBanner .
// @Description create headline banner into collections
// @Tags collections
// @Accept json
// @Produce json
// @Param banner body dtos.HeadlineBanner true "headline banner data request"
// @Success 201 {object} dtos.OK
// @Router /collections/headline [POST]
func (p *Article) UploadHeadlineBanner(c echo.Context) error {
	var banner dtos.HeadlineBanner
	if err := c.Bind(&banner); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&banner); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.Services.UploadHeadlineBanner(c.Request().Context(), banner); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "your headline has been created successfully",
	})
}

// UploadArticle .
// @Description create collections
// @Tags collections
// @Accept json
// @Produce json
// @Param collection body dtos.UploadArticle true "collections Request"
// @Success 201 {object} dtos.CollectionUpload
// @Router /collections [POST]
func (p *Article) UploadArticle(c echo.Context) error {
	var cardBanner dtos.UploadArticle
	if err := c.Bind(&cardBanner); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&cardBanner); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	id, err := p.Services.UploadArticle(c.Request().Context(), cardBanner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.CollectionUpload{
		Msg: "collection uploaded successfully",
		ID:  id,
	})
}

// UpdateCollectionsImage .
// @Description create collections
// @Tags collections
// @Accept json
// @Produce json
// @Param img formData file true "image of collections"
// @Param id formData string true "collections identifier"
// @Success 200 {object} dtos.OK
// @Router /collections/img [PUT]
func (p *Article) UpdateCollectionsImage(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	id := c.FormValue("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'id' field",
		})
	}

	if err := p.Services.UploadCollectionsImage(c.Request().Context(), id, file); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Error{
		Msg: "upload image of collection successfully",
	})
}

// GetArticleData .
// @Description create collections
// @Tags collections
// @Accept json
// @Produce json
// @Param position query string true "position of collections"
// @Param limit query number true "limit of cards carousel"
// @Success 200 {object} dtos.Article
// @Router /collections [GET]
func (p *Article) GetArticleData(c echo.Context) error {
	position := c.QueryParam("position")
	limit := c.QueryParam("limit")
	if position == "" || limit == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'position' or 'limit' field",
		})
	}

	_limit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}

	carousel, err := p.Services.GetCarousels(c.Request().Context(), position, _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, *carousel)
}