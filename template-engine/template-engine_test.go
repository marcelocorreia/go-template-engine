package template_engine_test

import (
	"encoding/json"
	"fmt"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"os"
	"github.com/marcelocorreia/go-template-engine/utils"
)

var TEST_DELIMS = []string{"{{{", "}}}"}
var DEFAULT_DELIMS = []string{"{{", "}}"}

func TestParseTemplateString(t *testing.T) {
	fmt.Println("Running Test with vars...\n\n")
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")
	fmt.Println("Finished Test with vars...\n")
}

func TestTemplateJson(t *testing.T) {
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
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
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	file, _ := ioutil.ReadFile("test_fixtures/vars.json-should-not-exist")
	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	_, err := engine.ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)
	fmt.Println("Finished Testing throwing error...\n")
}

func TestTemplateEngine_VariablesFileMerge(t *testing.T) {
	engine, err := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	assert.Nil(t, err)
	out, err := engine.VariablesFileMerge([]string{"test_fixtures/bb.yml", "test_fixtures/combo1.yml"}, getParams())
	assert.Nil(t, err)
	out, err = engine.VariablesFileMerge([]string{"test_fixtures/bb.yml", "/a/file/that/should/not/exist"}, getParams())
	assert.Error(t, err)
	os.Remove(out)
	out, err = engine.VariablesFileMerge([]string{"test_fixtures/bb-broken.yml"}, getParams())
	assert.Nil(t, err)
}

func TestTemplateEngine_GetFileList(t *testing.T) {
	dir := "/go/src/github.com/marcelocorreia/go-template-engine/template-engine"

	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	ll, _ := engine.GetFileList(dir, true, []string{}, []string{})
	assert.True(t, len(ll) > 0)
	_, err := engine.GetFileList("/a/dir/that/should/not/exist", true, []string{}, []string{})
	assert.Error(t, err)
}

func TestPrepareOutputDirectory(t *testing.T) {
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
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
	params["name"] = "Willie Nelson"
	return params
}

func TestTemplateEngine_LoadVars(t *testing.T) {
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
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
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	dir, _ := os.Getwd()
	tmpDir := os.TempDir()
	err := engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, []string{".templates"},[]string{})
	assert.Nil(t, err)
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, []string{},[]string{})
	assert.Nil(t, err)
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", "/a/dir/that/should/not/exist", nil, []string{},[]string{})
	err = engine.ProcessDirectory(dir+"/a/dir/that/should/not/exist", "/a/dir/that/should/not/exist", nil, nil,nil)
	assert.Error(t, err)
}

func TestDelims(t *testing.T) {
	var engine template_engine.Engine
	engine, _ = template_engine.GetEngine(TEST_DELIMS[0], TEST_DELIMS[1])
	vars, err := engine.LoadVars("test_fixtures/delim.yml")
	out, err := engine.ParseTemplateFile("test_fixtures/delim.tpl", vars)
	assert.Nil(t, err)
	assert.Contains(t, out, "Willie")
	assert.Contains(t, out, "horses")
	assert.Contains(t, out, "beer")
}

func TestGetEngine(t *testing.T) {
	gte, err := template_engine.GetEngine()
	assert.NotNil(t, gte)
	assert.Nil(t, err)
	gte, err = template_engine.GetEngine("{{{", "}}}")
	assert.NotNil(t, gte)
	assert.Nil(t, err)
}

func TestTemplateEngine_StaticInclude(t *testing.T) {
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, err := engine.ParseTemplateFile("test_fixtures/static-include.yml", params)
	assert.Nil(t,err)
	assert.NotNil(t,out)
	fmt.Println(out)
}

func TestTemplateEngine_replace(t *testing.T) {
	engine, _ := template_engine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["name"] = "Jolito"
	out, err := engine.ParseTemplateFile("test_fixtures/replace.yml", params)
	assert.Nil(t,err)
	assert.NotNil(t,out)
	fmt.Println(out)
}
