package template_engine

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestTemplateVars(t *testing.T) {

	fmt.Println("Running Test with vars...\n\n")

	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := ParseTemplateFile("test_fixtures/README.md.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")

	fmt.Println("Finished Test with vars...\n")

}

func TestTemplateJson(t *testing.T) {

	fmt.Println("Running Test with JSON file...")

	file, _ := ioutil.ReadFile("test_fixtures/vars.json")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	outJson, _ := ParseTemplateFile("test_fixtures/README.md.tpl", varsJson)
	//assert.Contains(t, out, []string{"# Blitzkrieg Bop","Hey ho, let's go"})
	assert.Contains(t, outJson, "Hey ho, let's go")

	fmt.Println("Finished Test with JSON file...\n")

}
func TestTemplateErrorJson(t *testing.T) {

	fmt.Println("Running Testing throwing error...")

	file, _ := ioutil.ReadFile("test_fixtures/vars.json-should-not-exist")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)

	_, err := ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)

	fmt.Println("Finished Testing throwing error...\n")

}
