package image

import (
	"context"
	"image"

	"github.com/RexterR/imger/cache"
	"github.com/RexterR/imger/imger"
)

type cacheService struct {
	cache   cache.Image
	service imger.ImageService
}

// NewCacheService creates a cache wrapper around ImageService
func NewCacheService(cache cache.Image, service imger.ImageService) imger.ImageService {
	return &cacheService{
		cache:   cache,
		service: service,
	}
}

func (cs *cacheService) Process(ctx context.Context, imgSrc string, filters []imger.Filter) (image.Image, string, error) {
	img, format, err := cs.cache.Get(imgSrc, filters)

	if err == nil {
		return img, format, nil
	}

	img, format, err = cs.service.Process(ctx, imgSrc, filters)

	if err == nil {
		err := cs.cache.Set(imgSrc, filters, format, img)

		if err != nil {
			return img, format, err
		}
	}

	return img, format, err
}

func (cs *cacheService) Effects() ([]imger.Effect, error) {
	return cs.service.Effects()
}

func (cs *cacheService) Effect(id string) (imger.Effect, error) {
	return cs.service.Effect(id)
}
