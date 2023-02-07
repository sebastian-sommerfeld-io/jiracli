#!/bin/bash
# @file jiracli.sh
# @brief Run Jira CLI locally.
#
# @description The script is a wrapper for the `jiracli` command. The script runs the
# ``jiracli`` app locally. Before building, all tests are run as well. Everything related
# to go is delegated to the link:https://hub.docker.com/_/golang[golang] Docker image.
#
# === Script Arguments
#
# * *$@* (...): The same parameters as ``jiracli`` app (see ``./jiracli.sh --help``)
#
# === Script Example
#
# [source, bash]
# ```
# ./jiracli.sh
# ```


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../../util/bash-modules/log.sh"
source "../../util/bash-modules/go-wrapper.sh"


# @description Format go source code.
function format() {
  LOG_HEADER "Format code"
  go fmt ./...
}


# @description Run all test cases and security scanner.
function test() {
  LOG_HEADER "Run tests"
  
  COVERAGE_REPORT="go-code-coverage.out"
  go test -coverprofile="./$COVERAGE_REPORT" ./...
  mv "$COVERAGE_REPORT" "../../../target/$COVERAGE_REPORT"
}


# @description Run the app locally.
#
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
function run() {
  LOG_HEADER "Run app locally"
  go run . "$@"
}


if [ ! -f go.mod ]; then
  readonly MODULE="github.com/sebastian-sommerfeld-io/jiracli"

  LOG_INFO "Initialize $MODULE"
  go mod init "$MODULE"
  go mod tidy

  go get -u github.com/spf13/cobra@latest
fi


format
test
run "$@"
