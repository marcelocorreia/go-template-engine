# go-template-engine GTE

Based on Golang templates text SDK.

## TLDR;

- Accepts JSON and YAML variable file
- Looks up file extension and parses according to the file extension, accepts **.json .yml .yaml** extensions
- These example are pretty vanilla, go templates are actually pretty powerful, check the links for more info.
    - [https://golang.org/pkg/text/template/](https://golang.org/pkg/text/template/)
    - [https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html)
### Usage
files in the examples are located at template-engine/test_fixtures/
```
$> go-template-engine --source template-engine/test_fixtures/cfn.tpl.yml \
        --var-file template-engine/test_fixtures/cfn-vars.yml \
        --output cfn.yml

$> go-template-engine --source template-engine/test_fixtures/cfn.tpl.yml \
        --var-file template-engine/test_fixtures/cfn-vars.json \
        --output cfn.yml
```

### Examples

#### Simple CFN Template
```yaml
AWSTemplateFormatVersion: 2010-09-09
Description: VPC's sample

Resources: {{range .network.vpcs}}{{$vpc_name := .name}}
  {{.name}}:
    Type: 'AWS::EC2::VPC'
    Properties:
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      CidrBlock: {{.cidr}}
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'
{{range .subnets}}
  {{.name}}:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId: !Ref {{$vpc_name}}
      CidrBlock: {{.cidr}}
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'
{{end}}{{end}}
Outputs: {{range .network.vpcs}}
  {{.name}}:
    Description: VPC ID of {{.name}}
    Value: !Ref {{.name}}
{{range .subnets}}
  PrivateSubnet:
    Description: Subnet ID of {{.name}}
    Value: !Ref {{.name}}
{{end}}
{{end}}
```

#### YAML Variables
```yaml
network:
  vpcs:
    - name: VPCA
      cidr: 10.11.0.0/16
      subnets:
        - name: SubnetA1
          cidr: 10.11.1.0/24
        - name: SubnetA2
          cidr: 10.11.2.0/24
    - name: VPCB
      cidr: 10.12.0.0/16
      subnets:
        - name: SubnetB1
          cidr: 10.12.1.0/24
        - name: SubnetB2
          cidr: 10.12.2.0/24
```
#### JSON Variables
```json
{
  "network": {
    "vpcs": [
      {
        "cidr": "10.11.0.0/16",
        "name": "VPCA",
        "subnets": [
          {
            "cidr": "10.11.1.0/24",
            "name": "SubnetA1"
          },
          {
            "cidr": "10.11.2.0/24",
            "name": "SubnetA2"
          }
        ]
      },
      {
        "cidr": "10.12.0.0/16",
        "name": "VPCB",
        "subnets": [
          {
            "cidr": "10.12.1.0/24",
            "name": "SubnetB1"
          },
          {
            "cidr": "10.12.2.0/24",
            "name": "SubnetB2"
          }
        ]
      }
    ]
  }
}
```
#### Output
```yml
AWSTemplateFormatVersion: 2010-09-09
Description: VPC's sample

Resources:
  VPCA:
    Type: 'AWS::EC2::VPC'
    Properties:
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      CidrBlock: 10.11.0.0/16
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

  SubnetA1:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId: !Ref VPCA
      CidrBlock: 10.11.1.0/24
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

  SubnetA2:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId: !Ref VPCA
      CidrBlock: 10.11.2.0/24
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

  VPCB:
    Type: 'AWS::EC2::VPC'
    Properties:
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      CidrBlock: 10.12.0.0/16
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

  SubnetB1:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId: !Ref VPCB
      CidrBlock: 10.12.1.0/24
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

  SubnetB2:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId: !Ref VPCB
      CidrBlock: 10.12.2.0/24
      Tags:
        - Key: Application
          Value: !Ref 'AWS::StackName'

Outputs:
  VPCA:
    Description: VPC ID of VPCA
    Value: !Ref VPCA

  PrivateSubnet:
    Description: Subnet ID of SubnetA1
    Value: !Ref SubnetA1

  PrivateSubnet:
    Description: Subnet ID of SubnetA2
    Value: !Ref SubnetA2


  VPCB:
    Description: VPC ID of VPCB
    Value: !Ref VPCB

  PrivateSubnet:
    Description: Subnet ID of SubnetB1
    Value: !Ref SubnetB1

  PrivateSubnet:
    Description: Subnet ID of SubnetB2
    Value: !Ref SubnetB2
```



#### whatever else...
```text
# {{.package_name}}

{{.phrase1}}
{{.phrase1}}

They're forming in a straight line
They're going through a tight wind
The kids are losing their minds
The Blitzkrieg Bop

They're piling in the back seat
They're generating steam heat
Pulsating to the back beat
The Blitzkrieg Bop.

{{.phrase1}}
Shoot'em in the back now
What they want, I don't know
They're all reved up and ready to go

{{.the.end}}
```

```
package_name: Blitzkrieg Bop
phrase1: Hey ho, let's go
the: {end: Tommy & Dee Dee Ramone}
```

```markdown
# Blitzkrieg Bop

Hey ho, let's go
Hey ho, let's go

They're forming in a straight line
They're going through a tight wind
The kids are losing their minds
The Blitzkrieg Bop

They're piling in the back seat
They're generating steam heat
Pulsating to the back beat
The Blitzkrieg Bop.

Hey ho, let's go
Shoot'em in the back now
What they want, I don't know
They're all reved up and ready to go

Tommy & Dee Dee Ramone
```

## Development

```bash
$> go get github.com/marcelocorreia/go-template-engine/template-engine
```

```golang
var engine template_engine.Engine
engine = template_engine.TemplateEngine{}
file, _ := ioutil.ReadFile("test_fixtures/bb.json")
var varsJson interface{}
json.Unmarshal(file, &varsJson)
outJson, _ := engine.ParseTemplateFile("test_fixtures/bb.txt.tpl", varsJson)
```


### TODO's
- [ ] Accept multiple variable files