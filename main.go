package main

import (
	"log"
	"os"

	"github.com/pedrobertao/gotemplater/generator"
)

func main() {
	log.Println("Generating files from JSON structure...")
	outputPath := "./"
	filename := "./generator/templates/web-api/structure.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", filename, err)
	}
	if err := generator.GenerateFromJSON(data, outputPath); err != nil {
		log.Fatalf("Error generating files: %v\n", err)
	}
	log.Println("\nFiles generated successfully!")
	log.Printf("Check the '%s' directory for generated files.\n", outputPath)

}
