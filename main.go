package main

import (
	"encoding/json"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/marcelocorreia/go-template-engine/template-engine"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	app                    = kingpin.New("go-template-engine", "")
	templateFile           = app.Flag("source", "Template Source File").Short('s').Required().String()
	templateVars           = app.Flag("var", "Params & Variables. Example --var hey=ho --var lets=go").StringMap()
	templateVarsFile       = app.Flag("var-file", "Variables File").String()
	templateVarsFileOutput = app.Flag("output", "File output full path").Short('o').String()
)

func main() {
	theGracefulDeath()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	var engine template_engine.Engine
	engine = template_engine.TemplateEngine{}
	engine.ParseTemplateFile(*templateFile, *templateVars)
	if *templateVarsFile != "" {
		var varsFile interface{}

		file, _ := ioutil.ReadFile(*templateVarsFile)

		if strings.HasSuffix(*templateVarsFile, ".json") {
			err := json.Unmarshal(file, &varsFile)
			fmt.Println(&varsFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}
		} else if strings.HasSuffix(*templateVarsFile, ".yaml") || strings.HasSuffix(*templateVarsFile, ".yml") {
			err := yaml.Unmarshal(file, &varsFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		out, err := engine.ParseTemplateFile(*templateFile, varsFile)
		if err != nil {
			ct.Foreground(ct.Red, false)
			fmt.Println("Error: running template.\n", err)
			ct.ResetColor()
			fmt.Println(err)
			os.Exit(1)
		}
		if *templateVarsFileOutput != "" {
			err = ioutil.WriteFile(*templateVarsFileOutput, []byte(out), 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			ct.Foreground(ct.Black, false)
			fmt.Println(out)
		}
		ct.ResetColor()
	}
}

func theGracefulDeath() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Shutting down gracefully...")
		ct.ResetColor()
		defer fmt.Println("Done.")
		os.Exit(0)
	}()
}
