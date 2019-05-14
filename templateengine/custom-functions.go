package templateengine

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
)

func (gte TemplateEngine) staticInclude(sourceFile string) string {
	body, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return fmt.Sprintf("ERROR including file: %s\n", sourceFile)
	}
	return string(body)
}

func (gte TemplateEngine) replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func (gte TemplateEngine) inList(needle interface{}, haystack []interface{}) bool {
	for _, h := range haystack {
		if reflect.DeepEqual(needle, h) {
			return true
		}
	}
	return false
}

func (gte TemplateEngine) printf(pattern string, params ...string) string {
	return fmt.Sprintf(pattern, params)
}

//ListFuncs Lists Custom functions
func (gte TemplateEngine) ListFuncs() {
	funcs := make([]string, 0, len(gte.Funcs))
	for k := range gte.Funcs {
		funcs = append(funcs, k)
	}
	sort.Strings(funcs)

	for _, v := range funcs {
		fmt.Println(v)
	}
}

func (gte *TemplateEngine) loadFuncs() {
	for k, v := range sprig.GenericFuncMap() {
		gte.Funcs[k] = v
	}

	gte.Funcs["staticInclude"] = func(path string) string { return gte.staticInclude(path) }
	gte.Funcs["replace"] = func(input, from, to string) string { return gte.replace(input, from, to) }
	gte.Funcs["inList"] = func(needle interface{}, haystack []interface{}) bool { return gte.inList(needle, haystack) }
	gte.Funcs["printf"] = func(pattern string, params ...string) string { return gte.printf(pattern, params...) }
}