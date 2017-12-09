package machine_io

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl"
)

type MachineIO interface {
	Json2Yaml(content string) (string, error)
	Yaml2Json(content string) (string, error)
	Json2Hcl(content string) (string, error)
	Hcl2Json(content string) (string, error)
	JsonOutput(obj interface{}) (string, error)
	YamlOutput(obj interface{}) (string, error)
}

type Converter struct {}

type Envelope struct {
	Type string
	Msg  interface{}
}

func (c Converter) Yaml2Hcl(content string) (string, error) {
	json,err:=c.Yaml2Json(content)
	if err != nil {
		return "",err
	}
	fmt.Println(json)
	return "",nil
}

func (c Converter) Json2Hcl(content string) (string, error) {
	input := []byte(content)

	ast, err := jsonParser.Parse([]byte(input))

	if err != nil {
		return "", fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)

	if err != nil {
		return "", fmt.Errorf("unable to print HCL: %s", err)
	}

	return "", nil
}

func (c Converter) Hcl2JSON(content string) (string, error) {
	input := []byte(content)
	var v interface{}
	err := hcl.Unmarshal(input, &v)
	json, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		return "", fmt.Errorf("unable to marshal json: %s", err)
	}

	return string(json), nil
}

func (c Converter) Json2Yaml(content string) (string, error) {
	var body interface{}
	err := json.Unmarshal([]byte(content), &body)
	if err != nil {
		return "", err
	}
	out, err := yaml.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (c Converter) Yaml2Json(content string) (string, error) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(content), &body)

	out, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (c Converter) JsonOutput(obj interface{}) (string, error) {
	out, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (c Converter) YamlOutput(obj interface{}) (string, error) {
	out, err := yaml.Marshal(obj)
	if err != nil {
		return "", nil
	}
	return string(out), nil
}
