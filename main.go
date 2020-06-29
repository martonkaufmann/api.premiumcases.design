package main

import (
	"os"

	"api.premiumcases.design/api"
	"api.premiumcases.design/pkg/requestvalidator"
	"github.com/bugsnag/bugsnag-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: os.Getenv("BUGSNAG_API_KEY"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{os.Getenv("BUGSNAG_PROJECT")},
	})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &requestvalidator.RequestValidator{Validator: validator.New()}

	// Routes
	api.RegisterRoutes(e.Group("api"))

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
