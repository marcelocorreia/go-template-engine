package template_engine

import (
	"io/ioutil"
	"text/template"
	"bytes"
)


func    ParseTemplateFile(templateFile string, params interface{}) (string, error) {
	tplFile, err := ioutil.ReadFile(templateFile)

	if err != nil {
		return "", err
	}

	r, err := ParseTemplateString(string(tplFile), params)
	return r, nil
}

func ParseTemplateString(templateString string, params interface{}) (string, error) {
	t := template.Must(template.New("letter").Parse(templateString))

	var doc bytes.Buffer
	errParse := t.Execute(&doc, params)
	if errParse != nil {
		return "", errParse
	}
	resp := doc.String()

	return resp, nil
}