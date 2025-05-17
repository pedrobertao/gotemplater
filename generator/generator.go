package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// TemplatePath represents a file path within a template definition
type TemplatePath struct {
	Path string `yaml:"path"`
}

// Template represents a complete template definition with its files
type Template struct {
	Files []TemplatePath `yaml:"files"`
}

// Generate creates a new project based on the specified template
func Generate(projectName, templateName string) error {
	// Read and parse the template file
	tmpl, err := readTemplate(templateName)
	if err != nil {
		return err
	}

	// Generate all files defined in the template
	return generateFiles(projectName, templateName, tmpl)
}

// readTemplate reads and parses a template file from the embedded filesystem
func readTemplate(templateName string) (*Template, error) {
	data, err := Templates.ReadFile(templateName)
	if err != nil {
		return nil, fmt.Errorf("failed to read template '%s': %w", templateName, err)
	}

	var tmpl Template
	if err := yaml.Unmarshal(data, &tmpl); err != nil {
		return nil, fmt.Errorf("failed to parse YAML for template '%s': %w", templateName, err)
	}

	return &tmpl, nil
}

// generateFiles creates all files defined in the template
func generateFiles(projectName, templateName string, tmpl *Template) error {
	for _, file := range tmpl.Files {
		if err := generateFile(projectName, templateName, file); err != nil {
			return err
		}
	}
	return nil
}

// generateFile creates a single file with its directory structure
func generateFile(projectName, templateName string, file TemplatePath) error {
	fullPath := filepath.Join(projectName, file.Path)

	// Create all necessary directories
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", filepath.Dir(fullPath), err)
	}

	// Generate content based on template and file path
	content := generateContent(templateName, file.Path)

	// Write content to file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file '%s': %w", fullPath, err)
	}

	return nil
}
