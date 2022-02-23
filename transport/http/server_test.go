package http

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/talento90/go-health"

	"github.com/RexterR/imger/implementation/image"
	"github.com/RexterR/imger/implementation/profile"
	"github.com/RexterR/imger/mock"
	"github.com/RexterR/imger/pkg/log"
	"github.com/RexterR/imger/repository/memory"
)

func mockDependencies() *ServerDependencies {
	imgRepository := mock.NewImageRepository()
	effectRepo := memory.NewImageRepository(imgRepository)
	imgService := image.NewService(imgRepository, effectRepo)
	profileService := profile.NewService(mock.NewProfileRepository())
	logger, _ := log.NewLogger(log.Configuration{Output: ioutil.Discard})

	dep := &ServerDependencies{
		ImgService:     imgService,
		ProfileService: profileService,
		Logger:         logger,
		Health:         health.New("imger", health.Options{CheckersTimeout: time.Second}),
	}

	return dep
}

func createMockServer() *httptest.Server {
	dep := mockDependencies()
	handler := createRouter(dep)

	return httptest.NewServer(handler)
}

func TestNewServer(t *testing.T) {
	dep := mockDependencies()

	srv := NewServer(&Configuration{}, dep)

	if srv.Handler == nil {
		t.Error("Expect Handler to have a valid http.Hanlder")
	}
}
