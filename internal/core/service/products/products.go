package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/specifications"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/lib/errors"
	"swclabs/swipecore/pkg/utils"

	"github.com/shopspring/decimal"
)

var _ IProductService = (*ProductService)(nil)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	products products.IProductRepository,
	inventory inventories.IInventoryRepository,
	category categories.ICategoriesRepository,
	Spec specifications.ISpecifications,
) IProductService {
	return &ProductService{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
		Category:  category,
		Specs:     Spec,
	}
}

// ProductService struct for product service
type ProductService struct {
	Blob      blob.IBlobStorage
	Products  products.IProductRepository
	Inventory inventories.IInventoryRepository
	Category  categories.ICategoriesRepository
	Specs     specifications.ISpecifications
}

// ViewDataOf implements IProductService.
func (s *ProductService) ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error) {
	products, err := s.Products.GetByCategory(ctx, types, offset)
	if err != nil {
		return nil, err
	}
	var (
		productView []dtos.ProductView
	)
	for _, p := range products {
		_view := dtos.ProductView{
			ID:    p.ID,
			Price: p.Price,
			Desc:  p.Description,
			Name:  p.Name,
			Image: p.Image,
		}
		if p.Specs != "" && types&enum.ElectronicDevice != 0 {
			var specs dtos.ProductSpecs
			if err := json.Unmarshal([]byte(p.Specs), &specs); err != nil {
				return nil, err
			}
			_view.Specs = specs
		}
		productView = append(productView, _view)
	}
	return productView, nil
}

// GetInventoryByID implements IProductService.
func (s *ProductService) GetInventoryByID(ctx context.Context, inventoryID int64) (*dtos.Inventory, error) {
	stock, err := s.Inventory.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, stock.ProductID)
	if err != nil {
		return nil, err
	}
	var (
		result = dtos.Inventory{
			ID:           stock.ID,
			ProductName:  product.Name,
			ProductID:    strconv.Itoa(int(stock.ProductID)),
			Price:        stock.Price.String(),
			Available:    stock.Available,
			CurrencyCode: stock.CurrencyCode,
			Status:       stock.Status,
			Color:        stock.Color,
			ColorImg:     stock.ColorImg,
			Image:        strings.Split(stock.Image, ","),
			Specs:        nil,
		}
		invID, _ = strconv.ParseInt(result.ProductID, 10, 64)
		specs    []dtos.InventorySpecification
	)
	specOfproduct, err := s.Specs.GetByInventoryID(ctx, invID)
	if err != nil {
		return nil, err
	}
	for _, spec := range specOfproduct {
		var _spec dtos.InventorySpecification
		if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
			return nil, err
		}
		specs = append(specs, _spec)
	}
	result.Specs = specs
	return &result, nil

}

// ProductDetailOf implements IProductService.
func (s *ProductService) ProductDetailOf(ctx context.Context, productID int64) (*dtos.ProductDetail, error) {
	var (
		stocks       []dtos.Inventory
		productSpecs dtos.ProductSpecs
		details      dtos.ProductDetail
	)

	rawStocks, err := s.Inventory.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	rawProduct, err := s.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(rawProduct.Specs), &productSpecs); err != nil {
		return nil, err
	}

	details.Name = rawProduct.Name
	details.Screen = productSpecs.Screen
	details.Display = productSpecs.Display
	details.Image = strings.Split(rawProduct.Image, ",")

	for _, stock := range rawStocks {
		var (
			inventory = dtos.Inventory{
				ID:           stock.ID,
				ProductName:  rawProduct.Name,
				ProductID:    strconv.Itoa(int(stock.ProductID)),
				Price:        stock.Price.String(),
				Available:    stock.Available,
				CurrencyCode: stock.CurrencyCode,
				Status:       stock.Status,
				Color:        stock.Color,
				ColorImg:     stock.ColorImg,
				Image:        strings.Split(stock.Image, ","),
				Specs:        nil,
			}
			invID, _ = strconv.ParseInt(stock.ID, 10, 64)
			specs    []dtos.InventorySpecification
		)
		specOfproduct, err := s.Specs.GetByInventoryID(ctx, invID)
		if err != nil {
			return nil, err
		}
		for _, spec := range specOfproduct {
			var _spec dtos.InventorySpecification
			if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
				return nil, err
			}
			specs = append(specs, _spec)
		}
		inventory.Specs = specs
		stocks = append(stocks, inventory)
	}

	for _, stock := range stocks {
		var (
			detailColor = dtos.DetailColor{
				Name:    stock.Color,
				Img:     stock.ColorImg,
				Product: stock.Image,
				Specs:   nil,
			}
			detailSpec []dtos.DetailSpecs
		)
		for _, spec := range stock.Specs {
			detailSpec = append(detailSpec, dtos.DetailSpecs{
				RAM:   spec.RAM,
				SSD:   spec.SSD,
				Price: stock.Price,
			})
		}
		detailColor.Specs = append(detailColor.Specs, detailSpec...)
		details.Color = append(details.Color, detailColor)
	}

	return &details, nil
}

// UpdateInventory implements IProductService.
func (s *ProductService) UpdateInventory(ctx context.Context, inventory dtos.UpdateInventory) error {
	pid, _ := strconv.Atoi(inventory.ProductID)
	price, _ := decimal.NewFromString(inventory.Price)
	return s.Inventory.Update(ctx, entity.Inventories{
		Price:        price,
		ProductID:    int64(pid),
		ID:           inventory.ID,
		Status:       inventory.Status,
		Available:    inventory.Available,
		CurrencyCode: inventory.CurrencyCode,
	})
}

// UploadStockImage implements IProductService.
func (s *ProductService) UploadStockImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err == nil {
			if err = s.Inventory.UploadImage(ctx, ID, resp.SecureURL); err == nil {
				if err = file.Close(); err != nil {
					return err
				}
			}
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteInventoryByID implements IProductService.
func (s *ProductService) DeleteInventoryByID(ctx context.Context, inventoryID int64) error {
	return s.Inventory.DeleteByID(ctx, inventoryID)
}

// GetAllStock implements IProductService.
func (s *ProductService) GetAllStock(ctx context.Context, page int, limit int) (*dtos.StockInInventory, error) {
	inventories, err := s.Inventory.GetLimit(ctx, limit, page)
	if err != nil {
		return nil, errors.Service("get stock", err)
	}
	var stock dtos.StockInInventory

	for _, _inventory := range inventories {
		switch _inventory.Status {
		case "active":
			stock.Header.Active++
		case "draft":
			stock.Header.Draft++
		case "archived":
			stock.Header.Active++
		}
		var (
			specs    []dtos.InventorySpecification
			invID, _ = strconv.ParseInt(_inventory.ID, 10, 64)
		)
		specOfproduct, err := s.Specs.GetByInventoryID(ctx, invID)
		if err != nil {
			return nil, err
		}
		for _, spec := range specOfproduct {
			var _spec dtos.InventorySpecification
			if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
				return nil, err
			}
			specs = append(specs, _spec)
		}
		product, err := s.Products.GetByID(ctx, _inventory.ProductID)
		if err != nil {
			return nil, err
		}
		stock.Stock = append(stock.Stock, dtos.Inventory{
			Specs:        specs,
			ProductName:  product.Name,
			ProductID:    strconv.Itoa(int(_inventory.ProductID)),
			Image:        strings.Split(_inventory.Image, ","),
			ID:           _inventory.ID,
			Price:        _inventory.Price.String(),
			Available:    _inventory.Available,
			CurrencyCode: _inventory.CurrencyCode,
			Status:       _inventory.Status,
			ColorImg:     _inventory.ColorImg,
			Color:        _inventory.Color,
		})
	}

	stock.Page = page
	stock.Limit = limit
	stock.Header.All = len(inventories)

	return &stock, nil
}

// GetInventory implements IProductService.
func (s *ProductService) GetInventory(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	return s.Inventory.GetByProductID(ctx, productID)
}

// Search implements IProductService.
func (s *ProductService) Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error) {
	_products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var productSchema []dtos.ProductResponse
	for _, p := range _products {
		productSchema = append(productSchema, dtos.ProductResponse{
			ID:          p.ID,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Image:       strings.Split(p.Image, ","),
			Created:     utils.HanoiTimezone(p.Created),
		})
	}
	return productSchema, nil
}

// UpdateProductInfo implements IProductService.
func (s *ProductService) UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error {
	ID, _ := strconv.Atoi(product.CategoryID)
	_category, err := s.Category.GetByID(ctx, int64(ID))
	if err != nil {
		return fmt.Errorf("category not found %v", err)
	}
	var types enum.Category
	if err := types.Load(_category.Name); err != nil {
		return fmt.Errorf("category invalid %v", err)
	}
	_product := entity.Products{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		SupplierID:  product.SupplierID,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
	}
	return s.Products.Update(ctx, _product)

}

// UploadProductImage implements IProductService.
func (s *ProductService) UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err != nil {
			return err
		}
		if err := s.Products.UploadNewImage(ctx, resp.SecureURL, ID); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

// CreateProduct implements IProductService.
func (s *ProductService) CreateProduct(ctx context.Context, products dtos.Product) (int64, error) {
	ID, _ := strconv.Atoi(products.CategoryID)
	_category, err := s.Category.GetByID(ctx, int64(ID))
	if err != nil {
		return -1, fmt.Errorf("category not found %v", err)
	}

	var types enum.Category
	if err := types.Load(_category.Name); err != nil {
		return -1, fmt.Errorf("category invalid %v", err)
	}
	var prd = entity.Products{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Specs:       "{}",
	}
	if products.Specs != nil && types&enum.ElectronicDevice != 0 {
		var specs, ok = products.Specs.(dtos.ProductSpecs)
		if !ok {
			return -1, fmt.Errorf("invalid specifications")
		}
		specsByte, _ := json.Marshal(specs)
		prd.Specs = string(specsByte)
	}
	return s.Products.Insert(ctx, prd)
}

// DeleteProductByID implements IProductService.
func (s *ProductService) DeleteProductByID(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}

// InsertIntoInventory implements IProductService.
func (s *ProductService) InsertIntoInventory(ctx context.Context, product dtos.Inventory) error {
	var (
		pid, _    = strconv.Atoi(product.ProductID)
		price, _  = decimal.NewFromString(product.Price)
		inventory = entity.Inventories{
			Color:        product.Color,
			ColorImg:     product.ColorImg,
			Image:        strings.Join(product.Image, ","),
			ProductID:    int64(pid),
			Price:        price,
			Available:    product.Available,
			CurrencyCode: product.CurrencyCode,
			Status:       "active",
		}
		types enum.Category
	)
	pID, _ := strconv.ParseInt(product.ProductID, 10, 64)
	p, err := s.Products.GetByID(ctx, pID)
	if err != nil {
		return err
	}
	cID, _ := strconv.ParseInt(p.CategoryID, 10, 64)
	category, err := s.Category.GetByID(ctx, cID)
	if err != nil {
		return err
	}
	if err := types.Load(category.Name); err != nil {
		return err
	}
	invID, err := s.Inventory.InsertProduct(ctx, inventory)
	if err != nil {
		return err
	}
	if types&enum.ElectronicDevice != 0 {
		for _, spec := range product.Specs {
			bSpec, _ := json.Marshal(spec)
			if err := s.Specs.Insert(ctx, entity.Specifications{
				InventoryID: invID,
				Content:     string(bSpec),
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetProductsLimit implements IProductService.
func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]dtos.ProductResponse, error) {
	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var productResponse []dtos.ProductResponse
	for _, p := range products {
		var (
			product = dtos.ProductResponse{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     utils.HanoiTimezone(p.Created),
				Image:       strings.Split(p.Image, ","),
			}
			types         enum.Category
			categoryID, _ = strconv.ParseInt(p.CategoryID, 10, 64)
		)
		category, err := s.Category.GetByID(ctx, categoryID)
		if err != nil {
			return nil, err
		}
		if err := types.Load(category.Name); err != nil {
			return nil, err
		}
		productResponse = append(productResponse, product)
	}
	return productResponse, nil
}
