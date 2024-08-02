package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/domain/model"
	productRepo "swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/pkg/lib/logger"
	"testing"

	"go.uber.org/zap"
)

func TestProductView(t *testing.T) {
	var (
		product productRepo.Mock
		service = products.ProductService{
			Products: &product,
		}
		controller = controller.Products{
			Services: &service,
		}
		specs = []dtos.ProductSpecs{
			{
				Screen:  "6.1 inch",
				Display: "Super Retina XDR display",
				SSD:     []int{128, 256, 512},
				RAM:     []int{4, 8},
			},
			{
				Screen:  "6.7 inch",
				Display: "Super Retina XDR display",
				SSD:     []int{128, 256, 512},
				RAM:     []int{4, 8},
			},
			{
				Screen:  "5.4 inch",
				Display: "Super Retina XDR display",
				SSD:     []int{128, 256, 512},
				RAM:     []int{4, 8},
			},
		}

		sspecs []string
	)

	for _, spec := range specs {
		specs, _ := json.Marshal(spec)
		sspecs = append(sspecs, string(specs))
	}

	product.On("GetByCategory", context.Background(), enum.Phone, 0).Return(
		[]model.ProductXCategory{
			{
				ID:           1,
				Name:         "iPhone 12",
				Description:  "iPhone 12",
				Price:        "1.000.000 - 2.000.000",
				Image:        "https://example.com/iphone-12.jpg",
				CategoryName: enum.Phone.String(),
				Spec:         sspecs[0],
			},
			{
				ID:           2,
				Name:         "iPhone 12 Pro",
				Description:  "iPhone 12 Pro",
				Price:        "2.000.000 - 3.000.000",
				Image:        "https://example.com/iphone-12-pro.jpg",
				CategoryName: enum.Phone.String(),
				Spec:         sspecs[1],
			},
			{
				ID:           3,
				Name:         "iPhone 12 Mini",
				Description:  "iPhone 12 Mini",
				Price:        "1.000.000 - 2.000.000",
				Image:        "https://example.com/iphone-12-mini.jpg",
				CategoryName: enum.Phone.String(),
				Spec:         sspecs[2],
			},
		},
		nil)
	e.GET("/products/view", controller.GetProductView)
	req := httptest.NewRequest(http.MethodGet, "/products/view?type=phone", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body []dtos.ProductView
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fail()
	}

	file, err := os.Create("./products_view_out.json")
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

func TestProductViewAccessory(t *testing.T) {
	var (
		product productRepo.Mock
		service = products.ProductService{
			Products: &product,
		}
		controller = controller.Products{
			Services: &service,
		}
	)

	product.On("GetByCategory", context.Background(), enum.Accessories, 0).Return(
		[]model.ProductXCategory{
			{
				ID:           1,
				Name:         "Apple iPhone Adapter",
				Description:  "Apple iPhone Adapter 20W",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-adapter.jpg",
				Spec:         "",
				CategoryName: enum.Accessories.String(),
			},
			{
				ID:           2,
				Name:         "Apple iPhone Case",
				Description:  "Apple iPhone Case",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-case.jpg",
				Spec:         "",
				CategoryName: enum.Accessories.String(),
			},
			{
				ID:           3,
				Name:         "Apple iPhone Screen Protector",
				Description:  "Apple iPhone Screen Protector",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-screen-protector.jpg",
				Spec:         "",
				CategoryName: enum.Accessories.String(),
			},
		},
		nil)
	e.GET("/products/view", controller.GetProductView)
	req := httptest.NewRequest(http.MethodGet, "/products/view?type=accessories", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body []dtos.ProductView
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fail()
	}

	file, err := os.Create("./products_view_accessory_out.json")
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