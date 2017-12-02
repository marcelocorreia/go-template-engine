package template_engine

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

type Engine interface {
	ParseTemplateFile(templateFile string, params interface{}) (string, error)
	ParseTemplateString(templateString string, params interface{})(string, error)
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
