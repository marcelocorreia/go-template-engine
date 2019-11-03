package awstools

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
	"sort"
)

const maxResults = 100

type SecretsManagerService interface {
	GetSecretString(key string) (string, error)
	GetSecretField(key, field string) (string, error)
	Secrets() ([]string, error)
	PrintSecretsList() error
}

type SecretsManager struct {
	AwsProfile string
	Region     string
}

func NewSecretsManagerService() SecretsManagerService {
	sm := SecretsManager{}
	if os.Getenv("AWS_DEFAULT_REGION") != "" {
		sm.Region = os.Getenv("AWS_DEFAULT_REGION")
	} else {
		sm.Region = "ap-southeast-2"
	}

	return sm
}

func (sm SecretsManager) PrintSecretsList() error {
	list, err := sm.Secrets()
	if err != nil {
		return err
	}
	for _, sec := range list {
		fmt.Println(sec)
	}
	return nil
}

func (sm SecretsManager) Secrets() ([]string, error) {
	cli := secretsmanager.New(sm.getSession(sm.Region))
	var res int64
	res = maxResults
	var list, err = cli.ListSecrets(&secretsmanager.ListSecretsInput{
		MaxResults: &res,
	})

	if err != nil {
		return nil, err
	}
	var resp []string

	for _, sec := range list.SecretList {
		resp = append(resp, *sec.Name)
	}
	sort.Strings(resp)
	return resp, nil
}

func (sm SecretsManager) GetSecretString(key string) (string, error) {
	cli := secretsmanager.New(sm.getSession(sm.Region))
	input := secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	}

	o, err := cli.GetSecretValue(&input)
	if err != nil {
		return "", err
	}
	return *o.SecretString, nil
}

func (sm SecretsManager) GetSecretField(key, field string) (string, error) {
	cli := secretsmanager.New(sm.getSession(sm.Region))
	input := secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	}

	o, err := cli.GetSecretValue(&input)
	if err != nil {
		return "", err
	}

	raw := *o.SecretString
	var resp map[string]string
	err = json.Unmarshal([]byte(raw), &resp)

	if _, ok := resp[field]; ok {
		return resp[field], err
	} else {
		return "", err
	}
}

func (sm SecretsManager) getSession(region string) *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	}))
	return sess
}

func (sm SecretsManager) whoami() (*sts.GetCallerIdentityOutput, error) {
	input := &sts.GetCallerIdentityInput{}
	svc := sts.New(sm.getSession(sm.Region))
	result, err := svc.GetCallerIdentity(input)

	if err != nil {
		return nil, err
	}
	return result, nil
}
