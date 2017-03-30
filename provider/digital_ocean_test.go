package provider

import (
	"os"
	"testing"
)

func TestWriteInstanceFiles(t *testing.T) {
	instances := []DOInstance{
		{ResourceName: "some content", Image: "some content", Name: "some content",
			Region: "some content", Size: "some content"},
	}

	err := writeInstanceFiles(instances)
	if err != nil {
		t.Error("Expected no error, got: ", err.Error())
	}
	os.Remove("instances.tf")
}

func TestWriteVarsFiles(t *testing.T) {
	variables := DOVariables{
		DOAuthentication: DOAuthentication{
			Token: "some content",
		},
	}

	err := writeVarsFiles(variables)
	if err != nil {
		t.Error("Expected no error, got: ", err.Error())
	}
	os.Remove("vars.tfvars")
}

func TestWriteKeyFiles(t *testing.T) {
	keys := []DOKey{
		{ResourceName: "some content", Name: "some content", Path: "some content"},
	}

	err := writeKeyfiles(keys)
	if err != nil {
		t.Error("Expected no error, got: ", err.Error())
	}
	os.Remove("keys.tf")
}
