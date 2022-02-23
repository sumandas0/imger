package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
)

type gamma struct {
	effect
}

func (r *gamma) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	gamma, err := floatBinder("gamma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustGamma(img, gamma)

	return img, ctx.Err()
}
