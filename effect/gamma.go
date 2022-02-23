package effect

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
	"github.com/disintegration/imaging"
)

type gamma struct {
	effect
}

func NewGamma() imger.Effect {
	return &gamma{
		effect: effect{
			id:          "gamma",
			description: "Gamma - Change the image gamma",
			parameters: imger.Parameters{
				"gamma": imger.Parameter{
					Description: "Percentage of the gamma correction.",
					Required:    true,
					Example:     0.75,
					Type:        "float",
				},
			},
		},
	}
}

func (r *gamma) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	gamma, err := floatBinder("gamma", params)

	if err != nil {
		return nil, err
	}

	img = imaging.AdjustGamma(img, gamma)

	return img, ctx.Err()
}
