// cmd/orbitdb-store/main.go
package main

import (
	"flag"
	"log"
	"os"

	"orbitdb-store/internal/orbitdb-store"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	input := flag.String("input", "", "Input file path")
	output := flag.String("output", "", "Output file path")
	flag.Parse()

	app := orbitdb-store.NewApp(*verbose)
	if err := app.Run(*input, *output); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
