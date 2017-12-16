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
	"errors"
)

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

func (gte TemplateEngine) ParseTemplateString(templateString string, params interface{}) (string, error) {
	t, err := template.New("letter").Parse(templateString)
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

func (gte TemplateEngine) LoadVars(filePath string) (interface{}, error) {
	var varsFile interface{}
	file, _ := ioutil.ReadFile(filePath)

	if strings.HasSuffix(filePath, ".json") {
		err := json.Unmarshal(file, &varsFile)
		fmt.Println(&varsFile)
		if err != nil {
			return nil, err
		}
	} else if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		err := yaml.Unmarshal(file, &varsFile)
		if err != nil {
			return nil, err
		}
	}
	return varsFile, nil
}

func (gte TemplateEngine) VariablesFileMerge(varsFile []string, extra_vars map[string]string) (string, error) {
	tmpFile, _ := ioutil.TempFile("/tmp", "vars")

	for _, file := range varsFile {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		tmpFile.Write(content)
		tmpFile.WriteString("\n")
	}

	for k, v := range extra_vars {
		tmpFile.WriteString(fmt.Sprintf("%s: %s\n", k, v))
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
	err := gte.PrepareOutputDirectory(sourceDir, targetDir, exclusions)
	if err != nil {
		return err
	}
	list, err := gte.GetFileList(sourceDir, false, exclusions)

	if err != nil {
		return err
	}
	for _, f := range list {
		sourceFile := fmt.Sprintf("%s/%s", sourceDir, f)
		targetFile := fmt.Sprintf("%s/%s", targetDir, f)
		body, err := gte.ParseTemplateFile(sourceFile, params)
		if err != nil {
			fmt.Printf("File: %s can't be loaded as template,\n\tContent writen without modifications.\n\tPlease check the tags is case this is not correct.\n-----------------------------\n%s\n-----------------------------\n", sourceFile, body)
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
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
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
		return errors.New("output must be provided when source is a directory")
	}

	utils.CreateNewDirectoryIfNil(targetDir)
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	for _, d := range files {
		if !utils.StringInSlice(d.Name(), exclusions) {
			if info, err := os.Stat(sourceDir + "/" + d.Name()); err == nil && info.IsDir() {
				utils.CreateNewDirectoryIfNil(targetDir + "/" + d.Name())
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


