package cache

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"time"

	"github.com/RexterR/imger/imger"
)

// Image caches images by the given URL and filters
type Image interface {
	Get(url string, filters []imger.Filter) (image.Image, string, error)
	Set(url string, filters []imger.Filter, format string, value image.Image) error
}

// NewImage creates a new image cache
func NewImage(cache imger.Cache) Image {
	return &imageCache{
		cache: cache,
	}
}

type imageCache struct {
	cache imger.Cache
}

func generateHash(url string, filters []imger.Filter) (string, error) {
	arrBytes := []byte{}

	arrBytes = append(arrBytes, url...)

	for _, filter := range filters {
		jsonBytes, err := json.Marshal(filter)

		if err != nil {
			return "", err
		}

		arrBytes = append(arrBytes, jsonBytes...)
	}

	hash := md5.Sum(arrBytes)

	return fmt.Sprintf("%x", hash), nil
}

func (c *imageCache) Get(url string, filters []imger.Filter) (image.Image, string, error) {
	hash, err := generateHash(url, filters)

	if err != nil {
		return nil, "", err
	}

	imgBytes, err := c.cache.Get(hash)

	if err != nil {
		return nil, "", err
	}

	r := bytes.NewReader(imgBytes)

	img, format, err := image.Decode(r)

	if err != nil {
		return nil, format, err
	}

	return img, format, err
}

func (c *imageCache) Set(url string, filters []imger.Filter, format string, img image.Image) error {
	hash, err := generateHash(url, filters)

	if err != nil {
		return err
	}

	bytes, err := imger.Encode(format, img, jpeg.DefaultQuality)

	if err != nil {
		return err
	}

	return c.cache.Set(hash, bytes, time.Minute)
}
