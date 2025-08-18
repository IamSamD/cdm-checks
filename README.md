# cdm-checks

Shared library for CDM checks to live

All CDM checks are written in this repository as small Go programs whch compile to binaries.

The [cdm-cli](https://github.com/IamSamD/cdm-cli) can treat these binaries as plugins and therefore can execute any check available to it.

## Repository Structure

Checks are organised in a hierarchical structure as shown below:

```bash
.
└── cloud_provider
    └── cloud_resource
        └── check_name
```

For example, a check that notifies for an Azure Kubernetes Service upgrade would be structured as below:

```bash
.
└── azure
    └── aks
        └── upgrade
```

Where the `azure/aks/upgrade` directroy hold the code for the check.

## Checks

A check is written using a standard template which can be generated using the [cdm-cli](https://github.com/IamSamD/cdm-cli).

The structure for a check:

```bash
.
├── check
│   └── check.go
├── go.mod
├── go.sum
└── main.go
```

Once you have generated a new check, the only file you need to edit is `check/check.go`.
This is wher you define the config values your check requires and the logic for your check.

```go
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

var ConfigValues []string = []string{
    // Declare your env var names here
}

/*
Check is the function in which you write your check.

All env vars you declared in ConfigValues are accessible in the config object.

Example:

	config["ENV_VAR_ONE"]
*/

func Check(config cdm_framework.Config) error {
	// Write your check logic here
}
```

This ensures that as a developer writing new checks, you do not need to be concerned with the checks execution,
only the logic of your check maintaining a more simple developer experience.

### Testing locally

The `general/source_control/depandabit_prs` check will be used as an example for this section.

Checks use godotenv to enable local testing of a check

To test a check locally create a `.env` file in the root dir of your check:

```bash
touch ./general/source_control/dependabot_prs/.env
```

Set the required env vars in the `.env` file:

```bash
GITHUB_TOKEN="<GITHUB_PAT>"
GITHUB_OWNER="<REPOSITORY_OWNER>"
GITHUB_REPO="<REPOSITORY_NAME>"
```

You can now run the check locally with

```bash
go run main.go
```

## CI
The CI pipeline for this repo is a Github Actions Workflow that will dynamically detect check with new commits and run a build only on the those checks. 

### PR
On raising a PR the CI workflow will trigger

Checks with new commits will be detected against the main branch

Only checks with new commits will be built to check that the build succeeds

Once a successful build has been verified you will be cleared to merge to main

### Merge to main
On merge to main the workflow will:

- Build the new/updated check
- Bump the version of the check
- Create a new tag and a new release for the check on the new version

This will allow for specifying the version of a check that should be used in the clients CDM pipeline. 