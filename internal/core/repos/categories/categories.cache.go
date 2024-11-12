package categories

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/cache"
)

type _cache struct {
	cache    cache.ICache
	category ICategories
}

var _ ICategories = (*_cache)(nil)

func useCache(cache cache.ICache, repo ICategories) ICategories {
	return &_cache{
		category: repo,
		cache:    cache,
	}
}

// GetByID implements ICategoriesRepository.
func (c *_cache) GetByID(ctx context.Context, ID int64) (*entity.Category, error) {
	return c.category.GetByID(ctx, ID)
}

// GetLimit implements ICategoriesRepository.
func (c *_cache) GetLimit(ctx context.Context, limit string) ([]entity.Category, error) {
	return c.category.GetLimit(ctx, limit)
}

// Insert implements ICategoriesRepository.
func (c *_cache) Insert(ctx context.Context, ctg entity.Category) error {
	return c.category.Insert(ctx, ctg)
}

func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.category.DeleteByID(ctx, ID)
}

// Update implements IProductRepository.
func (c *_cache) Update(ctx context.Context, ctg entity.Category) error {
	return c.category.Update(ctx, ctg)
}
