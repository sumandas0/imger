package http

import "github.com/RexterR/imger/imger"

type effectModel struct {
	ID          string           `json:"id"`
	Description string           `json:"description"`
	Parameters  imger.Parameters `json:"parameters"`
}

func newEffectModel(e imger.Effect) effectModel {
	return effectModel{
		ID:          e.ID(),
		Description: e.Description(),
		Parameters:  e.Parameters(),
	}
}
