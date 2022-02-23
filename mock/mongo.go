package mock

import (
	"sync"
	"time"

	"github.com/RexterR/imger/imger"
	"github.com/RexterR/imger/pkg/errors"
)

type profileRepo struct {
	mutex      *sync.Mutex
	repository []imger.Profile
}

// NewProfileRepository returns a mock implemation of ProfileRepository interface
func NewProfileRepository() imger.ProfileRepository {
	return &profileRepo{
		mutex:      &sync.Mutex{},
		repository: []imger.Profile{},
	}
}

func (r *profileRepo) GetAll(limit int64, skip int64) (*[]imger.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	p := r.repository[skip : limit+skip]

	return &p, nil
}

func (r *profileRepo) Get(id string) (*imger.Profile, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, v := range r.repository {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.ENotExists("Profile does not exists", nil)
}

func (r *profileRepo) Create(profile *imger.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	profile.Created = time.Now().UTC()
	profile.Updated = time.Now().UTC()

	r.repository = append(r.repository, *profile)

	return nil
}

func (r *profileRepo) Update(profile *imger.Profile) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, v := range r.repository {
		if profile.ID == v.ID {
			profile.Updated = time.Now().UTC()
			r.repository[i] = *profile
			return nil
		}
	}

	return errors.ENotExists("Profile does not exist", nil)
}

func (r *profileRepo) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, v := range r.repository {
		if v.ID == id {
			r.repository = append(r.repository[:i], r.repository[i+1:]...)
			return nil
		}
	}

	return errors.ENotExists("Profile does not exist", nil)
}
