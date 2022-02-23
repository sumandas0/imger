package mock

import (
	"context"
	"image"

	"github.com/RexterR/imger/imger"
)

// NewImageRepository returns a mock implementation of ImageRepository
func NewImageRepository() imger.ImageRepository {
	return &httpMock{}
}

type httpMock struct{}

func (h *httpMock) Get(ctx context.Context, path string) (image.Image, string, error) {
	return image.NewRGBA(image.Rect(0, 0, 100, 50)), "jpeg", nil
}
