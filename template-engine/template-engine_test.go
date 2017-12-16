package template_engine_test

import (
	"encoding/json"
	"fmt"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"os"
	"github.com/marcelocorreia/go/utils"
)

func TestParseTemplateString(t *testing.T) {
	fmt.Println("Running Test with vars...\n\n")
	engine := *getEngine()
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")
	fmt.Println("Finished Test with vars...\n")
}

func TestTemplateJson(t *testing.T) {
	engine := *getEngine()
	fmt.Println("Running Test with JSON file...")
	fmt.Println("===================================================")
	file, _ := ioutil.ReadFile("test_fixtures/bb.json")
	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	outJson, _ := engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", varsJson)
	assert.Contains(t, outJson, "Blitzkrieg Bop")
	assert.Contains(t, outJson, "The kids are losing their minds")
	assert.Contains(t, outJson, "Hey ho, let's go")
	fmt.Println(outJson)
	fmt.Println("===================================================")
	fmt.Println("Finished Test with JSON file...\n")
}

func TestTemplateErrorJson(t *testing.T) {
	fmt.Println("Running Testing throwing error...")
	engine := *getEngine()
	file, _ := ioutil.ReadFile("test_fixtures/vars.json-should-not-exist")
	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	_, err := engine.ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)
	fmt.Println("Finished Testing throwing error...\n")
}

func TestTemplateEngine_VariablesFileMerge(t *testing.T) {
	engine := *getEngine()
	out, _ := engine.VariablesFileMerge([]string{"test_fixtures/bb.yml", "test_fixtures/combo1.yml"}, getParams())
	fmt.Println(out)
	comboVars, _ := ioutil.ReadFile(out)
	fmt.Println(string(comboVars))
	out, err := engine.VariablesFileMerge([]string{"test_fixtures/bb.yml", "/a/file/that/should/not/exist"}, getParams())
	assert.Error(t,err)
	os.Remove(out)
}

func TestTemplateEngine_GetFileList(t *testing.T) {
	dir := "/go/src/github.com/marcelocorreia/go-template-engine/template-engine"

	engine := *getEngine()
	ll, _ := engine.GetFileList(dir, true, []string{})
	for _, f := range ll {
		fmt.Println(f)
	}
	_, err := engine.GetFileList("/a/dir/that/should/not/exist", true, []string{})
	assert.Error(t,err)
}

func getEngine() *template_engine.Engine {
	var engine template_engine.Engine
	engine = template_engine.TemplateEngine{}

	return &engine
}

func TestPrepareOutputDirectory(t *testing.T) {
	engine := *getEngine()
	dir := "go-template-engine/template-engine/test_fixtures/base"
	tmpDir, err := ioutil.TempDir("/tmp", "gteTest-")
	if err != nil {
		panic(err)
	}
	engine.PrepareOutputDirectory(dir, tmpDir, []string{".templates", "ci"})
	exists, err := utils.Exists(tmpDir)
	if err != nil {
		panic(err)
	}
	assert.True(t, exists)
	os.RemoveAll(tmpDir)
	exists, _ = utils.Exists(tmpDir)
	assert.False(t, exists)
	tmpDir, err = ioutil.TempDir("/bogus", "gteTest-")

	assert.Error(t, err)

	err = engine.PrepareOutputDirectory(dir, "", []string{})
	assert.Error(t, err)
}

func getParams() (map[string]string) {
	params := make(map[string]string)
	params["hey"] = "Ho"
	params["Lets"] = "go"
	return params
}

func TestTemplateEngine_LoadVars(t *testing.T) {
	engine := *getEngine()
	dir, _ := os.Getwd()
	vars, _ := engine.LoadVars(dir + "/test_fixtures/bb.yml")
	assert.NotNil(t, vars)
	vars, err := engine.LoadVars(dir + "/test_fixtures/bb-broken.yml")
	assert.Nil(t, vars)
	assert.Error(t, err)
	vars, _ = engine.LoadVars(dir + "/test_fixtures/bb.json")
	assert.NotNil(t, vars)
	vars, err = engine.LoadVars(dir + "/test_fixtures/bb-broken.json")
	assert.Nil(t, vars)
	assert.Error(t, err)
}

func TestTemplateEngine_ProcessDirectory(t *testing.T) {
	engine := *getEngine()
	dir, _ := os.Getwd()
	tmpDir := os.TempDir()
	err := engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, []string{".templates"})
	assert.Nil(t,err)
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, []string{})
	assert.Nil(t,err)
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", "/a/dir/that/should/not/exist", nil, []string{})
	assert.Error(t,err)
	err = engine.ProcessDirectory(dir+"/a/dir/that/should/not/exist", "/a/dir/that/should/not/exist", nil, nil)
	assert.Error(t,err)
}
