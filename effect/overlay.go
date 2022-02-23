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

func NewOverlay(imgRepository imger.ImageRepository) imger.Effect {
	return &overlay{
		imgRepository: imgRepository,
		effect: effect{
			id:          "overlay",
			description: "Overlay - Overlay image",
			parameters: imger.Parameters{
				"position": imger.Parameter{
					Description: "Position for the overlay image",
					Required:    true,
					Example:     "[1,2]",
					Type:        "array[int]",
				},
				"url": imger.Parameter{
					Description: "Url for overlay image",
					Required:    true,
					Example:     "http://image.png",
					Type:        "string",
				},
				"opacity": imger.Parameter{
					Description: "Opacity for overlay image",
					Required:    false,
					Example:     90,
					Type:        "float",
					Default:     100,
				},
			},
		},
	}
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
