package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
)

type contrast struct {
	effect
}

func (r *contrast) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustContrast(img, percentage)

	return img, ctx.Err()
}
