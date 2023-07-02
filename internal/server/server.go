package server

import (
	"fmt"

	"party-calc/internal/logger"
	"party-calc/internal/server/handlers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	cfg               *ServerConfig
	personHandler     handlers.PersonHandler
	eventHandler      handlers.EventHandler
	persEventsHandler handlers.PersEventsHandler
	calcHandler       handlers.CalcHandler
}

func NewServer(
	cfg ServerConfig,
	ph *handlers.PersonHandler,
	eh *handlers.EventHandler,
	peh *handlers.PersEventsHandler,
	ch *handlers.CalcHandler,
) *Server {
	return &Server{
		cfg:               &cfg,
		personHandler:     *ph,
		eventHandler:      *eh,
		persEventsHandler: *peh,
		calcHandler:       *ch,
	}
}

func (s *Server) Start() {
	router := gin.Default()

	router.POST("/person", s.personHandler.Add)
	router.GET("/person", s.personHandler.Get)
	router.PUT("/person", s.personHandler.Update)
	router.DELETE("/person", s.personHandler.Delete)

	router.POST("/event", s.eventHandler.Add)
	router.GET("/event", s.eventHandler.Get)
	router.PUT("/event", s.eventHandler.Update)
	router.DELETE("/event", s.eventHandler.Delete)

	router.POST("/persEvents", s.persEventsHandler.Add)
	router.GET("/persEvents", s.persEventsHandler.Get)
	router.PUT("/persEvents", s.persEventsHandler.Update)
	router.DELETE("/persEvents", s.persEventsHandler.Delete)

	router.GET("/calcEvent", s.calcHandler.GetEvent)

	err := router.Run(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
	if err != nil {
		logger.Logger.Error("Server couldn`t start:", zap.Error(err))
		return
	}
}
