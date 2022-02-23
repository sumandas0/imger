package http

import (
	"github.com/RexterR/imger/imger"
	"github.com/RexterR/imger/pkg/errors"
)

type createProfileModel struct {
	ID      string         `json:"id"`
	Filters []imger.Filter `json:"filters"`
}

type updateProfileModel struct {
	Filters []imger.Filter `json:"filters"`
}

func (m *createProfileModel) toProfile() (*imger.Profile, error) {
	if m.ID == "" {
		return nil, errors.EValidation("id is missing", nil)
	}

	if len(m.Filters) == 0 {
		return nil, errors.EValidation("filters are empty", nil)
	}

	return &imger.Profile{ID: m.ID, Filters: m.Filters}, nil
}

func (m *updateProfileModel) toProfile(profile *imger.Profile) (*imger.Profile, error) {
	if len(m.Filters) == 0 {
		return nil, errors.EValidation("effects are empty", nil)
	}

	profile.Filters = m.Filters

	return profile, nil
}
