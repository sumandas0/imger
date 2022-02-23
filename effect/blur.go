package effect

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
	"github.com/disintegration/imaging"
)

type blur struct {
	effect
}

// NewBlur creates an Effect that blurrs an image
func NewBlur() imger.Effect {
	return &blur{
		effect: effect{
			id:          "blur",
			description: "Blur - Gaussian Blur",
			parameters: imger.Parameters{
				"sigma": imger.Parameter{
					Description: "How much the image will be blurred.",
					Required:    true,
					Example:     0.5,
					Type:        "float",
				},
			},
		},
	}
}
func (r *blur) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {

	sigma, err := floatBinder("sigma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Blur(img, sigma)

	return img, ctx.Err()
}
