package provider

import (
	"io"
	"text/template"
)

type DOInstance struct {
	ResourceName string
	Image        string
	Name         string
	Region       string
	Size         string
}

func GenerateInstanceFile(instances []DOInstance, writer io.Writer) error {
	t, _ := template.New("droplet.tmpl").ParseFiles("../templates/digital_ocean/droplet.tmpl")
	return t.Execute(writer, instances)
}
