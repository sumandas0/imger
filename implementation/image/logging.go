package image

import (
	"context"
	"image"
	"time"

	"github.com/RexterR/imger/pkg/log"

	"github.com/RexterR/imger/imger"
)

type logService struct {
	logger  log.Logger
	service imger.ImageService
}

// NewLogService creates a log wrapper around ImageService
func NewLogService(logger log.Logger, service imger.ImageService) imger.ImageService {
	return &logService{
		logger:  logger,
		service: service,
	}
}

func (ls *logService) Process(ctx context.Context, imgSrc string, filters []imger.Filter) (image.Image, string, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"imgSrc": imgSrc, "time": time.Since(start)}, "ImageService:Process")
	}(time.Now())

	return ls.service.Process(ctx, imgSrc, filters)
}

func (ls *logService) Effects() ([]imger.Effect, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"time": time.Since(start)}, "ImageService:Effects")
	}(time.Now())

	return ls.service.Effects()
}

func (ls *logService) Effect(id string) (imger.Effect, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(log.Fields{"id": id, "time": time.Since(start)}, "ImageService:Effect")
	}(time.Now())

	return ls.service.Effect(id)
}
