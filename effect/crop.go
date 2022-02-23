package effect

import (
	"context"
	"image"

	"github.com/disintegration/imaging"
)

type crop struct {
	effect
}

func (o *crop) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	rectangle, err := rectangleBinder("rectangle", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Crop(img, rectangle)

	return img, ctx.Err()
}
