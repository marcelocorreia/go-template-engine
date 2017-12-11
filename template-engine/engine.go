package template_engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"bufio"
	"github.com/marcelocorreia/go-template-engine/utils"
)

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

func (gte TemplateEngine) LoadVars(filePath string) (interface{}, error) {
	var varsFile interface{}
	file, _ := ioutil.ReadFile(filePath)

	if strings.HasSuffix(filePath, ".json") {
		err := json.Unmarshal(file, &varsFile)
		fmt.Println(&varsFile)
		if err != nil {
			return "", err
		}
	} else if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		err := yaml.Unmarshal(file, &varsFile)
		if err != nil {
			return "", err
		}
	}
	return varsFile, nil
}

func (gte TemplateEngine) VariablesFileMerge(varsFile []string, extra_vars map[string]string) (string, error) {
	tmpFile, err := ioutil.TempFile("/tmp", "vars")
	if err != nil {
		return "", err
	}
	for _, file := range varsFile {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		tmpFile.Write(content)
		tmpFile.WriteString("\n")
	}

	for k, v := range extra_vars {
		tmpFile.WriteString(fmt.Sprintf("%s: %s\n",k,v))
	}
	tmpFile.Close()
	cleanFile, err := cleanYamlFile(tmpFile.Name())
	if err != nil {
		return "", err
	}
	os.Remove(tmpFile.Name())
	utils.CopyFile(cleanFile, cleanFile+".yml")

	return cleanFile + ".yml", nil
}

func cleanYamlFile(file string) (string, error) {
	tmpFile, err := ioutil.TempFile("/tmp", "vars")
	if err != nil {
		return "", err
	}
	inFile, _ := os.Open(file)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "---") && len(line) > 0 {
			tmpFile.WriteString(line + "\n")
		}
	}
	return tmpFile.Name(), nil
}

func (gte TemplateEngine) ProcessDirectory(sourceDir string, targetDir string, params interface{}, exclusions []string) (error) {
	gte.PrepareOutputDirectory(sourceDir, targetDir, exclusions)

	list, err := gte.GetFileList(sourceDir, false, exclusions)

	if err != nil {
		fmt.Println("Error processing:", sourceDir)
		panic(err)
	}
	for _, f := range list {
		sourceFile := fmt.Sprintf("%s/%s", sourceDir, f)
		targetFile := fmt.Sprintf("%s/%s", targetDir, f)
		body, err := gte.ParseTemplateFile(sourceFile, params)
		if err != nil {
			return err
		}
		err = output(body, targetFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gte TemplateEngine) GetFileList(dir string, fullPath bool, exclusions []string) ([]string, error) {
	var fileList *[]string
	fileList = &[]string{}
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if !utils.StringInSlice(f.Name(), exclusions) {
			if info, err := os.Stat(dir + "/" + f.Name()); err == nil && info.IsDir() {
				gte.getTempList(dir+"/"+f.Name(), fileList)
			} else {
				*fileList = append(*fileList, dir+"/"+f.Name())
			}
		}
	}

	if fullPath {
		return *fileList, nil
	}

	list := []string{}
	for _, f := range *fileList {
		root := filepath.Base(dir)
		list = append(list, strings.Replace(strings.Replace(f, dir, root, -1), root+"/", "", -1))
	}
	return list, nil
}

func (gte TemplateEngine) getTempList(dir string, fileList *[]string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if info, err := os.Stat(dir + "/" + f.Name()); err == nil && info.IsDir() {
			gte.getTempList(dir+"/"+f.Name(), fileList)
		} else {
			*fileList = append(*fileList, dir+"/"+f.Name())
		}
	}
}

func (gte TemplateEngine) PrepareOutputDirectory(sourceDir string, targetDir string, exclusions []string) (error) {
	if targetDir == "" {
		fmt.Println("Error: output must be provided when source is a directory")
		os.Exit(1)
	}

	utils.CreateNewDirectory(targetDir)
	files, _ := ioutil.ReadDir(sourceDir)
	for _, d := range files {
		if !utils.StringInSlice(d.Name(), exclusions) {
			if info, err := os.Stat(sourceDir + "/" + d.Name()); err == nil && info.IsDir() {
				utils.CreateNewDirectory(targetDir + "/" + d.Name())
			}
		}
	}

	return nil
}

func output(out string, templateFileOutput string) (error) {
	err := ioutil.WriteFile(templateFileOutput, []byte(out), 0755)
	if err != nil {
		return err
	}
	return nil
}

///go/src/github.com/marcelocorreia/go-template-engine/bin/go-template-engine --source ci/go-template-engine.rb             --var dist_file=dist/go-template-engine-darwin-amd64-1.39.0.zip             --var version=1.39.0             --var hash_sum=123             > /Users/marcelo/IdeaProjects/tardis/homebrew-taps/go-template-engine.rb
