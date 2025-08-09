package check

import (
	"log/slog"

	"github.com/iamsamd/cdm_framework"
)

var log *slog.Logger = cdm_framework.Logger

/*
ConfigValues is a slice of strings, each string being the name of
an environment variable your check depends on for it's configuration.

Add the names of the env vars your check needs to access, to ConfigValues below.
You can add as many as you need.

Example:

	var ConfigValues []string = []string{
		"ENV_VAR_ONE",
		"ENV_VAR_TWO",
		"ENV_VAR_THREE",
	}
*/

var ConfigValues []string = []string{}

/*
Check is the function in which you write your check.

All env vars you declared in ConfigValues are accessible in the config object.

Example:

	config["ENV_VAR_ONE"]
*/

func Check(config cdm_framework.Config) error {
	// Write your check logic here
}
