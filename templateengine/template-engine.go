package templateengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/hashicorp/hcl"
	"github.com/marcelocorreia/go-utils/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//DELIMS Go Template delims
var DELIMS = []string{"{{", "}}"}

//Engine Engine
type Engine interface {
	ParseTemplateFile(templateFile string, params interface{}) (string, error)
	ParseTemplateString(templateString string, params interface{}) (string, error)
	LoadVars(filePath string) (interface{}, error)
	ProcessDirectory(sourceDir string, targetDir string, params interface{}, dirExclusions []string, fileExclusions []string, fileIgnores []string) error
	GetFileList(dir string, dirExclusions []string, fileExclusions []string) ([]string, error)
	PrepareOutputDirectory(sourceDir string, targetDir string, exclusions []string) error
	loadFuncs()
	ListFuncs()
	staticInclude(sourceFile string) string
	replace(input, from, to string) string
	inList(needle interface{}, haystack []interface{}) bool
	printf(pattern string, params ...string) string
}

// TemplateEngine templateengine
type TemplateEngine struct {
	Delims []string
	Funcs  map[string]interface{}
}

//GetEngine returns engine
func GetEngine(delims ...string) (*TemplateEngine, error) {
	if len(delims) == 2 {
		DELIMS = delims
	}

	engine := TemplateEngine{
		Delims: DELIMS,
		Funcs:  make(map[string]interface{}),
	}
	engine.loadFuncs()
	return &engine, nil
}

//ParseTemplateFile Parses file
func (gte TemplateEngine) ParseTemplateFile(templateFile string, params interface{}) (string, error) {
	tplFile, err := ioutil.ReadFile(templateFile)

	if err != nil {
		return "", err
	}

	r, err := gte.ParseTemplateString(string(tplFile), params)
	if err != nil {
		return r, err
	}
	return r, nil
}

//ParseTemplateString Parses string
func (gte TemplateEngine) ParseTemplateString(templateString string, params interface{}) (string, error) {
	funcMap := template.FuncMap{
		"staticInclude": func(path string) string { return gte.staticInclude(path) },
		"replace":       func(input, from, to string) string { return gte.replace(input, from, to) },
		"inList":        func(needle interface{}, haystack []interface{}) bool { return gte.inList(needle, haystack) },
	}

	t, err := template.New("gte").Delims(gte.Delims[0], gte.Delims[1]).Funcs(funcMap).Funcs(sprig.GenericFuncMap()).Parse(templateString)
	if err != nil {
		return templateString, err
	}
	var doc bytes.Buffer
	errParse := t.Execute(&doc, params)

	if errParse != nil {
		return "", errParse
	}
	resp := doc.String()

	return resp, nil
}

//LoadVars Load all vars for job
func (gte TemplateEngine) LoadVars(filePath string) (interface{}, error) {
	var varsFile interface{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(filePath, ".json") {
		err := json.Unmarshal(file, &varsFile)
		if err != nil {
			return nil, err
		}
	} else if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		err := yaml.Unmarshal(file, &varsFile)
		if err != nil {
			return nil, err
		}
	} else if strings.HasSuffix(filePath, ".tf") || strings.HasSuffix(filePath, ".tfvars") {
		err := hcl.Unmarshal(file, &varsFile)
		if err != nil {
			return nil, err
		}
	} else {
		varsFile = make(map[interface{}]interface{})
	}
	return varsFile, nil
}

//ProcessDirectory Parses a full directory
func (gte TemplateEngine) ProcessDirectory(sourceDir string, targetDir string, params interface{}, dirExclusions []string, fileExclusions []string, fileIgnores []string) error {
	err := gte.PrepareOutputDirectory(sourceDir, targetDir, dirExclusions)
	if err != nil {
		return err
	}
	list, err := gte.GetFileList(sourceDir, dirExclusions, fileExclusions)

	if err != nil {
		return err
	}
	for _, f := range list {
		sourceFile := f
		targetFile := strings.Replace(f, sourceDir, targetDir, -1)
		isDir, err := IsDirectory(sourceFile)
		if err != nil {
			return err
		}
		if !isDir {
			_, err := ioutil.ReadFile(sourceFile)
			if err != nil {
				return err
			}
			file, err := os.Stat(sourceFile)
			if err != nil {
				return err
			}
			if !utils.StringInSlice(file.Name(), fileExclusions) {
				b, err := gte.ParseTemplateFile(sourceFile, params)
				if err != nil {
					fmt.Printf("File: %s can't be loaded as template,\n\tContent written without modifications.\n\tPlease check the flags is case this is not correct.\n-----------------------------\n", sourceFile)
				}
				if err := Output(b, targetFile); err != nil {
					return err
				}
			}

		}
	}

	return nil
}

//GetFileList file list
func (gte TemplateEngine) GetFileList(dir string, dirExclusions []string, fileExclusions []string) ([]string, error) {
	var files []string
	exists, err := Exists(dir)
	if err != nil || !exists {
		return nil, errors.New(dir + "does not exist")
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

//PrepareOutputDirectory Prepares output directory
func (gte TemplateEngine) PrepareOutputDirectory(sourceDir string, targetDir string, exclusions []string) error {
	if targetDir == "" {
		return errors.New("Output must be provided when source is a directory")
	}

	CreateNewDirectoryIfNil(targetDir)
	files, err := gte.GetFileList(sourceDir, exclusions, exclusions)
	if err != nil {
		return err
	}
	for _, d := range files {

		if info, err := os.Stat(d); err == nil && info.IsDir() {
			newDir := strings.Replace(d, sourceDir, targetDir, -1)
			CreateNewDirectoryIfNil(newDir)
		}

	}

	return nil
}

//Output Output
func Output(out string, templateFileOutput string) error {
	return ioutil.WriteFile(templateFileOutput, []byte(out), 0755)
}
