// Code generated by gentmpl.sh; DO NOT EDIT.

package gomarkdoc

var templates = map[string]string{
	"doc": `{{- range .Blocks -}}
	{{- if eq .Kind "paragraph" -}}
		{{- paragraph .Text -}}
	{{- else if eq .Kind "code" -}}
		{{- codeBlock "" .Text -}}
	{{- else if eq .Kind "header" -}}
		{{- header .Level .Text -}}
	{{- end -}}
{{- end -}}

`,
	"example": `{{- accordionHeader .Title -}}

{{- template "doc" .Doc -}}

{{- codeBlock "go" .Code -}}

{{- if .HasOutput -}}

	{{- header 4 "Output" -}}

	{{- codeBlock "" .Output -}}
    
{{- end -}}

{{- accordionTerminator -}}

`,
	"file": `<!-- Code generated by gomarkdoc. DO NOT EDIT -->

{{.Header -}}

{{- range .Packages -}}
	{{- template "package" . -}}
{{- end -}}

{{- .Footer}}

Generated by {{link "gomarkdoc" "https://github.com/huner2/gomarkdoc"}}
`,
	"func": `{{- if .Receiver -}}
	{{- codeHref .Location | link (escape .Name) | printf "func \\(%s\\) %s" (escape .Receiver) | rawHeader .Level -}}
{{- else -}}
	{{- codeHref .Location | link (escape .Name) | printf "func %s" | rawHeader .Level -}}
{{- end -}}

{{- codeBlock "go" .Signature -}}

{{- template "doc" .Doc -}}

{{- range .Examples -}}
	{{- template "example" . -}}
{{- end -}}

`,
	"index": `{{- if len .Consts -}}

	{{- localHref "Constants" | link "Constants" | listEntry 0 -}}
	
{{- end -}}

{{- if len .Vars -}}

	{{- localHref "Variables" | link "Variables" | listEntry 0 -}}

{{- end -}}

{{- range .Funcs -}}

	{{- if .Receiver -}}
		{{- codeHref .Location | link (escape .Name) | printf "func \\(%s\\) %s" (escape .Receiver) | localHref | link .Signature | listEntry 0 -}}
	{{- else -}}
		{{- codeHref .Location | link (escape .Name) | printf "func %s" | localHref | link .Signature | listEntry 0 -}}
	{{- end -}}

{{- end -}}

{{- range .Types -}}

	{{- codeHref .Location | link (escape .Name) | printf "type %s" | localHref | link .Title | listEntry 0 -}}

	{{- range .Funcs -}}
		{{- if .Receiver -}}
			{{- codeHref .Location | link (escape .Name) | printf "func \\(%s\\) %s" (escape .Receiver) | localHref | link .Signature | listEntry 1 -}}
		{{- else -}}
			{{- codeHref .Location | link (escape .Name) | printf "func %s" | localHref | link .Signature | listEntry 1 -}}
		{{- end -}}
	{{- end -}}

	{{- range .Methods -}}
		{{- if .Receiver -}}
			{{- codeHref .Location | link (escape .Name) | printf "func \\(%s\\) %s" (escape .Receiver) | localHref | link .Signature | listEntry 1 -}}
		{{- else -}}
			{{- codeHref .Location | link (escape .Name) | printf "func %s" | localHref | link .Signature | listEntry 1 -}}
		{{- end -}}
	{{- end -}}

{{- end -}}

{{- spacer -}}
`,
	"package": `{{- if eq .Name "main" -}}
	{{- header .Level .Dirname -}}
{{- else -}}
	{{- header .Level .Name -}}
{{- end -}}

{{- codeBlock "go" .Import -}}

{{- template "doc" .Doc -}}

{{- range .Examples -}}
	{{- template "example" . -}}
{{- end -}}

{{- header (add .Level 1) "Index" -}}

{{- template "index" . -}}

{{- if len .Consts -}}

	{{- header (add .Level 1) "Constants" -}}

	{{- range .Consts -}}
		{{- template "value" . -}}
	{{- end -}}

{{- end -}}

{{- if len .Vars -}}

	{{- header (add .Level 1) "Variables" -}}

	{{- range .Vars -}}
		{{- template "value" . -}}
	{{- end -}}

{{- end -}}

{{- range .Funcs -}}
	{{- template "func" . -}}
{{- end -}}

{{- range .Types -}}
	{{- template "type" . -}}
{{- end -}}
`,
	"type": `{{- codeHref .Location | link (escape .Name) | printf "type %s" | rawHeader .Level -}}

{{- template "doc" .Doc -}}

{{- codeBlock "go" .Decl -}}

{{- range .Consts -}}
	{{- template "value" . -}}
{{- end -}}

{{- range .Vars -}}
	{{- template "value" . -}}
{{- end -}}

{{- range .Examples -}}
	{{- template "example" . -}}
{{- end -}}

{{- range .Funcs -}}
	{{- template "func" . -}}
{{- end -}}

{{- range .Methods -}}
	{{- template "func" . -}}
{{- end -}}

`,
	"value": `{{- template "doc" .Doc -}}

{{- codeBlock "go" .Decl -}}

`,
}
