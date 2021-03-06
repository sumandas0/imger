package effect

import (
	"context"
	"image"
	"testing"

	"github.com/RexterR/imger/pkg/errors"
)

func TestRotateTransform(t *testing.T) {
	tt := []struct {
		name   string
		params map[string]interface{}
		err    errors.Type
	}{
		{
			name: "transform successfully",
			params: map[string]interface{}{
				"angle":   0.9,
				"bgcolor": "black",
			},
		},
		{
			name: "missing angle",
			params: map[string]interface{}{
				"bgcolor": "black",
			},
			err: errors.Validation,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, 100, 50))
			rotate := NewRotate()

			_, err := rotate.Transform(context.Background(), img, tc.params)

			if tc.err != "" {
				if err == nil || !errors.Is(tc.err, err) {
					t.Error("Expected validation error", err)
				}
			}
		})
	}
}
