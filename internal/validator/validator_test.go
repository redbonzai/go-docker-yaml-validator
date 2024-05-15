package validator

import (
	"fmt"
	"os"
	"testing"
)

// Helper function to read test files
func readFile(t *testing.T, filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading test file %s: %v", filePath, err)
	}
	fmt.Printf("compose data: %v\n", data)
	return data
}

func TestValidateComposeFile(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		expectErr bool
	}{
		{"ValidComposeFile", "../../testdata/valid-compose.yml", false},
		{"InvalidComposeFile", "../../testdata/invalid-compose.yml", true},
		{"InvalidSchemaComposeFile", "../../testdata/invalid-schema-compose.yml", true},
		{"MinimalValidComposeFile", "../../testdata/minimal-valid-compose.yml", false},
		{"AllOptionalFieldsComposeFile", "../../testdata/all-optional-fields-compose.yml", false},
		{"UnsupportedFieldsComposeFile", "../../testdata/unsupported-fields-compose.yml", true},
		{"LargeComposeFile", "../../testdata/large-compose.yml", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readFile(t, tt.filePath)
			errors, err := ValidateComposeFile(data)
			if tt.expectErr {
				if err == nil && len(errors) == 0 {
					t.Errorf("Expected validation errors, got none")
				} else {
					t.Logf("Received expected errors: %v, %v", err, errors)
				}
			} else {
				if err != nil || len(errors) != 0 {
					t.Errorf("Expected no validation errors, got %v, %v", err, errors)
				}
			}
		})
	}
}
