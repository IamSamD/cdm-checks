package main

import (
	"dependabotprs/check"
	"fmt"
	"log"

	"github.com/iamsamd/cdm_framework"
	"github.com/joho/godotenv"
)

var config map[string]string

func init() {
	// Use godotenv for testing locally
	goDotEnvConfig, err := godotenv.Read()
	if err != nil {
		configValuesMap, err := cdm_framework.LoadEnvVars(check.ConfigValues)
		if err != nil {
			log.Fatal(err)
			// TODO: Ensure fail on error does not raise Ado ticket
			cdm_framework.FailCheck()
		}

		config = configValuesMap
		return
	}

	config = goDotEnvConfig
}

func main() {
	if err := check.RunCheck(config); err != nil {
		fmt.Println(err)
		cdm_framework.FailCheck()
	}
}
