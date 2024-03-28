package main

import (
	"fmt"
	"os"

	"github.com/babelcoder-dummy/intro-devops/api/config"
	"github.com/babelcoder-dummy/intro-devops/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	app := fiber.New()

	if os.Getenv("APP_ENV") == "production" {
		tracer.Start()
		defer tracer.Stop()
		
		app.Use(fibertrace.Middleware(fibertrace.WithServiceName("api")))
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	config.SetupEnv()
	config.SetupDB()
	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%d", config.Env.Port))
}
