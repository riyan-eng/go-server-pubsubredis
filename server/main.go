package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"server/config"
	"server/docs"
	"server/env"
	"server/infrastructure"
	"server/internal/repository"
	"server/internal/route"
	"server/internal/service"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
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
	infrastructure.NewLocalizer()
	config.NewValidation()
	os.Setenv("TZ", env.NewEnvironment().SERVER_TIMEZONE)
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer access token here
func main() {
	// swagger
	docs.SwaggerInfo.Title = "Golang Boilerplate One"

	// fiber
	fiberApp := fiber.New(config.NewFiberConfig())
	fiberApp.Use(cors.New(config.NewCorsConfig()))
	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())
	fiberApp.Use(infrastructure.Localizer)

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(infrastructure.Localize("WELCOME"))
	})
	fiberApp.Get("/metrics", monitor.New())
	fiberApp.Get("/docs/*", swagger.New(config.NewSwaggerConfig()))

	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.GormDB, infrastructure.Redis, config.NewEnforcer())
	exampleService := service.NewExampleService(dao)
	authenticationService := service.NewAuthenticationService(dao)
	objectService := service.NewObjectService(dao)
	route.SetupSubApp(fiberApp, exampleService, authenticationService, objectService)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		infrastructure.SqlDB.Close()
		infrastructure.SqlxDB.Close()
		infrastructure.Redis.Close()
		gorm, _ := infrastructure.GormDB.DB()
		gorm.Close()
		fiberApp.Shutdown()
		fmt.Println("Gracefully shutting down...")
	}()

	// Start the server
	if err := fiberApp.Listen(env.NewEnvironment().SERVER_HOST + ":" + env.NewEnvironment().SERVER_PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Println("Running cleanup tasks...")
}
