<!-- Auto generated file, DO NOT EDIT. Please refer to hammer.yml -->
<!-- Alternatively you can set global properties, please refer to the docs on Hammer configuration -->
{{- $github_user := .main_user.github_user}}{{- $name := .name }}
# {{ .name }}
{{ if .readme.banner }}![banner]({{.readme.banner}}){{ end }}
---
{{- range .readme.shields }}
[![shield]({{ . }})]({{ . }})
{{- end }}
{{ template "full_body" . }}
---
{{ template "tldr" . }}
{{ template "overview" . }}
{{ template "usage_file" . }}
{{ template "usage" . }}
{{ template "description" . }}
{{ template "description_file" . }}
{{ template "extra_content" . }}
{{ template "install" . }}
{{ template "created_with" . }}
<!-- Anchors -->
{{ template "anchors" . }}
<!-- end -->
{{ define "full_body" }}{{- if .readme.full_body }}
{{ .readme.full_body }}{{- end }}{{- end }}
{{ define "usage" }}{{- if .readme.usage }}
## Usage
{{ .readme.usage }}
{{- end }}{{- end}}

{{- define "usage_file" }}{{- if .readme.usage_file }}{{- $file := printf "%s/%s" .destination .readme.usage_file }}
{{ staticInclude $file }}{{- end }}{{- end}}

{{- define "tldr" }}{{- if .readme.tldr }}
## TLDR;{{ end }} 
{{- range .readme.tldr }}
- {{ . }}
{{- end }}{{- end }}

{{- define "overview" }}{{- if .readme.overview }}
## Overview
{{ .readme.overview }}
{{- end }}{{ end }} 

{{- define "description" }}{{- if .readme.description }}
## Description
{{ .readme.description }}
{{- end }}{{- end }}

{{- define "description_file" }}{{- if .readme.description_file }}{{- $file := printf "%s/%s" .destination .readme.description_file }}
{{- staticInclude $file }}{{- end }}{{- end }}

{{- define "extra_content" }}{{- if .readme.extra_content }}
{{ .readme.extra_content }}
{{- end }}{{- end }} 

{{- define "install" }}{{- if .readme.install }}
## Install
{{ .readme.install }}
{{- end -}}{{ end -}}

{{- define "created_with" }}
---
[:hammer:**Created with a Hammer**:hammer:](https://github.com/marcelocorreia/hammer)
{{- end }}

{{- define "anchors" }}
{{ if .main_user.logo }}[logo]: {{ .main_user.logo }}{{ end }}
{{ if .main_user.docs }}[docs]: {{ .main_user.docs }}{{ end }}
{{ if .main_user.github }}[github]: {{ .main_user.github }}{{ end }}
{{ if .main_user.dockerhub }}[dockerhub]: {{ .main_user.dockerhub }}{{ end }}
{{ if .main_user.linkedin }}[linkedin]: {{ .main_user.linkedin }}{{ end }}
{{ if .main_user.website }}[website]: {{ .main_user.website }}{{ end }}
{{ if .main_user.slack }}[slack]: {{ .main_user.slack }}{{ end }}
{{ if .main_user.email }}[email]: {{ .main_user.email }}{{ end }}
{{ if .main_user.asciinema }}[asciinema]: {{ .main_user.asciinema }}{{ end }}
{{ if .main_user.ansible_galaxy_user }}[ansible_galaxy_user]: {{ .main_user.ansible_galaxy_user }}{{ end }}
{{- end }}
