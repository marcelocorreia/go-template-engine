package main

import (
	"fmt"
	"github.com/marcelocorreia/go-template-engine/templateengine"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app                 = kingpin.New("go-template-engine", "go-template-engine")
	exitOnError         = app.Flag("exit-on-error", "Exits on error").Short('e').Bool()
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
	templateOptions     = app.Flag("option", "Go template options").Strings()

	delimLeft           = app.Flag("delim-left", "Left Delimiter").Default("{{").String()
	delimRight          = app.Flag("delim-right", "Right Delimiter").Default("}}").String()
	listCustomFunctions = app.Flag("list-custom-functions", "List Custom Commands").Short('c').Bool()
	logLevel            = app.Flag("log", "Log Level. Default Info").Default("info").String()
	noColor             = app.Flag("no-color", "No color output").Bool()
	//VERSION application
	VERSION string
)

func main() {
	app.Version(VERSION).VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('h')
	if len(os.Args) <= 1 {
		kingpin.Usage()
		os.Exit(1)
	}

	kingpin.MustParse(app.Parse(os.Args[1:]))

	switch *logLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)

	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	var engine templateengine.Engine
	engine, err := templateengine.GetEngine(*exitOnError, []string{*delimLeft, *delimRight}, *templateOptions...)
	if err != nil {
		handleErrorExit(err, "Error Loading engine")
	}

	if *listCustomFunctions {
		engine.ListFuncs()
		os.Exit(0)
	}

	var jobVars map[string]interface{}
	if templateVarsFile != nil && *templateVarsFile != "" {
		jobVars, err = engine.LoadVars(*templateVarsFile)

		if err != nil {
			handleErrorExit(err, "Error:")
		}

	}

	if jobVars == nil {
		jobVars = make(map[string]interface{})
	}

	for k, v := range *templateVars {
		jobVars[k] = v
	}
	render(jobVars, engine)

	os.Exit(0)

}

func render(jobVars interface{}, engine templateengine.Engine) {
	if info, err := os.Stat(*templateFile); err == nil && info.IsDir() {
		err := engine.ProcessDirectory(*templateFile, *templateFileOutput, jobVars, *templateExcludesDir, *templateExcludes, *templateIgnores)
		if err != nil {
			handleErrorExit(err, fmt.Sprintf("Error Processing templates @ dir: %s\n", *templateFile))
		}
	} else {
		var out string
		for _, headerFile := range *headerTemplateFiles {
			out += parse(headerFile, jobVars, engine)
		}

		out += parse(*templateFile, jobVars, engine)
		for _, footerFile := range *footerTemplateFiles {
			out += parse(footerFile, jobVars, engine)
		}

		if err != nil {
			handleErrorExit(err, "Error running template.\n")
		}
		output(out)
	}
}

func parse(template string, jobVars interface{}, engine templateengine.Engine) string {
	out, err := engine.ParseTemplateFile(template, jobVars)
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
