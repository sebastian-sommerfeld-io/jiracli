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


readonly LOG_INFO="[\e[34mINFO\e[0m]" # Avoid 'unbound variable' errors in pipeline
readonly TARGET_DIR="../../../target"

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
bash .go-wrapper.sh


echo -e "$LOG_INFO Run $0"


# @description Print log output in a header-style.
#
# @arg $@ String The line to print.
function LOG() {
  echo -e "$LOG_INFO ------------------------------------------------------------------------"
  echo -e "$LOG_INFO $1"
  echo -e "$LOG_INFO ------------------------------------------------------------------------"
}


# @description Format go source code.
function format() {
  LOG "Format code"
  go fmt ./...
}


# @description Run all test cases.
function test() {
  LOG "Run tests"
  go test ./...
}


# @description Run the app locally.
#
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
function run() {
  LOG "Run app locally"
  go run . "$@"
}


# @description Build the app.
function build() {
  LOG "Build app"
  go build .

  echo -e "$LOG_INFO Move binary to target directory"
  mkdir -p "$TARGET_DIR"
  mv jiracli "$TARGET_DIR"
}


if [ ! -f go.mod ]; then
  readonly MODULE="sebastian-sommerfeld-io/jiracli"

  echo -e "$LOG_INFO Initialize $MODULE"
  go mod init "$MODULE"
  go mod tidy

  go get -u github.com/spf13/cobra@latest
fi


format
test
run "$@"
build
