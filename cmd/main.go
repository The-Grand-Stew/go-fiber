package main

import "os"

func main() {

	// Define a new Fiber app with config.
	app := fiber.New()

	routes.PublicRoutes(app) // Register a public routes for app.
	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "prod" {
		utils.StartServerWithGracefulShutdown(app)
	} else {
		utils.StartServer(app)
	}
}
