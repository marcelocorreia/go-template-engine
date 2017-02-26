package template_engine

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"encoding/json"
	"github.com/daviddengcn/go-colortext"
	"fmt"
)

func TestTemplateVars(t *testing.T) {
	ct.Foreground(ct.Cyan,false)
	fmt.Println("Running Test with vars...\n\n")
	ct.ResetColor()
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := ParseTemplateFile("test_fixtures/README.md.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")
	ct.Foreground(ct.Green,false)
	fmt.Println("Finished Test with vars...\n")
	ct.ResetColor()
}

func TestTemplateJson(t *testing.T) {
	ct.Foreground(ct.Cyan,false)
	fmt.Println("Running Test with JSON file...")
	ct.ResetColor()
	file, _ := ioutil.ReadFile("test_fixtures/vars.json")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	outJson, _:= ParseTemplateFile("test_fixtures/README.md.tpl", varsJson)
	//assert.Contains(t, out, []string{"# Blitzkrieg Bop","Hey ho, let's go"})
	assert.Contains(t, outJson, "Hey ho, let's go")
	ct.Foreground(ct.Green,false)
	fmt.Println("Finished Test with JSON file...\n")
	ct.ResetColor()
}
func TestTemplateErrorJson(t *testing.T) {
	ct.Foreground(ct.Cyan,false)
	fmt.Println("Running Testing throwing error...")
	ct.ResetColor()
	file, _ := ioutil.ReadFile("test_fixtures/vars.json-should-not-exist")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)

	_,err:= ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t,err)
	ct.Foreground(ct.Green,false)
	fmt.Println("Finished Testing throwing error...\n")
	ct.ResetColor()
}