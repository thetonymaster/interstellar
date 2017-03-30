package provider

import (
	"bufio"
	"os"
)

type DigitalOcean struct {
	Instances []DOInstance
	Variables DOVariables
	Keys      []DOKey
}

func (do DigitalOcean) WriteFiles() error {
	return nil
}

func writeInstanceFiles(instances []DOInstance) error {
	f, err := os.Create("instances.tf")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = GenerateInstanceFile(instances, w)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func writeVarsFiles(variables DOVariables) error {
	f, err := os.Create("vars.tfvars")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = GenerateVariables(variables, w)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func writeKeyfiles(keys []DOKey) error {
	f, err := os.Create("keys.tf")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = GenerateKeysFile(keys, w)
	if err != nil {
		return err
	}
	w.Flush()
	return nil

}
