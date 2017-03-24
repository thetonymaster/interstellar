package provider

import "os"

func ExampleGenerateInstanceFile() {
	instances := []DigitalOceanInstance{{
		ResourceName: "web",
		Name:         "web-1",
		Image:        "ubuntu-14-04-x64",
		Region:       "nyc2",
		Size:         "512mb",
	}}
	GenerateInstanceFile(instances, os.Stdout)

	// Output: resource "digitalocean_droplet" "web" {
	//   image  = "ubuntu-14-04-x64"
	//   name   = "web-1"
	//   region = "nyc2"
	//   size   = "512mb"
	// }
}
