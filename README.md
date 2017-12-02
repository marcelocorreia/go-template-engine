# go-template-engine GTE

Based on Golang templates text SDK. 
- [https://golang.org/pkg/text/template/](https://golang.org/pkg/text/template/)
   

##TLDR;

- Accepts JSON and YAML variable file
- Looks up file extension, accepts **.json .yml .yaml** extensions
- 

###Usage
```
$> go-template-engine --source README.md.tpl \
        -var-file template-engine/test_fixtures/vars.yml
```

###Example

#####Template
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
