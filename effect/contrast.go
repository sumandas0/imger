package effect

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
	"github.com/disintegration/imaging"
)

type contrast struct {
	effect
}

func NewContrast() imger.Effect {
	return &contrast{
		effect: effect{
			id:          "contrast",
			description: "Contrast - Change the image contrast",
			parameters: imger.Parameters{
				"percentage": imger.Parameter{
					Description: "Percentage of the contrast.",
					Required:    true,
					Example:     10,
					Type:        "float",
				},
			},
		},
	}
}

func (r *contrast) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	percentage, err := floatBinder("percentage", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustContrast(img, percentage)

	return img, ctx.Err()
}
