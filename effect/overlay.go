package effect

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
	"github.com/disintegration/imaging"
)

type overlay struct {
	imgRepository imger.ImageRepository
	effect
}

func (o *overlay) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	position, err := pointBinder("position", params)

	if err != nil {
		return nil, err
	}

	url, err := urlBinder("url", params)

	if err != nil {
		return nil, err
	}

	opacity, err := floatBinder("opacity", params)

	if err != nil {
		opacity = 100
	}

	overlayImg, _, err := o.imgRepository.Get(ctx, url.String())

	if err != nil {
		return nil, err
	}

	img = imaging.Overlay(img, overlayImg, position, opacity)

	return img, ctx.Err()
}
