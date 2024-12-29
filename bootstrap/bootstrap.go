package bootstrap

import (
	"bookstore/adapter/database"
	controller "bookstore/adapter/http/book"
	routes "bookstore/adapter/http/routes"
	service "bookstore/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type bootstrap struct {
	routes *routes.Routes
}

var externalModule = fx.Module("external",
	fx.Provide(gin.Default))

var portModule = fx.Module("adapter",
	fx.Provide(routes.NewRoute,
		routes.NewBookRoute,
		controller.NewBookController,
		database.NewDatabaseAdapter))

var serviceModule = fx.Module("service",
	fx.Provide(service.NewBookServiceFactory,
		service.NewBookService))

var Module = fx.Options(
	externalModule,
	portModule,
	serviceModule,
	fx.Invoke(newBootstrap))

func newBootstrap(lifecycle fx.Lifecycle, routes *routes.Routes, httpServer *gin.Engine) {
	bs := &bootstrap{
		routes: routes,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go bs.startUpHttpServer(httpServer, routes)
			log.Println("HTTP server started successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("HTTP server ended successfully")
			return nil
		}})
}

func (bs *bootstrap) startUpHttpServer(httpServer *gin.Engine, routes *routes.Routes) {
	routes.Setup()
	httpServer.Run()
}
