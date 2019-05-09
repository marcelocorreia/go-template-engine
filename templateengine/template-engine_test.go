package templateengine_test

import (
	"encoding/json"
	"fmt"
	"github.com/marcelocorreia/go-template-engine/templateengine"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var TEST_DELIMS = []string{"{{{", "}}}"}
var DEFAULT_DELIMS = []string{"{{", "}}"}

var path string

func init() {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = p
	fmt.Print(path)
}

func TestParseTemplateString(t *testing.T) {
	fmt.Println("Running Test with vars...")
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := engine.ParseTemplateFile(path+"/test_fixtures/bb.txt.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")
	fmt.Println("Finished Test with vars...")
}
func TestListFuncs(t *testing.T) {
	engine, _ := templateengine.GetEngine()
	engine.ListFuncs()
}
func TestTemplateJson(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	fmt.Println("Running Test with JSON file...")
	fmt.Println("===================================================")
	file, _ := ioutil.ReadFile(path + "/test_fixtures/bb.json")
	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	outJson, _ := engine.ParseTemplateFile(path+"/test_fixtures/bb.txt.tpl", varsJson)
	assert.Contains(t, outJson, "Blitzkrieg Bop")
	assert.Contains(t, outJson, "The kids are losing their minds")
	assert.Contains(t, outJson, "Hey ho, let's go")
	fmt.Println(outJson)
	fmt.Println("===================================================")
	fmt.Println("Finished Test with JSON file...")
}

func TestTemplateErrorJson(t *testing.T) {
	fmt.Println("Running Testing throwing error...")
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	file, _ := ioutil.ReadFile(path + "/test_fixtures/vars.json-should-not-exist")
	var varsJson interface{}
	json.Unmarshal(file, &varsJson)
	_, err := engine.ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)
	fmt.Println("Finished Testing throwing error...")
}

func TestTemplateEngine_GetFileList(t *testing.T) {
	dir := path
	//dir := "/go/src/github.com/marcelocorreia/badwolf-templates/templates/badwolf/terraform-stack"

	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	ll, _ := engine.GetFileList(dir, []string{}, []string{})
	assert.True(t, len(ll) > 0)
	//_, err := engine.GetFileList("/a/dir/that/should/not/exist", true, []string{}, []string{})
	//assert.Error(t, err)
}

func TestPrepareOutputDirectory(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	dir := path + "/test_fixtures/base"
	tmpDir, err := ioutil.TempDir("/tmp", "gteTest-")
	if err != nil {
		panic(err)
	}
	engine.PrepareOutputDirectory(dir, tmpDir, []string{".templates", "ci"})
	exists, err := templateengine.Exists(tmpDir)
	if err != nil {
		panic(err)
	}
	assert.True(t, exists)
	os.RemoveAll(tmpDir)
	exists, _ = templateengine.Exists(tmpDir)
	assert.False(t, exists)
	_, err = ioutil.TempDir("/bogus", "gteTest-")

	assert.Error(t, err)

	err = engine.PrepareOutputDirectory(dir, "", []string{})
	assert.Error(t, err)
}

func getParams() map[string]string {
	params := make(map[string]string)
	params["hey"] = "Ho"
	params["Lets"] = "go"
	params["name"] = "Willie Nelson"
	return params
}

func TestTemplateEngine_LoadVars(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
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
	vars, err = engine.LoadVars(dir + "/test_fixtures/variables.tfvars")
	assert.NotNil(t, vars)
	assert.Nil(t, err)
}

func TestTemplateEngine_ProcessDirectory(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	dir, _ := os.Getwd()
	tmpDir := os.TempDir()
	err := engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, nil, nil, nil)
	assert.Nil(t, err)
	os.RemoveAll(tmpDir)

	tmpDir = os.TempDir()
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, nil, nil, nil)
	assert.Nil(t, err)
	os.RemoveAll(tmpDir)

	tmpDir = os.TempDir()
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", tmpDir, nil, nil, nil, nil)
	assert.Nil(t, err)
	exists, err := templateengine.Exists(tmpDir + "/.variables.tfvars")
	assert.True(t, exists)
	assert.NoError(t, err)
	os.RemoveAll(tmpDir)

	tmpDir = os.TempDir()
	err = engine.ProcessDirectory(dir+"/test_fixtures/base", "/a/dir/that/should/not/exist", nil, nil, nil, nil)
	assert.Error(t, err)
	os.RemoveAll(tmpDir)

	tmpDir = os.TempDir()
	err = engine.ProcessDirectory(dir+"/a/dir/that/should/not/exist", "/a/dir/that/should/not/exist", nil, nil, nil, nil)
	assert.Error(t, err)
	os.RemoveAll(tmpDir)
}

func TestDelims(t *testing.T) {
	var engine templateengine.Engine
	engine, _ = templateengine.GetEngine(TEST_DELIMS[0], TEST_DELIMS[1])
	vars, err := engine.LoadVars("test_fixtures/delim.yml")
	assert.NoError(t, err)
	out, err := engine.ParseTemplateFile("test_fixtures/delim.tpl", vars)
	assert.NoError(t, err)
	assert.Contains(t, out, "Willie")
	assert.Contains(t, out, "horses")
	assert.Contains(t, out, "beer")
}

func TestGetEngine(t *testing.T) {
	gte, err := templateengine.GetEngine()
	assert.NotNil(t, gte)
	assert.Nil(t, err)
	gte, err = templateengine.GetEngine("{{{", "}}}")
	assert.NotNil(t, gte)
	assert.Nil(t, err)
}

func TestTemplateEngine_StaticInclude(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, err := engine.ParseTemplateFile("test_fixtures/static-include.yml", params)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	fmt.Println(out)
}

func TestTemplateEngine_replace(t *testing.T) {
	engine, _ := templateengine.GetEngine(DEFAULT_DELIMS[0], DEFAULT_DELIMS[1])
	params := make(map[string]string)
	params["name"] = "Jolito"
	out, err := engine.ParseTemplateFile("test_fixtures/replace.yml", params)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	fmt.Println(out)
}
