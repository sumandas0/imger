package profile

import (
	"github.com/RexterR/imger/imger"
)

type service struct {
	repository imger.ProfileRepository
}

// NewService returns a profile service
func NewService(repository imger.ProfileRepository) imger.ProfileService {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(limit int64, skip int64) (*[]imger.Profile, error) {
	return s.repository.GetAll(limit, skip)
}

func (s *service) Get(id string) (*imger.Profile, error) {
	return s.repository.Get(id)
}

func (s *service) Create(profile *imger.Profile) error {
	return s.repository.Create(profile)
}

func (s *service) Update(profile *imger.Profile) error {
	return s.repository.Update(profile)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
