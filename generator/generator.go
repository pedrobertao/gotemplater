package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileStructure represents the structure for generating files
type FileStructure struct {
	Name     string                 `json:"name"`
	Type     string                 `json:"type"` // "file" or "directory"
	Content  string                 `json:"content,omitempty"`
	Children []FileStructure        `json:"children,omitempty"`
	Template string                 `json:"template,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

// GenerateFiles creates files and directories from JSON structure
func GenerateFiles(structure FileStructure, basePath string) error {
	fullPath := filepath.Join(basePath, structure.Name)

	switch structure.Type {
	case "directory":
		// Create directory
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", fullPath, err)
		}
		fmt.Printf("Created directory: %s\n", fullPath)

		// Process children
		for _, child := range structure.Children {
			if err := GenerateFiles(child, fullPath); err != nil {
				return err
			}
		}

	case "file":
		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", fullPath, err)
		}

		// Determine file content
		content := structure.Content
		if structure.Template != "" {
			content = processTemplate(structure.Template, structure.Data)
		}

		// Create file
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", fullPath, err)
		}
		fmt.Printf("Created file: %s\n", fullPath)

	default:
		return fmt.Errorf("unknown type: %s", structure.Type)
	}

	return nil
}

// processTemplate performs simple template substitution
func processTemplate(template string, data map[string]interface{}) string {
	result := template
	for key, value := range data {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}
	return result
}

// GenerateFromJSON parses JSON and generates files
func GenerateFromJSON(jsonData []byte, outputPath string) error {
	var structure FileStructure
	if err := json.Unmarshal(jsonData, &structure); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return GenerateFiles(structure, outputPath)
}

// GenerateFromJSONFile reads JSON from file and generates files
func GenerateFromJSONFile(jsonFilePath, outputPath string) error {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	return GenerateFromJSON(data, outputPath)
}
