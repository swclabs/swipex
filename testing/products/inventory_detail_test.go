package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/inventories"
	"swclabs/swix/pkg/lib/logger"
	"testing"

	productContainer "swclabs/swix/internal/apis/container/products"
	productRepo "swclabs/swix/internal/core/repos/products"
	productService "swclabs/swix/internal/core/service/products"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var e = echo.New()

func TestGetInventory(t *testing.T) {
	var (
		inventory  inventories.Mock
		product    productRepo.Mock
		category   categories.Mock
		service    = productService.New(nil, &product, &inventory, &category)
		controller = productContainer.NewController(service)
	)
	specs, _ := json.Marshal(dtos.Specs{
		SSD: "256",
		RAM: "128",
	})
	category.On("GetByID", context.Background(), int64(1)).Return(&entity.Categories{
		ID:          1,
		Name:        "phone",
		Description: "iPhone",
	}, nil)

	inventory.On("GetByID", context.Background(), int64(1)).Return(&entity.Inventories{
		ID:           1,
		ProductID:    1,
		Available:    1000,
		Price:        decimal.NewFromInt(10000),
		CurrencyCode: "VND",
		Status:       "active",
		Color:        "Black Titanium",
		ColorImg:     "https://example.com/black-titanium.jpg",
		Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
		Specs:        string(specs),
	}, nil)

	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Products{
		Name:       "iPhone 12",
		CategoryID: 1,
		ID:         1,
	}, nil)

	e.GET("/inventories/details", controller.GetInvDetails)
	req := httptest.NewRequest(http.MethodGet, "/inventories/details?id=1", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body dtos.Inventory
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fail()
	}

	file, err := os.Create("./inventory_detail_out.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	logger := logger.Write(file)
	logger.Info("Response body", zap.Any("body", body), zap.Int("status", rr.Code))
}
