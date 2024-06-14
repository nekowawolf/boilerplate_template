package main

import (
	"log"

	"github.com/nekowawolf/ws-andikamf/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/nekowawolf/ws-andikamf/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/nekowawolf/ws-andikamf/docs"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/nekowawolf
// @contact.email andikamusk@gmail.com

// @host  ws-andikamf-8a9eea7e7e50.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
