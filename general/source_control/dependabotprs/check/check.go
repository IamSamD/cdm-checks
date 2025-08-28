package check

import (
	"context"
	"fmt"

	"github.com/google/go-github/v74/github"
	"github.com/iamsamd/cdm_framework"
)

var log *cdm_framework.Logger = cdm_framework.NewLogger()

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

var ConfigValues []string = []string{
	"GITHUB_TOKEN",
	"GITHUB_OWNER",
	"GITHUB_REPO",
}

/*
Check is the function in which you write your check.

All env vars you declared in ConfigValues are accessible in the config object.

Example:

	config["ENV_VAR_ONE"]
*/

func Check(config cdm_framework.Config) error {
	// Initialize the github client
	client := github.NewClient(nil).WithAuthToken(config["GITHUB_TOKEN"])

	// Fetch Dependabot PRs
	prs, _, err := client.PullRequests.List(context.Background(), config["GITHUB_OWNER"], config["GITHUB_REPO"], &github.PullRequestListOptions{
		State: "open",
		Base:  "main",
	})
	if err != nil {
		return fmt.Errorf("failed to list github PRs due to error: %v", err)
	}

	log.Debug(fmt.Sprintf("Found %d open PRs for repo: %s", len(prs), config["GITHUB_REPO"]))

	// Filter for dependabot PRs
	var dependabotPRs []*github.PullRequest

	for _, pr := range prs {
		if pr.User.GetLogin() == "dependabot[bot]" {
			dependabotPRs = append(dependabotPRs, pr)
		}
	}

	log.Debug(fmt.Sprintf("Found %d open dependabot PRs for repo: %s", len(dependabotPRs), config["GITHUB_REPO"]))

	// Check assertion - fail check if more than 5 open Dependabot PRs
	if len(dependabotPRs) >= 5 {
		cdm_framework.CheckFailedReport(fmt.Sprintf("There are more than 5 open Dependabot PRs for repo: %s", config["GITHUB_REPO"]))
		cdm_framework.FailCheck()
	}

	return nil
}
