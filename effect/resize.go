package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
)

type resize struct {
	effect
}

func (r *resize) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	width, err := integerBinder("width", params)

	if err != nil {
		return nil, err
	}

	height, err := integerBinder("height", params)

	if err != nil {
		return nil, err
	}

	filter, err := filterBinder("filter", params)

	if err != nil {
		filter = imaging.Linear
	}

	img = imaging.Resize(img, width, height, filter)

	return img, ctx.Err()
}
