package app

import (
	"context"
	"fmt"
	"github.com/Hellsker/Todo-List/config"
	v1 "github.com/Hellsker/Todo-List/internal/controller/http/v1"
	"github.com/Hellsker/Todo-List/internal/logger"
	repository "github.com/Hellsker/Todo-List/internal/repository/task"
	service "github.com/Hellsker/Todo-List/internal/service/task"
	"github.com/Hellsker/Todo-List/pkg/postgres"
	"github.com/Hellsker/Todo-List/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	//init logger
	logger := logger.New(cfg.Env)
	logger.Debug("Logger initialized")

	//init db
	dbConf := postgres.NewConfig(cfg)
	pg := postgres.New(dbConf)
	defer pg.Close()
	if pg == nil {
		logger.Error("Database init error!")
		return
	}
	err := pg.Ping(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	logger.Debug("Database initialization was successful!")
	//init repo
	repoTask := repository.New(pg)
	logger.Debug("Repository initialization was successful")
	// init Service
	taskServ := service.New(repoTask)
	logger.Debug("Service initialization was successful")
	// init HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, logger, taskServ)
	httpServer := server.New(handler)
	logger.Debug("Server initialization was successful")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}
