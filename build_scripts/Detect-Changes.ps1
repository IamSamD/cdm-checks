[CmdletBinding()]
param (
    [Parameter(Mandatory)]
    [string]
    $isPr
)

Set-Location "../"

$ErrorActionPreference = "Stop"

Write-Output "Running git diff...`n"

if ($isPr -eq "true") {
    $files = git diff --name-only origin/main HEAD HEAD
} else {
    $files = git diff --name-only HEAD~1 HEAD
}

Write-Output "Changed files: $files`n"

$paths = $files -split "\n"

$parentDirs = @()

Write-Output "Detecting checks for building..."
foreach ($path in $paths) {
    $parentDir = $path

    while ($true) {
        # Write-Host $parentDir
        if (Test-Path "$parentDir/go.mod") {
            $parentDirs += $parentDir
            break
        }

        if ($null -ne $parentDir -and $parentDir -ne "") {
            $parentDir = Split-Path $parentDir -Parent
            continue
        }

        break
    }
}

if ($parentDirs.Count -eq 0) {
    Write-Output "No checks to build"
    "skipbuild=true" >> $env:GITHUB_OUTPUT
    exit 0
}

$parentDirs = $parentDirs | Sort-Object -Unique

Write-Output "Checks to be built: $parentDirs"

"checks=$($parentDirs | ConvertTo-Json -AsArray -Compress)" >> $env:GITHUB_OUTPUT