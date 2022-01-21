package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/domain/engin"
	"github.com/lits01/xiaozhan/pkg/configs"
	"log"
)

var (
	configFile = "app.ini"
)

type Server struct {
	logr   *logr.Logger
	gin    *gin.Engine
	engine *engin.Engin
	config configs.Configuration
}

func NewServer(log *logr.Logger, config configs.Configuration, engin *engin.Engin) *Server {
	return &Server{
		logr:   log,
		config: config,
		gin:    gin.Default(),
		engine: engin,
	}
}

func NewConfig() configs.Configuration {
	config, err := configs.LoadConfig(configs.ConfigFilePath(configFile))
	if err != nil {
		log.Panicln(err)
	}
	return config
}

func (s *Server) Run() error {
	go s.engine.Run()
	s.gin.Run(":80")
	return nil
}
