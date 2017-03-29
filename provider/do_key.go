package provider

import (
	"html/template"
	"io"
)

type DOKey struct {
	ResourceName string
	Name         string
	Path         string
}

func GenerateKeysFile(keys []DOKey, wr io.Writer) error {
	t, _ := template.New("keys.tmpl").ParseFiles("../templates/digital_ocean/keys.tmpl")
	return t.Execute(wr, keys)
}
