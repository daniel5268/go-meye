package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/daniel5268/go-meye/src/config"
	"github.com/daniel5268/go-meye/src/handler"
	"github.com/daniel5268/go-meye/src/infrastructure"
	"github.com/daniel5268/go-meye/src/repository"
	"github.com/daniel5268/go-meye/src/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const apiName = "gomeye"

type userDependencies struct {
	handler    handler.UserHandler
	repository repository.UserRepository
}

type dependencies struct {
	user userDependencies
}

type App struct {
	Server        *echo.Echo
	ServerReady   chan bool
	ServerStopped chan bool
	DB            *gorm.DB
	dependencies  *dependencies
}

// NewApp initializes the app
func NewApp() *App {
	a := &App{}
	a.setupInfrastructure()
	a.setupDependencies()
	a.setupServer()
	return a
}

func (a *App) setupInfrastructure() {
	a.DB = infrastructure.NewGormPostgresClient()
}

func (a *App) setupDependencies() {
	ur := repository.NewUserRepository(a.DB)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)
	ud := userDependencies{
		handler:    uh,
		repository: ur,
	}
	a.dependencies = &dependencies{
		user: ud,
	}
}

func (a *App) setupServer() {
	a.Server = echo.New()
	a.Server.HTTPErrorHandler = ErrorHandler
	a.Server.Validator = NewValidator()
	baseGroup := a.Server.Group(fmt.Sprintf("/api/%s", apiName))
	a.setupRoutes(baseGroup)
}

// StartApp initializes the server
func (a *App) StartApp() {
	go a.startServer()
	if a.ServerReady != nil {
		a.ServerReady <- true
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.stopApp()
}

func (a *App) startServer() {
	if err := a.Server.Start(fmt.Sprintf(":%s", config.Port)); err != nil {
		log.Print("Shutting down the server")
	}
}

func (a *App) stopApp() {
	ctx := context.Background()
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Print("Error shutting down the server", "error:", err)
	}
}
