package collections

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/infra/db"
)

var _ ICollections = (*Collections)(nil)
var _ = app.Repos(Init)

// New creates a new Collections object
func New(conn db.IDatabase) ICollections {
	return &Collections{
		db: conn,
	}
}

// Init initializes the Collections object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) ICollections {
	return useCache(cache, New(conn))
}

// Collections struct for collections
type Collections struct {
	db db.IDatabase
}

// UploadCollectionImage implements domain.ICollections.
func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionID string, url string) error {
	return collection.db.SafeWrite(
		ctx, updateCollectionImage,
		url, collectionID,
	)
}

// Create implements domain.ICollections.
func (collection *Collections) Create(
	ctx context.Context, collectionType entity.Collection) (int64, error) {
	return collection.db.SafeWriteReturn(
		ctx, insertIntoCollections,
		collectionType.Position, collectionType.Headline, collectionType.Body,
	)
}

// GetMany implements domain.ICollections.
func (collection *Collections) GetMany(
	ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	rows, err := collection.db.Query(ctx, selectCollectionByPosition, position, limit)
	if err != nil {
		return nil, err
	}
	collections, err := db.CollectRows[entity.Collection](rows)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// AddHeadlineBanner implements domain.IHeadlineBannerCollections.
func (collection *Collections) AddHeadlineBanner(
	ctx context.Context, headline entity.Collection) error {
	return collection.db.SafeWrite(
		ctx, insertIntoCollections, headline.Position, "", headline.Body)
}
