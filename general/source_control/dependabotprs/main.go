package main

import (
	"dependabotprs/check"

	"github.com/iamsamd/cdm_framework"
)

var config cdm_framework.Config

var log *cdm_framework.Logger = cdm_framework.NewLogger()

func init() {
	configValues, err := cdm_framework.InitEnvironment(check.ConfigValues)
	if err != nil {
		log.Error(err.Error())
		cdm_framework.FailCheck()
	}
	config = configValues
}

func main() {
	if err := cdm_framework.RunCheck(config, check.Check); err != nil {
		log.Error(err.Error())
		cdm_framework.FailCheck()
	}
}
