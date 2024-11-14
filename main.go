package main

import (
	"fmt"
	"git-project-management/config"
	"git-project-management/internal/app"
	"log"

	"gitea.com/logicamp/lc"
	"github.com/danielgtaylor/huma/v2"
	humaFiber "github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig, _ := lc.GetConfig[config.Config](&config.Config{})
	fiberApp := fiber.New()
	config := huma.DefaultConfig("Git Project Management", "1.0.0")
	config.Servers = []*huma.Server{{URL: appConfig.BASE_URL}}
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"auth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	api := humaFiber.New(fiberApp, config)
	app.Setup(&api, *appConfig)

	log.Fatal(fiberApp.Listen(fmt.Sprintf(":%s", appConfig.PORT)))
}
