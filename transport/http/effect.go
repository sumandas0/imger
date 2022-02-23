package http

import (
	"net/http"

	"github.com/RexterR/imger/imger"
	"github.com/julienschmidt/httprouter"
)

type effectsController struct {
	service imger.ImageService
}

func newEffectsController(service imger.ImageService) *effectsController {
	return &effectsController{
		service: service,
	}
}

func (c *effectsController) get(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	effect, err := c.service.Effect(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, newEffectModel(effect))
}

func (c *effectsController) getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	effects, err := c.service.Effects()

	if err != nil {
		return errResponse(err)
	}

	var effectDesc = make([]effectModel, 0, len(effects))

	for _, e := range effects {
		effectDesc = append(effectDesc, newEffectModel(e))
	}

	return response(http.StatusOK, effectDesc)
}
