package profile

import (
	"time"

	"github.com/RexterR/imger/pkg/log"

	"github.com/RexterR/imger/imger"
)

type logService struct {
	logger  log.Logger
	service imger.ProfileService
}

// NewLogService creates a log wrapper around ProfileService
func NewLogService(logger log.Logger, service imger.ProfileService) imger.ProfileService {
	return &logService{
		logger:  logger,
		service: service,
	}
}

func (ls *logService) GetAll(limit int64, skip int64) (*[]imger.Profile, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"limit": limit,
				"skip":  skip,
				"time":  time.Since(start)},
			"ProfileService:GetAll")
	}(time.Now())

	return ls.service.GetAll(limit, skip)
}

func (ls *logService) Get(id string) (*imger.Profile, error) {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"id":   id,
				"time": time.Since(start)},
			"ProfileService:Get")
	}(time.Now())

	return ls.service.Get(id)
}

func (ls *logService) Create(profile *imger.Profile) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"profile": profile,
				"time":    time.Since(start)},
			"ProfileService:Create")
	}(time.Now())

	return ls.service.Create(profile)
}

func (ls *logService) Update(profile *imger.Profile) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"profile": profile,
				"time":    time.Since(start)},
			"ProfileService:Update")
	}(time.Now())

	return ls.service.Update(profile)
}

func (ls *logService) Delete(id string) error {
	defer func(start time.Time) {
		ls.logger.DebugWithFields(
			log.Fields{
				"id":   id,
				"time": time.Since(start)},
			"ProfileService:Delete")
	}(time.Now())

	return ls.service.Delete(id)
}
