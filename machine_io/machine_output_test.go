package machine_io

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

const jsonInput  =`
{
	"name":"Jose Mane",
	"age": 50,
	"children":[
		{
			"name":"Jose Manezin",
			"age": 30
		},
		{
			"name":"Jose Roela",
			"age": 24
		}
	]
}`

const yamlInput  =`
name: Jose Mane
age: 50
children:
- age: 30
  name: Jose Manezin
- age: 24
  name: Jose Roela
`
type TestType struct {
	Name     string `json:"name" yaml:"name"`
	Age      int `json:"age" yaml:"age"`
	Children []TestType `json:"children" yaml:"children"`
}

func TestOutput(t *testing.T) {
	dad := Sample()
	outJ, errJ := JsonOutput(dad)
	assert.Nil(t, errJ)
	fmt.Println(outJ)
	outY, errY := YamlOutput(dad)
	assert.Nil(t, errY)
	fmt.Println(outY)
	c:=Converter{}
	j2y,_:=c.Json2Yaml(jsonInput)
	fmt.Println(j2y)
	y2j,_:=c.Yaml2Json(yamlInput)
	fmt.Println(y2j)
}

func Sample() (TestType) {
	son := TestType{
		Name: "Jim", Age:50,
	}
	
	daughter := TestType{
		Name: "Mary", Age:10,
	}
	
	dad := TestType{
		Name: "John", Age:98,
		Children:[]TestType{son, daughter},
	}
	return dad
}


