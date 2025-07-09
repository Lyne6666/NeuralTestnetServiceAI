// cmd/neuraltestnetserviceai/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"neuraltestnetserviceai/internal/neuraltestnetserviceai"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := neuraltestnetserviceai.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
