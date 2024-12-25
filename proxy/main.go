package main

import (
	"log"
)

func main() {
	geoService := NewGeoService("90a5dd26d0ba58ea94f25f085aa113ad67f2af27", "eb3066ce98823788c54dafb9e5e66d87a3c92d9d")
	if err := geoService.Serve(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
