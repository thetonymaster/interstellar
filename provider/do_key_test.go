package provider

import "os"

func ExampleGenerateKeysFile() {
	keys := []DOKey{
		{ResourceName: "default",
			Name: "Terraform Example",
			Path: "/Users/terraform/.ssh/id_rsa.pub"},
	}

	GenerateKeysFile(keys, os.Stdout)

	// Output: resource "digitalocean_ssh_key" "default" {
	//   name       = "Terraform Example"
	//   public_key = "${file("/Users/terraform/.ssh/id_rsa.pub")}"
	// }

}
