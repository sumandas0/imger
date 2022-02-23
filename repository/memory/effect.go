package memory

import (
	"fmt"

	"github.com/RexterR/imger/effect"
	"github.com/RexterR/imger/errors"
	"github.com/RexterR/imger/imger"
)

type effectRepository struct {
	effects []imger.Effect
}

// NewImageRepository creates a memory repository for Effect entity
func NewImageRepository(imgRepo imger.ImageRepository) imger.EffectRepository {
	return &effectRepository{
		effects: []imger.Effect{
			effect.NewRotate(),
			effect.NewResize(),
			effect.NewOverlay(imgRepo),
			effect.NewBlur(),
			effect.NewBrightness(),
			effect.NewGamma(),
			effect.NewContrast(),
			effect.NewCrop(),
		},
	}
}

func (r *effectRepository) GetEffects() ([]imger.Effect, error) {
	return r.effects, nil
}

func (r *effectRepository) GetEffect(id string) (imger.Effect, error) {
	for _, effect := range r.effects {
		if effect.ID() == id {
			return effect, nil
		}
	}

	return nil, errors.ENotExists(fmt.Sprintf("Effect %s does not exist", id), nil)
}
