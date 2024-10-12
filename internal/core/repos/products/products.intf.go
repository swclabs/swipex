package products

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/enum"
	"swclabs/swix/internal/core/domain/model"
)

// IProducts defines methods to interact with product (Products) data.
type IProducts interface {
	// Insert adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// prd is a pointer to the Products object to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	Insert(ctx context.Context, prd entity.Products) (int64, error)

	// GetLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductSchema objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit int, offset int) ([]entity.Products, error)

	// UploadNewImage updates the image URL of a specified product.
	// ctx is the context to manage the request's lifecycle.
	// urlImg is the new image URL to be uploaded.
	// id is the ID of the product to be updated.
	// Returns an error if any issues occur during the update process.
	UploadNewImage(ctx context.Context, urlImg string, ID int) error

	UploadShopImage(ctx context.Context, urlImg string, ID int) error

	// GetByID retrieves a product by its ID.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve.
	// Returns a pointer to the Products object and an error if any issues occur during the retrieval process.
	GetByID(ctx context.Context, productID int64) (*entity.Products, error)

	// DeleteByID deletes a product by its ID.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the product to delete.
	// Returns an error if any issues occur during the deletion process.
	DeleteByID(ctx context.Context, ID int64) error

	// Update updates a product's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// product contains the updated product details.
	// Returns an error if any issues occur during the update process.
	Update(ctx context.Context, product entity.Products) error

	// Search retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of Products objects and an error if any issues occur during the retrieval process.
	Search(ctx context.Context, keyword string) ([]entity.Products, error)

	// GetByCategory retrieves a list of products based on a specified category.
	// ctx is the context to manage the request's lifecycle.
	// types is the category type.
	// offset is the number of products to skip.
	GetByCategory(ctx context.Context, types enum.Category, offset int) ([]model.ProductXCategory, error)
}
