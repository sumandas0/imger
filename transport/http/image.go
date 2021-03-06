package http

import (
	"context"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RexterR/imger/imger"
	"github.com/RexterR/imger/pkg/errors"
	"github.com/julienschmidt/httprouter"
)

type imagesController struct {
	service        imger.ImageService
	profileService imger.ProfileService
}

func newImagesController(service imger.ImageService, profile imger.ProfileService) *imagesController {
	return &imagesController{
		service:        service,
		profileService: profile,
	}
}

func getQuality(r *http.Request) int {
	h := r.Header.Get("accept")
	values := strings.Split(h, ";")

	for _, v := range values {
		if i := strings.Index(v, "q="); i > -1 {
			q, err := strconv.Atoi(v[i+2:])

			if err != nil {
				return jpeg.DefaultQuality
			}

			return q
		}
	}

	return jpeg.DefaultQuality
}

func getParameters(srv imger.ProfileService, r *http.Request) (string, []imger.Filter, error) {
	var filters []imger.Filter
	imgSrc := r.URL.Query().Get("imgSrc")
	filtersJSON := r.URL.Query().Get("filters")
	profileID := r.URL.Query().Get("profile")

	if imgSrc == "" {
		return imgSrc, filters, errors.EMalformed("Missing imgSrc query parameter", nil)
	}

	if filtersJSON != "" {
		err := json.Unmarshal([]byte(filtersJSON), &filters)

		if err != nil {
			return imgSrc, filters, errors.EMalformed("effects query parameter is malformed", err)
		}
	}

	if profileID != "" {
		profile, err := srv.Get(profileID)

		if err == nil {
			filters = append(profile.Filters, filters...)
		}
	}

	return imgSrc, filters, nil
}

func (c *imagesController) transformImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	imgSrc, filters, err := getParameters(c.profileService, r)
	if err != nil {
		return errResponse(err)
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	img, format, err := c.service.Process(ctx, imgSrc, filters)

	if err != nil {
		return errResponse(err)
	}

	q := getQuality(r)

	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))

	bytes, err := imger.Encode(format, img, q)

	if err != nil {
		return errResponse(err)
	}

	if _, err := w.Write(bytes); err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, nil)
}
