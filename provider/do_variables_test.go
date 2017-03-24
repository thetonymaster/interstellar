package provider

import "os"

func ExampleGenerateVariables() {
	variables := DOVariables{
		DOAuthentication: DOAuthentication{
			Token: "77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3",
		},
	}

	GenerateVariables(variables, os.Stdout)

	// Output: DO_TOKEN = "77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3"
}
