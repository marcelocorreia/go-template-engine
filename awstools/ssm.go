package awstools

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"os"
)

type parameterStore struct {
	AwsProfile string
	Region     string
}

type ParameterStore interface {
	GetParameter(name string) (string, error)
	PutParameter(value, name string) error
	GetParameterField(name, field string) (string, error)
	GetParametersByPath(path string, recursive bool, encrypted bool) error
}

func NewParameterStore() ParameterStore {
	p := parameterStore{}
	if os.Getenv("AWS_DEFAULT_REGION") != "" {
		p.Region = os.Getenv("AWS_DEFAULT_REGION")
	} else {
		p.Region = "ap-southeast-2"
	}

	return p
}

func (pm parameterStore) PutParameter(name, value string) error {
	ssmsvc := ssm.New(GetSession(pm.Region))
	_, err := ssmsvc.PutParameter(&ssm.PutParameterInput{
		Value: aws.String(value),
		Name:  aws.String(name),
		Type:  aws.String("String"),
	})
	return err
}

func (pm parameterStore) GetParameter(name string) (string, error) {
	ssmsvc := ssm.New(GetSession(pm.Region))
	out, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(name),
	})
	return *out.Parameter.Value, err
}

func (pm parameterStore) GetParameterField(name, field string) (string, error) {
	ssmsvc := ssm.New(GetSession(pm.Region))
	out, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(name),
	})
	val := *out.Parameter.Value
	var f map[string]interface{}
	err = json.Unmarshal([]byte(val), &f)
	if err != nil {
		return "", err
	}

	return f[field].(string), err
}


func (pm parameterStore) GetParametersByPath(path string, recursive bool, encrypted bool) error {
	ssmsvc := ssm.New(GetSession(pm.Region))
	out, err := ssmsvc.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(recursive),
		WithDecryption: aws.Bool(encrypted),
	})

	for _, p := range out.Parameters {
		fmt.Println(p.Value)
	}

	return err
}


func GetSession(region string) *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	}))
	return sess
}


