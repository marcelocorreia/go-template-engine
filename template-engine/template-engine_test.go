package template_engine_test

import (
	"encoding/json"
	"fmt"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParseTemplateString(t *testing.T) {

	fmt.Println("Running Test with vars...\n\n")

	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := template_engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")

	fmt.Println("Finished Test with vars...\n")

}

func TestTemplateJson(t *testing.T) {

	fmt.Println("Running Test with JSON file...")
	fmt.Println("===================================================")

	file, _ := ioutil.ReadFile("test_fixtures/bb.json")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	outJson, _ := template_engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", varsJson)
	assert.Contains(t, outJson, "Blitzkrieg Bop")
	assert.Contains(t, outJson, "The kids are losing their minds")
	assert.Contains(t, outJson, "Hey ho, let's go")

	fmt.Println(outJson)
	fmt.Println("===================================================")
	fmt.Println("Finished Test with JSON file...\n")

}
func TestTemplateErrorJson(t *testing.T) {

	fmt.Println("Running Testing throwing error...")

	file, _ := ioutil.ReadFile("test_fixtures/vars.json-should-not-exist")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)

	_, err := template_engine.ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)

	fmt.Println("Finished Testing throwing error...\n")

}
