package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/vanya/backend/internal/config"
	handler "github.com/vanya/backend/internal/handler/http"
	"github.com/vanya/backend/internal/service"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(
		&service.Services{},
		&geoip2.Reader{},
		&logrus.Logger{},
	)

	require.IsType(t, &handler.Handler{}, h)
}

func TestNewHandler_InitRoutes(t *testing.T) {
	h := handler.NewHandler(
		&service.Services{},
		&geoip2.Reader{},
		&logrus.Logger{},
	)
	router := h.InitRoutes(&config.Config{})

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}
