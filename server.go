package main

import (
	"./routes"

	"github.com/zenazn/goji"
)

func main() {
	// Serve static files
	goji.Use(Static("public"))

	// Add routes
	routes.Include()

	// Run Goji
	goji.Serve()
}
