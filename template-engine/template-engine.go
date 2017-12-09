package template_engine

type Engine interface {
	ParseTemplateFile(templateFile string, params interface{}) (string, error)
	ParseTemplateString(templateString string, params interface{}) (string, error)
	VariablesFileMerge(varsFile []string) (string, error)
	//LoadVars(varsFile []string) (string, error)
}

type TemplateEngine struct {}

