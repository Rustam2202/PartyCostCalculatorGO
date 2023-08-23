package http

import (
	//"context"
	"context"
	"fmt"
	"net/http"
	"sync"

	//"time"

	"party-calc/docs"
	"party-calc/internal/logger"
	"party-calc/internal/server/http/handlers/calculation"
	"party-calc/internal/server/http/handlers/events"
	"party-calc/internal/server/http/handlers/persons"
	personsevents "party-calc/internal/server/http/handlers/persons_events"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	cfg               *ServerHTTPConfig
	personHandler     persons.PersonHandler
	eventHandler      events.EventHandler
	persEventsHandler personsevents.PersEventsHandler
	calcHandler       calculation.CalcHandler
	HttpServer        *http.Server
}

func NewServer(
	cfg ServerHTTPConfig,
	handlers *HTTPHandlers,
) *Server {
	return &Server{
		cfg:               &cfg,
		personHandler:     *handlers.PersonHandler,
		eventHandler:      *handlers.EventHandler,
		persEventsHandler: *handlers.PersEventsHandler,
		calcHandler:       *handlers.CalcHandler,
	}
}

// @title			Party Cost Calculator API
// @version		1.0
// @description	This is a sample app server.
// @BasePath		/
func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) { //
	//defer wg.Done()
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	{
		router.POST("/person", s.personHandler.Add)
		router.GET("/person/:id", s.personHandler.Get)
		router.PUT("/person", s.personHandler.Update)
		router.DELETE("/person/:id", s.personHandler.Delete)

		router.POST("/event", s.eventHandler.Add)
		router.GET("/event/:id", s.eventHandler.Get)
		router.PUT("/event", s.eventHandler.Update)
		router.DELETE("/event/:id", s.eventHandler.Delete)

		router.POST("/persEvents", s.persEventsHandler.Add)
		router.GET("/persEvents/:event_id", s.persEventsHandler.Get)
		router.PUT("/persEvents", s.persEventsHandler.Update)
		router.DELETE("/persEvents/:id", s.persEventsHandler.Delete)

		router.GET("/calcEvent", s.calcHandler.GetEvent)

		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	s.HttpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: router,
	}

	go func() {
		defer wg.Done()
		logger.Logger.Info("Starting HTTP server ...")
		err := s.HttpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start HTTP server", zap.Error(err))
		}
	}()
	<-ctx.Done()
	logger.Logger.Info("Shutting down HTTP server ...")
	s.HttpServer.Shutdown(context.Background())

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	err := s.httpServer.ListenAndServe()
	// 	if err != nil && err != http.ErrServerClosed {
	// 		logger.Logger.Error("HTTP server error:", zap.Error(err))
	// 	}
	// }()

	//<-ctx.Done()
	// shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()
	// shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// logger.Logger.Info("Shutting down HTTP server ...")
	// err = s.httpServer.Shutdown(shutdownCtx)
	// if err != nil {
	// 	logger.Logger.Error("HTTP server shutdown error:", zap.Error(err))
	// }

	// err := router.Run(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
	// if err != nil {
	// 	logger.Logger.Error("Server couldn`t start:", zap.Error(err))
	// 	return
	// }
}
