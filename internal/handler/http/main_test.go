package http

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vanya/backend/internal/config"
)

var cfg *config.HandlerTest

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	testCfg := config.InitTestConfig("../../../configs")

	cfg = new(config.HandlerTest)
	*cfg = *testCfg

	os.Exit(m.Run())
}
