package api

import (
	"api.premiumcases.design/api/webhook"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/webhook/printify/product/publish", webhook.PrintifyProductPublish)
}
