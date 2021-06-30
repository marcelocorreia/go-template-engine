package awstools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parameterStore_GetParameter(t *testing.T) {
	ps := NewParameterStore()
	//
	//o,err:=ps.PutParameter("hey","ho")
	//assert.NoError(t, err)
	//fmt.Println(o)

	oo, err := ps.GetParameter("ho")
	assert.NoError(t, err)
	fmt.Println(oo)
}

func Test_parameterStore_GetParameters(t *testing.T) {
	ps := NewParameterStore()

	ps.GetParametersByPath("/dev/app01/*", true, false)
}
