package v1

import (
	_ "github.com/Hellsker/Todo-List/docs"
	"github.com/Hellsker/Todo-List/internal/logger"
	"github.com/Hellsker/Todo-List/internal/service"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
// Swagger spec:
// @title       ToDo API
// @description Todo API example
// @version     1.0
// @contact.name Marat Khasbiullin
// @contact.email marat.xasbiullin@gmail.com
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, logger logger.Interface, service service.TaskInterface) {
	// Options
	cfg := sloggin.Config{
		WithUserAgent:      false,
		WithRequestID:      false,
		WithRequestBody:    false,
		WithRequestHeader:  false,
		WithResponseBody:   false,
		WithResponseHeader: false,
		WithSpanID:         true,
		WithTraceID:        true,
	}
	handler.Use(sloggin.NewWithConfig(logger.GetLogger(), cfg))
	handler.Use(gin.Recovery())
	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		NewTaskRoutes(h, logger, service)
	}
}
