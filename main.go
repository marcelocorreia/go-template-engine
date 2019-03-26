package main

import (
	"fmt"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app = kingpin.New("go-template-engine", "go-template-engine")

	templateFile        = app.Flag("source", "Template Source File").Short('s').String()
	headerTemplateFiles = app.Flag("header-sources", "Extra Source Files to append as HEADER to the main template before processing."+
		"Useful to preload embedded templates").Strings()
	footerTemplateFiles = app.Flag("footer-sources", "Extra Source Files to append as FOOTER to the main template before processing."+
		"Useful to preload embedded templates").Strings()

	templateVars        = app.Flag("var", "Params & Variables. Example --var hey=ho --var lets=go").StringMap()
	templateVarsFile    = app.Flag("var-file", "Variables File").String()
	templateIgnores     = app.Flag("skip-parsing", "Includes file but skips parsing").Strings()
	templateExcludes    = app.Flag("exclude", "Excludes File from template job").Strings()
	templateExcludesDir = app.Flag("exclude-dir", "Excludes directory from template job").Strings()
	templateFileOutput  = app.Flag("output", "File output full path").Short('o').String()

	delimLeft           = app.Flag("delim-left", "Left Delimiter").Default("{{").String()
	delimRight          = app.Flag("delim-right", "Right Delimiter").Default("}}").String()
	listCustomFunctions = app.Flag("list-custom-functions", "List Custom Commands").Short('c').Bool()
	VERSION             string
)

func main() {
	app.Version(VERSION).VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('h')
	if len(os.Args) <= 1 {
		kingpin.Usage()
		os.Exit(1)
	}

	kingpin.MustParse(app.Parse(os.Args[1:]))

	var engine template_engine.Engine
	engine, err := template_engine.GetEngine(*delimLeft, *delimRight)
	if err != nil {
		handleErrorExit(err, "Error Loading engine")
	}

	if *listCustomFunctions {
		engine.ListFuncs()
		os.Exit(0)
	}

	var jobVars interface{}

	jobVars, err = engine.LoadVars(*templateVarsFile)
	if err != nil {
		handleErrorExit(err, "Error:")
	}

	for k, v := range *templateVars {
		jobVars.(map[interface{}]interface{})[k] = v
	}
	render(jobVars, engine)

	os.Exit(0)

}

func render(jobVars interface{}, engine template_engine.Engine) {
	if info, err := os.Stat(*templateFile); err == nil && info.IsDir() {
		err := engine.ProcessDirectory(*templateFile, *templateFileOutput, jobVars, *templateExcludesDir, *templateExcludes, *templateIgnores)
		if err != nil {
			handleErrorExit(err, fmt.Sprintf("Error Processing templates @ dir: %s\n", *templateFile))
		}
	} else {
		var out string
		for _, headerFile := range *headerTemplateFiles {
			out+= parse(headerFile, jobVars,engine)
		}

		out+= parse(*templateFile, jobVars,engine)
		for _, footerFile := range *footerTemplateFiles {
			out+= parse(footerFile, jobVars,engine)
		}

		if err != nil {
			handleErrorExit(err, "Error running template.\n")
		}
		output(out)
	}
}

func parse(template string, jobVars interface{},engine template_engine.Engine)(string){
	out,err:=engine.ParseTemplateFile(template,jobVars)
	if err != nil {
		handleErrorExit(err, "Error running template.\n")
	}
	return out
}

func handleErrorExit(err error, msg string) {
	fmt.Println(msg, err)
	os.Exit(1)
}

func output(out string) {
	if *templateFileOutput != "" {
		err := ioutil.WriteFile(*templateFileOutput, []byte(out), 0755)
		if err != nil {
			handleErrorExit(err, "Error writing file to "+*templateFileOutput)
		}
	} else {
		fmt.Println(out)
	}
}
