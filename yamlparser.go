package yamlparser

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ParseToStruct(file string, out any) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}

func ParseToMap(file string) (map[string]any, error) {
	var out map[string]any
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}
