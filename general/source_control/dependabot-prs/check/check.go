package check

import (
	"context"
	"fmt"

	"github.com/google/go-github/v74/github"
	"github.com/iamsamd/cdm_framework"
)

// Required configuration values
// - `GITHUB_TOKEN`: GitHub personal access token with permissions to read repository data
// - `GITHUB_OWNER`: GitHub repository owner
// - `GITHUB_REPO`: GitHub repository name

var ConfigValues []string = []string{
	// Place your Env Var names here to load all config values
	"GITHUB_TOKEN",
	"GITHUB_OWNER",
	"GITHUB_REPO",
}

func RunCheck(config map[string]string) error {
	/*
		--------------------------------
		 Write your check logic here
		--------------------------------
	*/
	log := cdm_framework.Logger

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

	log.Info(fmt.Sprintf("Found %d open dependabot PRs for repo: %s", len(dependabotPRs), config["GITHUB_REPO"]))

	// Check assertion - fail check if more than 50 open Dependabot PRs
	if len(dependabotPRs) > 50 {
		log.Info(fmt.Sprintf("There are more than 50 open Dependabot PRs for repo: %s", config["GITHUB_REPO"]))
		cdm_framework.FailCheck()
	}

	return nil
}
