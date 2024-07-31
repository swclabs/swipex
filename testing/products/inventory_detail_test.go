package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/pkg/lib/logger"
	"testing"

	productRepo "swclabs/swipecore/internal/core/repository/products"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var e = echo.New()

func TestGetInventory(t *testing.T) {
	var (
		specs = dtos.InventorySpecification{
			RAM: "8GB",
			SSD: "256GB",
		}

		inventory inventories.Mock
		product   productRepo.Mock
		service   = products.ProductService{
			Inventory: &inventory,
			Products:  &product,
		}
		controller = controller.Products{
			Services: &service,
		}
	)

	sspecs, _ := json.Marshal(specs)

	inventory.On("GetByID", context.Background(), int64(1)).Return(&entity.Inventories{
		ID:           "1",
		ProductID:    1,
		Available:    "1000",
		Price:        decimal.NewFromInt(10000),
		CurrencyCode: "VND",
		Status:       "active",
		Color:        "Black Titanium",
		ColorImg:     "https://example.com/black-titanium.jpg",
		Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
		Specs:        string(sspecs),
	}, nil)

	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Products{
		Name: "iPhone 12",
	}, nil)

	e.GET("/inventories/details", controller.GetInventoryDetails)
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
