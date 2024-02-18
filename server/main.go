package main

import (
	"log"
	"os"
	"runtime"
	"server/config"
	"server/docs"
	"server/env"
	"server/infrastructure"
	"server/internal/repository"
	"server/internal/route"
	"server/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	env.LoadEnvironmentFile()
	env.NewEnvironment()
	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
	infrastructure.ConnectGormDB()
	infrastructure.ConnRedis()
	os.Setenv("TZ", env.SERVER_TIMEZONE)
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer access token here
func main() {
	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.GormDB, infrastructure.Redis, config.NewEnforcer())
	exampleService := service.NewExampleService(dao)
	authenticationService := service.NewAuthenticationService(dao)
	objectService := service.NewObjectService(dao)

	// swagger
	docs.SwaggerInfo.Title = "Golang Boilerplate One"

	// fiber
	fiberApp := fiber.New(config.NewFiberConfig())
	fiberApp.Use(cors.New(config.NewCorsConfig()))
	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())
	fiberApp.Get("/", func(c *fiber.Ctx) error { return c.SendString("Welcome to Golang Boilerplate One APIs") })
	fiberApp.Get("/metrics", monitor.New())
	fiberApp.Get("/docs/*", swagger.New(config.NewSwaggerConfig()))
	route.NewRoute(fiberApp, exampleService, authenticationService, objectService)
	if err := fiberApp.Listen(env.SERVER_HOST + ":" + env.SERVER_PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
