package template_engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Engine interface {
	ParseTemplateFile(templateFile string, params interface{}) (string, error)
	ParseTemplateString(templateString string, params interface{}) (string, error)
	VariablesFileMerge(varsFile []string) (string, error)
}

type TemplateEngine struct {
}

func (gte TemplateEngine) ParseTemplateFile(templateFile string, params interface{}) (string, error) {
	tplFile, err := ioutil.ReadFile(templateFile)

	if err != nil {
		return "", err
	}

	r, err := gte.ParseTemplateString(string(tplFile), params)
	return r, nil
}

func (gte TemplateEngine) ParseTemplateString(templateString string, params interface{}) (string, error) {
	t := template.Must(template.New("letter").Parse(templateString))

	var doc bytes.Buffer
	errParse := t.Execute(&doc, params)
	if errParse != nil {
		return "", errParse
	}
	resp := doc.String()

	return resp, nil
}

func (gte TemplateEngine) VariablesFileMerge(varsFile []string) (string, error) {
	var payload string

	for _, payload := range varsFile {

		output, _ := loadVarsIntoStruct(payload)

		fmt.Println(output)
	}
	return payload, nil
}

func loadVarsIntoStruct(ffile string) (string, error) {
	var varsFile interface{}
	file, _ := ioutil.ReadFile(ffile)

	if strings.HasSuffix(ffile, ".json") {
		err := json.Unmarshal(file, &varsFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)

		}
	} else if strings.HasSuffix(ffile, ".yaml") || strings.HasSuffix(ffile, ".yml") {
		err := yaml.Unmarshal(file, &varsFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)

		}

	}

	return varsFile.(string), nil
}
