package main

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app                  = kingpin.New("go-template-engine", "")
	templateFile         = app.Flag("source", "Template Source File").Short('s').Required().String()
	templateVars         = app.Flag("var", "Params & Variables. Example --var hey=ho --var lets=go").StringMap()
	templateVarsFile     = app.Flag("var-file", "Variables File").Strings()
	templatesExcludesDir = app.Flag("exclude-dir", "Variables File").Strings()
	templateFileOutput   = app.Flag("output", "File output full path").Short('o').String()
	versionFlag          = app.Flag("version", "App Version").Short('v').Bool()
	VERSION              string
)

func main() {
	engine := *getEngine()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *versionFlag {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	var varsBundle interface{}

	varsBundle, _ = engine.VariablesFileMerge(*templateVarsFile)
	jobVars, err := engine.LoadVars(varsBundle.(string))
	if err != nil {
		handleErrorExit("Eita:", err)
	}
	if info, err := os.Stat(*templateFile); err == nil && info.IsDir() {
		engine.ProcessDirectory(*templateFile, *templateFileOutput, jobVars, *templatesExcludesDir)
	} else {
		out, err := engine.ParseTemplateFile(*templateFile, jobVars)
		if err != nil {
			handleErrorExit("Error running template.\n", err)
		}
		output(out)
	}

	ct.ResetColor()
	os.Exit(0)

}

func handleErrorExit(msg string, err error) {
	ct.Foreground(ct.Red, false)
	fmt.Println(msg, err)
	ct.ResetColor()
	os.Exit(1)
}

func output(out string) {
	ct.ResetColor()
	if *templateFileOutput != "" {
		err := ioutil.WriteFile(*templateFileOutput, []byte(out), 0755)
		if err != nil {
			handleErrorExit("Error writing file to " + *templateFileOutput, err)
		}
	} else {
		ct.Foreground(ct.Green, false)
		fmt.Println(out)
		ct.ResetColor()
	}
}

func getEngine() *template_engine.Engine {
	var engine template_engine.Engine
	engine = template_engine.TemplateEngine{}

	return &engine
}
