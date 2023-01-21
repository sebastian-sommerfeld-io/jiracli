#!/bin/bash
# @file jiracli.sh
# @brief Build and run Jira CLI locally.
#
# @description The script is a wrapper for the `jiracli` command. The script builds and runs
# the ``jiracli`` locally. Before building, all tests are run as well. Everything go-related
# is delegated to the link:https://hub.docker.com/_/golang[golang] Docker image.
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


readonly TARGET_DIR="../../../target"

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../bash-modules/log.sh"
source "../bash-modules/go-wrapper.sh"


LOG_INFO "Run $0"
which go

# @description Format go source code.
function format() {
  LOG_HEADER "Format code"
  go fmt ./...
}


# @description Run all test cases.
function test() {
  LOG_HEADER "Run tests"
  go test ./...
}


# @description Run the app locally.
#
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
function run() {
  LOG_HEADER "Run app locally"
  go run . "$@"
}


# @description Build the app.
function build() {
  LOG_HEADER "Build app"
  go build .

  LOG_INFO "Move binary to target directory"
  mkdir -p "$TARGET_DIR"
  mv jiracli "$TARGET_DIR"
}


if [ ! -f go.mod ]; then
  readonly MODULE="sebastian-sommerfeld-io/jiracli"

  LOG_INFO "Initialize $MODULE"
  go mod init "$MODULE"
  go mod tidy

  go get -u github.com/spf13/cobra@latest
fi


format
test
run "$@"
build
