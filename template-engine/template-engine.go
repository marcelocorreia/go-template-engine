package template_engine

type Engine interface {
	ParseTemplateFile(templateFile string, params interface{}) (string, error)
	ParseTemplateString(templateString string, params interface{}) (string, error)
	VariablesFileMerge(varsFile []string, extra_vars map[string]string) (string, error)
	LoadVars(filePath string) (interface{}, error)
	ProcessDirectory(sourceDir string, targetDir string, params interface{}, exclusions []string) (error)
	GetFileList(dir string, fullPath bool, exclusions []string) ([]string, error)
	PrepareOutputDirectory(sourceDir string, targetDir string, exclusions []string) (error)

}

type TemplateEngine struct{}
