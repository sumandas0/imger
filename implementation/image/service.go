package image

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
)

// NewService creates an ImageService
func NewService(imageRepo imger.ImageRepository, effectRepo imger.EffectRepository) imger.ImageService {
	return &service{
		imageRepo:  imageRepo,
		effectRepo: effectRepo,
	}
}

type service struct {
	imageRepo  imger.ImageRepository
	effectRepo imger.EffectRepository
}

func (s *service) Process(ctx context.Context, imgSrc string, filters []imger.Filter) (image.Image, string, error) {
	img, imgType, err := s.imageRepo.Get(ctx, imgSrc)

	if err != nil {
		return nil, imgType, err
	}

	for _, filter := range filters {
		effect, err := s.effectRepo.GetEffect(filter.ID)

		if err != nil {
			return nil, imgType, err
		}

		img, err = effect.Transform(ctx, img, filter.Parameters)

		if err != nil {
			return nil, imgType, err
		}
	}

	return img, imgType, ctx.Err()
}

func (s *service) Effects() ([]imger.Effect, error) {
	return s.effectRepo.GetEffects()
}

func (s *service) Effect(id string) (imger.Effect, error) {
	return s.effectRepo.GetEffect(id)
}
