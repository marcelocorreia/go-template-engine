package machine_io

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func JsonOutput(obj interface{}) (string, error) {
	out, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return "", nil
	}
	return string(out), nil
}

func YamlOutput(obj interface{}) (string, error) {
	out, err := yaml.Marshal(obj)
	if err != nil {
		return "", nil
	}
	return string(out), nil
}