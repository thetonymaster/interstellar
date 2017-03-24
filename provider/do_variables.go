package provider

import (
	"io"
	"text/template"
)

type DOVariables struct {
	DOAuthentication DOAuthentication
}

func GenerateVariables(variables DOVariables, wr io.Writer) error {
	t, _ := template.New("vars.tmpl").ParseFiles("../templates/digital_ocean/vars.tmpl")
	return t.Execute(wr, variables)
}
