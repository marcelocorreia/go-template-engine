<!-- Auto generated file, DO NOT EDIT. Please refer to hammer.yml -->
<!-- Alternatively you can set global properties, please refer to the docs on Hammer configuration -->
# go-template-engine

---

![LOGO](resources/gte-logo.png)

![shield](https://img.shields.io/github/release/marcelocorreia/go-template-engine.svg)
![shield](https://img.shields.io/github/last-commit/marcelocorreia/go-template-engine.svg)
![shield](https://img.shields.io/github/repo-size/marcelocorreia/go-template-engine.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcelocorreia/go-template-engine)](https://goreportcard.com/report/github.com/marcelocorreia/go-template-engine)

Easy tool to apply Go Templates in BAU jobs.

Based on Golang templates text SDK.

## TLDR;
- **Install**
	- Mac OS X
		- ```$ brew tap marcelocorreia/homebrew-taps; brew install go-template-engine```
	- Other Platforms
		- Download binaries from https://github.com/marcelocorreia/go-template-engine/releases
	- Docker way
		- cmd ```$ docker run --rm -it -v $\(pwd\):/app -w /app marcelocorreia/go-template-engine\"```
		- alias ```$ alias go-template-engine="docker run --rm marcelocorreia/go-template-engine"```
		- automated_alias_install ```$ curl -L https://github.com/marcelocorreia/go-template-engine/releases/download/{REPLACE_WITH_LASTEST_VERSION}/docker-alias-install.sh | sh```
- Added support ***AWS Secrets Manager***
    - Supports field Lookup
    - Tags
        - {{ secretsManagerField "myJsonSecret" "username"}} 		
        - {{ secretsManagerField "myJsonSecret" "password"}} 		
        - {{ secretsManager "myOtherSecret" }} 		
- Added support to [HCL](https://github.com/hashicorp/hcl) formart for variables file input
- Added support to all [Masterminds Sprig](https://github.com/Masterminds/sprig) functions
- Added static file include. Tag {{staticInclude "path/to/file.txt"}}
- Added replace tag. Tag {{replace .var "FROM_THIS" "TO_THIS"}}
- Accepts JSON and YAML variables files
- Lookup on file extension and parses accordingly, accepts **.json .yml .yaml** extensions
- Custom variable delimeter can be set using flags. Default: {{ , }}. Left and Right respectively. Check help menu.
- If --source points to a directory, it will run recursively, keeping the directory structure. Good for scaffolding
- Accepts multiple variables files, merging them. YAML only. (It will override duplicated variables if the exits in more than one file)
- These examples are pretty vanilla, go templates are actually pretty powerful, check the links for more info.
    - [https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html)
    - [https://golang.org/pkg/text/template/](https://golang.org/pkg/text/template/)
- Can be extended
    - ```$> go get github.com/marcelocorreia/go-template-engine/template-engine```

### Custom functions
|Function|Source|Desc|
|--------|------|----|
|secretsManager|GTE|na|
|secretsManagerField|GTE|na|
|abbrev|[sprig](https://github.com/Masterminds/sprig)|na|
|abbrevboth|[sprig](https://github.com/Masterminds/sprig)|na|
|add|[sprig](https://github.com/Masterminds/sprig)|na|
|add1|[sprig](https://github.com/Masterminds/sprig)|na|
|ago|[sprig](https://github.com/Masterminds/sprig)|na|
|append|[sprig](https://github.com/Masterminds/sprig)|na|
|atoi|[sprig](https://github.com/Masterminds/sprig)|na|
|b32dec|[sprig](https://github.com/Masterminds/sprig)|na|
|b32enc|[sprig](https://github.com/Masterminds/sprig)|na|
|b64dec|[sprig](https://github.com/Masterminds/sprig)|na|
|b64enc|[sprig](https://github.com/Masterminds/sprig)|na|
|base|[sprig](https://github.com/Masterminds/sprig)|na|
|biggest|[sprig](https://github.com/Masterminds/sprig)|na|
|buildCustomCert|[sprig](https://github.com/Masterminds/sprig)|na|
|camelcase|[sprig](https://github.com/Masterminds/sprig)|na|
|cat|[sprig](https://github.com/Masterminds/sprig)|na|
|ceil|[sprig](https://github.com/Masterminds/sprig)|na|
|clean|[sprig](https://github.com/Masterminds/sprig)|na|
|coalesce|[sprig](https://github.com/Masterminds/sprig)|na|
|compact|[sprig](https://github.com/Masterminds/sprig)|na|
|contains|[sprig](https://github.com/Masterminds/sprig)|na|
|date|[sprig](https://github.com/Masterminds/sprig)|na|
|dateInZone|[sprig](https://github.com/Masterminds/sprig)|na|
|dateModify|[sprig](https://github.com/Masterminds/sprig)|na|
|date_in_zone|[sprig](https://github.com/Masterminds/sprig)|na|
|date_modify|[sprig](https://github.com/Masterminds/sprig)|na|
|default|[sprig](https://github.com/Masterminds/sprig)|na|
|derivePassword|[sprig](https://github.com/Masterminds/sprig)|na|
|dict|[sprig](https://github.com/Masterminds/sprig)|na|
|dir|[sprig](https://github.com/Masterminds/sprig)|na|
|div|[sprig](https://github.com/Masterminds/sprig)|na|
|empty|[sprig](https://github.com/Masterminds/sprig)|na|
|env|[sprig](https://github.com/Masterminds/sprig)|na|
|expandenv|[sprig](https://github.com/Masterminds/sprig)|na|
|ext|[sprig](https://github.com/Masterminds/sprig)|na|
|fail|[sprig](https://github.com/Masterminds/sprig)|na|
|first|[sprig](https://github.com/Masterminds/sprig)|na|
|float64|[sprig](https://github.com/Masterminds/sprig)|na|
|floor|[sprig](https://github.com/Masterminds/sprig)|na|
|genCA|[sprig](https://github.com/Masterminds/sprig)|na|
|genPrivateKey|[sprig](https://github.com/Masterminds/sprig)|na|
|genSelfSignedCert|[sprig](https://github.com/Masterminds/sprig)|na|
|genSignedCert|[sprig](https://github.com/Masterminds/sprig)|na|
|has|[sprig](https://github.com/Masterminds/sprig)|na|
|hasKey|[sprig](https://github.com/Masterminds/sprig)|na|
|hasPrefix|[sprig](https://github.com/Masterminds/sprig)|na|
|hasSuffix|[sprig](https://github.com/Masterminds/sprig)|na|
|hello|[sprig](https://github.com/Masterminds/sprig)|na|
|htmlDate|[sprig](https://github.com/Masterminds/sprig)|na|
|htmlDateInZone|[sprig](https://github.com/Masterminds/sprig)|na|
|indent|[sprig](https://github.com/Masterminds/sprig)|na|
|initial|[sprig](https://github.com/Masterminds/sprig)|na|
|initials|[sprig](https://github.com/Masterminds/sprig)|na|
|int|[sprig](https://github.com/Masterminds/sprig)|na|
|int64|[sprig](https://github.com/Masterminds/sprig)|na|
|isAbs|[sprig](https://github.com/Masterminds/sprig)|na|
|join|[sprig](https://github.com/Masterminds/sprig)|na|
|keys|[sprig](https://github.com/Masterminds/sprig)|na|
|kindIs|[sprig](https://github.com/Masterminds/sprig)|na|
|kindOf|[sprig](https://github.com/Masterminds/sprig)|na|
|last|[sprig](https://github.com/Masterminds/sprig)|na|
|list|[sprig](https://github.com/Masterminds/sprig)|na|
|lower|[sprig](https://github.com/Masterminds/sprig)|na|
|max|[sprig](https://github.com/Masterminds/sprig)|na|
|merge|[sprig](https://github.com/Masterminds/sprig)|na|
|min|[sprig](https://github.com/Masterminds/sprig)|na|
|mod|[sprig](https://github.com/Masterminds/sprig)|na|
|mul|[sprig](https://github.com/Masterminds/sprig)|na|
|nindent|[sprig](https://github.com/Masterminds/sprig)|na|
|nospace|[sprig](https://github.com/Masterminds/sprig)|na|
|now|[sprig](https://github.com/Masterminds/sprig)|na|
|omit|[sprig](https://github.com/Masterminds/sprig)|na|
|pick|[sprig](https://github.com/Masterminds/sprig)|na|
|pluck|[sprig](https://github.com/Masterminds/sprig)|na|
|plural|[sprig](https://github.com/Masterminds/sprig)|na|
|prepend|[sprig](https://github.com/Masterminds/sprig)|na|
|push|[sprig](https://github.com/Masterminds/sprig)|na|
|quote|[sprig](https://github.com/Masterminds/sprig)|na|
|randAlpha|[sprig](https://github.com/Masterminds/sprig)|na|
|randAlphaNum|[sprig](https://github.com/Masterminds/sprig)|na|
|randAscii|[sprig](https://github.com/Masterminds/sprig)|na|
|randNumeric|[sprig](https://github.com/Masterminds/sprig)|na|
|regexFind|[sprig](https://github.com/Masterminds/sprig)|na|
|regexFindAll|[sprig](https://github.com/Masterminds/sprig)|na|
|regexMatch|[sprig](https://github.com/Masterminds/sprig)|na|
|regexReplaceAll|[sprig](https://github.com/Masterminds/sprig)|na|
|regexReplaceAllLiteral|[sprig](https://github.com/Masterminds/sprig)|na|
|regexSplit|[sprig](https://github.com/Masterminds/sprig)|na|
|repeat|[sprig](https://github.com/Masterminds/sprig)|na|
|replace|GTE|na|
|rest|[sprig](https://github.com/Masterminds/sprig)|na|
|reverse|[sprig](https://github.com/Masterminds/sprig)|na|
|round|[sprig](https://github.com/Masterminds/sprig)|na|
|semver|[sprig](https://github.com/Masterminds/sprig)|na|
|semverCompare|[sprig](https://github.com/Masterminds/sprig)|na|
|set|[sprig](https://github.com/Masterminds/sprig)|na|
|sha1sum|[sprig](https://github.com/Masterminds/sprig)|na|
|sha256sum|[sprig](https://github.com/Masterminds/sprig)|na|
|shuffle|[sprig](https://github.com/Masterminds/sprig)|na|
|snakecase|[sprig](https://github.com/Masterminds/sprig)|na|
|sortAlpha|[sprig](https://github.com/Masterminds/sprig)|na|
|split|[sprig](https://github.com/Masterminds/sprig)|na|
|splitList|[sprig](https://github.com/Masterminds/sprig)|na|
|squote|[sprig](https://github.com/Masterminds/sprig)|na|
|staticInclude|GTE|na|
|sub|[sprig](https://github.com/Masterminds/sprig)|na|
|substr|[sprig](https://github.com/Masterminds/sprig)|na|
|swapcase|[sprig](https://github.com/Masterminds/sprig)|na|
|ternary|[sprig](https://github.com/Masterminds/sprig)|na|
|title|[sprig](https://github.com/Masterminds/sprig)|na|
|toDate|[sprig](https://github.com/Masterminds/sprig)|na|
|toJson|[sprig](https://github.com/Masterminds/sprig)|na|
|toPrettyJson|[sprig](https://github.com/Masterminds/sprig)|na|
|toString|[sprig](https://github.com/Masterminds/sprig)|na|
|toStrings|[sprig](https://github.com/Masterminds/sprig)|na|
|trim|[sprig](https://github.com/Masterminds/sprig)|na|
|trimAll|[sprig](https://github.com/Masterminds/sprig)|na|
|trimPrefix|[sprig](https://github.com/Masterminds/sprig)|na|
|trimSuffix|[sprig](https://github.com/Masterminds/sprig)|na|
|trimall|[sprig](https://github.com/Masterminds/sprig)|na|
|trunc|[sprig](https://github.com/Masterminds/sprig)|na|
|tuple|[sprig](https://github.com/Masterminds/sprig)|na|
|typeIs|[sprig](https://github.com/Masterminds/sprig)|na|
|typeIsLike|[sprig](https://github.com/Masterminds/sprig)|na|
|typeOf|[sprig](https://github.com/Masterminds/sprig)|na|
|uniq|[sprig](https://github.com/Masterminds/sprig)|na|
|unset|[sprig](https://github.com/Masterminds/sprig)|na|
|until|[sprig](https://github.com/Masterminds/sprig)|na|
|untilStep|[sprig](https://github.com/Masterminds/sprig)|na|
|untitle|[sprig](https://github.com/Masterminds/sprig)|na|
|upper|[sprig](https://github.com/Masterminds/sprig)|na|
|uuidv4|[sprig](https://github.com/Masterminds/sprig)|na|
|without|[sprig](https://github.com/Masterminds/sprig)|na|
|wrap|[sprig](https://github.com/Masterminds/sprig)|na|
|wrapWith|[sprig](https://github.com/Masterminds/sprig)|na|


### Options
```
$> go-template-engine --help
  usage: go-template-engine --source=SOURCE [<flags>]

  Flags:
        --help                   Show context-sensitive help (also try --help-long and --help-man).
    -s, --source=SOURCE          Template Source File
        --var=VAR ...            Params & Variables. Example --var hey=ho --var lets=go
        --var-file=VAR-FILE ...  Variables File
        --exclude-dir=EXCLUDE-DIR ...
                                 Variables File
    -o, --output=OUTPUT          File output full path
        --delim-left="{{"        Left Delimiter
        --delim-right="}}"       Right Delimiter
    -v, --version                App Version
```

### Usage
files in the examples are located at template-engine/test_fixtures/
```
$> go-template-engine run --source template-engine/test_fixtures/cfn.tpl.yml \
        --var-file template-engine/test_fixtures/cfn-vars.yml \
        --output cfn.yml

$> go-template-engine run --source template-engine/test_fixtures/cfn.tpl.yml \
        --var-file template-engine/test_fixtures/cfn-vars.json \
        --output cfn.yml
```

### Install
#### Mac OS
```bash
$> brew tap marcelocorreia/homebrew-taps
   brew install go-template-engine
```
Other Systems Download latest binary from [https://github.com/marcelocorreia/go-template-engine/releases](https://github.com/marcelocorreia/go-template-engine/releases)

#### Docker
```bash
$ docker run --rm marcelocorreia/go-template-engine ...
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

#### Static Include
```yaml
---
list:
{{staticInclude "test_fixtures/list1.txt"}}
```
##### Result
```yaml
---
list:
  - hey
  - ho
  - lets
  - go
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

#### Simple vars passed on call
```bash
$> go-template-engine --source template-engine/test_fixtures/simple.txt.tpl \
    --var easy=simple \
    --var who=we
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
- [x] Accept multiple variable files
- [x] Recursive processing
- [x] Custom Delimeters
- [x] Static Include tag
- [x] Replace tag
- [x] Extra functions(tons.... thanks to [Masterminds Spring](https://github.com/Masterminds/sprig)

---









---
[:hammer:**Created with a Hammer**:hammer:](https://github.com/marcelocorreia/hammer)
<!-- Anchors -->





[linkedin]: https://www.linkedin.com/in/marcelocorreia
[website]: https://marcelo.correia.io
[slack]: https://correia-group.slack.com
[email]: marcelo@correia.io
[asciinema]: https://asciinema.org/~marcelocorreia
[ansible_galaxy_user]: marcelocorreia
<!-- end -->


