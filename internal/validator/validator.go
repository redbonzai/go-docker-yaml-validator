package validator

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

type ComposeFile struct {
	Version  string                 `yaml:"version"`
	Services map[string]interface{} `yaml:"services"`
}

func ValidateComposeFile(data []byte) ([]string, error) {
	var composeFile ComposeFile
	err := yaml.Unmarshal(data, &composeFile)
	if err != nil {
		return nil, errors.New("failed to unmarshal YAML: " + err.Error())
	}

	// Convert the YAML to JSON
	jsonData, err := yamlToJSON(data)
	if err != nil {
		return nil, errors.New("failed to convert YAML to JSON: " + err.Error())
	}

	schemaLoader := gojsonschema.NewReferenceLoader("https://raw.githubusercontent.com/compose-spec/compose-spec/master/schema/compose-spec.json")
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, errors.New("failed to validate schema: " + err.Error())
	}

	if result.Valid() {
		return nil, nil
	}

	var validationErrors []string
	for _, desc := range result.Errors() {
		validationErrors = append(validationErrors, desc.String())
	}
	return validationErrors, nil
}

func yamlToJSON(data []byte) ([]byte, error) {
	var yamlData interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return nil, err
	}
	return json.Marshal(yamlData)
}
