package effect

import "github.com/RexterR/imger/imger"

type effect struct {
	id          string
	description string
	parameters  imger.Parameters
}

func (e *effect) ID() string {
	return e.id
}
