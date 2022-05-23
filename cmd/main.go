package main

import (
	"{{ .appName }}/app/routes"
    "{{ .appName }}/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func health(c *fiber.Ctx)error{
    return c.JSON("Health Check from {{ .appName }}")

}


func main(){
    os.Setenv("SERVER_PORT","{{ .appPort }}")
   app := fiber.New() 
   app.Get("/{{.appName}}/health", health)
   {{ .routes }}
   // Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "prod" {
		utils.StartServerWithGracefulShutdown(app)
	} else {
		utils.StartServer(app)
	}
}
