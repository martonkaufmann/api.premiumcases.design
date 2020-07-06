package main

import (
	"net/url"
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
		ProjectPackages:     []string{os.Getenv("BUGSNAG_PROJECT")},
		NotifyReleaseStages: []string{"production"},
		ReleaseStage:        os.Getenv("ENV"),
	})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Validator
	e.Validator = &requestvalidator.RequestValidator{Validator: validator.New()}

	// Routes
	api.RegisterRoutes(e)

	// Proxy imaginary requests
	url, err := url.Parse(os.Getenv("IMAGINARY_HOST"))
	if err != nil {
		e.Logger.Fatal(err)
	}

	targets := []*middleware.ProxyTarget{
		{
			URL: url,
		},
	}

	imagesGroup := e.Group("/images")
	imagesGroup.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
