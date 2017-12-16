package main

import (
	"fmt"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app                  = kingpin.New("go-template-engine", "")
	templateFile         = app.Flag("source", "Template Source File").Short('s').String()
	templateVars         = app.Flag("var", "Params & Variables. Example --var hey=ho --var lets=go").StringMap()
	templateVarsFile     = app.Flag("var-file", "Variables File").Strings()
	templatesExcludesDir = app.Flag("exclude-dir", "Variables File").Strings()
	templateFileOutput   = app.Flag("output", "File output full path").Short('o').String()
	versionFlag          = app.Flag("version", "App Version").Short('v').Bool()
	VERSION              string
)

func main() {
	engine := *getEngine()
	if len(os.Args) <=1{
		kingpin.Usage()
		os.Exit(1)
	}
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *versionFlag {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	var varsBundle interface{}
	var jobVars interface{}

	varsBundle, _ = engine.VariablesFileMerge(*templateVarsFile, *templateVars)
	jobVars, err := engine.LoadVars(varsBundle.(string))
	if err != nil {
		handleErrorExit(err,"Error:")
	}

	render(jobVars, engine)

	os.Exit(0)

}

func render(jobVars interface{}, engine template_engine.Engine) {
	if info, err := os.Stat(*templateFile); err == nil && info.IsDir() {
		err := engine.ProcessDirectory(*templateFile, *templateFileOutput, jobVars, *templatesExcludesDir)
		if err != nil {
			handleErrorExit(err, fmt.Sprintf("Error Processing templates @ dir: %s\n", *templateFile))
		}
	} else {
		out, err := engine.ParseTemplateFile(*templateFile, jobVars)
		if err != nil {
			handleErrorExit(err, "Error running template.\n")
		}
		output(out)
	}
}
func handleErrorExit(err error, msg string) {
	fmt.Println(msg, err)
	os.Exit(1)
}

func output(out string) {
	if *templateFileOutput != "" {
		err := ioutil.WriteFile(*templateFileOutput, []byte(out), 0755)
		if err != nil {
			handleErrorExit(err,"Error writing file to " + *templateFileOutput)
		}
	} else {
		fmt.Println(out)
	}
}

func getEngine() *template_engine.Engine {
	var engine template_engine.Engine
	engine = template_engine.TemplateEngine{}

	return &engine
}
