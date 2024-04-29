package service

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/pkg/cloud"
)

type Posts struct {
	category   domain.ICategoriesRepository
	product    domain.IProductRepository
	newsletter domain.INewsletterRepository
}

func NewPost() domain.IPostsService {
	return &Posts{
		category:   repository.NewCategories(),
		product:    repository.NewProducts(),
		newsletter: repository.NewNewsletter(),
	}
}

// GetHomeBanner implements domain.IPostsService.
func (p *Posts) GetHomeBanner(ctx context.Context, limit int) ([]domain.HomeBanners, error) {
	return p.newsletter.GetHomeBanner(ctx, limit)
}

// GetNewsletter implements domain.IPostsService.
func (p *Posts) GetNewsletter(ctx context.Context, limit int) ([]domain.Newsletters, error) {
	return p.newsletter.Get(ctx, limit)
}

// UploadHomeBanner implements domain.IPostsService.
func (p *Posts) UploadHomeBanner(ctx context.Context, data domain.HomeBanners, fileHeader *multipart.FileHeader) error {
	url, err := cloud.UploadFile(ctx, cloud.Connection(), fileHeader)
	if err != nil {
		return err
	}
	data.Img = url
	return p.newsletter.InsertHomeBanner(ctx, data)
}

// UploadNewsletter implements domain.IPostsService.
func (p *Posts) UploadNewsletter(ctx context.Context, news domain.Newsletter, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UploadImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	news.Image = resp.SecureURL
	return repository.NewNewsletter().Insert(context.Background(), news)
}