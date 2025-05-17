package generator

import (
	"fmt"
	"strings"
)

// toCamel converts snake_case to CamelCase.
func toCamel(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func templatePathInFS(templateName string) string {
	return fmt.Sprintf("templates/%s/structure.yaml", templateName)
}
