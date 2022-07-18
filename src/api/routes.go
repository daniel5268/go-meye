package api

import "github.com/labstack/echo/v4"

func (a *App) setupRoutes(g *echo.Group) {
	userHandler := a.dependencies.user.handler
	v1Group := g.Group("/v1")
	v1UserGroup := v1Group.Group("/users")

	g.GET("/health-check", func(ctx echo.Context) error {
		return ctx.NoContent(200)
	})

	v1UserGroup.POST("/token", userHandler.SignIn)
	v1UserGroup.POST("", userHandler.Create, AuthAdmin(a.dependencies.user.repository))
}
