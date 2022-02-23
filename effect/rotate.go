package effect

import (
	"context"
	"image"
)

type rotate struct {
	effect
}

func (r *rotate) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {

	return img, ctx.Err()
}
