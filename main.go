package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/pedrobertao/gotemplater/core/generator"
)

func main() {
	project := flag.String("name", "myapp", "Name of the project to generate")
	template := flag.String("template", "api-simple", "Template to use (must exist in /templates)")
	flag.Parse()

	templatePath := filepath.Join("templates", *template, "structure.yaml")

	if err := generator.Generate(*project, templatePath); err != nil {
		log.Fatalf("Error generating project: %v", err)
	}

	fmt.Printf("✅ Project '%s' generated using '%s' template!\n", *project, *template)
}
