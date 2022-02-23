package effect

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
	"github.com/disintegration/imaging"
)

type crop struct {
	effect
}

// NewCrop creates an Effect that crops the image
func NewCrop() imger.Effect {
	return &crop{
		effect: effect{
			id:          "crop",
			description: "Crop - Crops image",
			parameters: imger.Parameters{
				"rectangle": imger.Parameter{
					Description: "Region to crop (x0,y0,x1,y1)",
					Required:    true,
					Example:     "[0,0,100,200]",
					Type:        "array[int]",
				},
			},
		},
	}
}

func (o *crop) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	rectangle, err := rectangleBinder("rectangle", params)

	if err != nil {
		return nil, err
	}

	img = imaging.Crop(img, rectangle)

	return img, ctx.Err()
}
