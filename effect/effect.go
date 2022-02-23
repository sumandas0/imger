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

func (e *effect) Description() string {
	return e.description
}

func (e *effect) Parameters() imger.Parameters {
	return e.parameters
}
