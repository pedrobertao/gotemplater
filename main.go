package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pedrobertao/gotemplater/generator"
)

type AvailableTemplates struct {
	Name         string
	TemplatePath string
}

var options = []AvailableTemplates{
	{
		Name:         "Simple Web API",
		TemplatePath: "api-simple",
	},
	{
		Name:         "Web API with Mongo",
		TemplatePath: "api-mongo",
	},
	{
		Name:         "Web API with Postgress",
		TemplatePath: "api-pgsql",
	},
	{
		Name:         "Web API SQLite",
		TemplatePath: "api-sqlite",
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📦 Project name (myapp): ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading project name %v \n", err)
	}
	projectName = strings.TrimSpace(projectName)

	if projectName == "" {
		projectName = "myapp"
	}

	fmt.Println("\n📁 Available templates:")
	for i, tmpl := range options {
		fmt.Printf("  %d. %s\n", i+1, tmpl.Name)
	}

	optionChose := AvailableTemplates{}
	for {
		fmt.Print("\nChoose your template by number: ")
		tmtlName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading template option %v \n", err)
		}
		tmtlName = strings.TrimSpace(tmtlName)

		t, exists := templateExists(tmtlName)
		if !exists {
			fmt.Printf("\nInvalid template chose, try again.")
			continue
		}
		optionChose = *t
		break
	}

	templatePath := fmt.Sprintf("templates/%s/structure.yaml", optionChose.TemplatePath)

	if err := generator.Generate(projectName, templatePath); err != nil {
		log.Fatalf("error generating project: %v", err)
	}

	fmt.Printf("✅ Project '%s' generated using '%s' template!\n", projectName, optionChose.Name)
}

func templateExists(tempPosi string) (*AvailableTemplates, bool) {
	posi, err := strconv.Atoi(tempPosi)
	if err != nil {
		return nil, false
	}

	if posi > len(options) || posi <= 0 {
		return nil, false
	}

	return &options[posi-1], true
}
