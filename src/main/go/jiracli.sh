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
source "../utils/bash-modules/log.sh"
source "../utils/bash-modules/go-wrapper.sh"


# @description Format go source code.
function format() {
  LOG_HEADER "Format code"
  go fmt ./...
}


# @description Run local linters.
function linters() {
  LOG_HEADER "Run local linters"

  # (
  #   cd ../../../ || exit

  #   mkdir -p target/cache/.m2

  #   docker run --rm -it \
  #     --volume /etc/passwd:/etc/passwd:ro \
  #     --volume /etc/group:/etc/group:ro \
  #     --user "$(id -u):$(id -g)" \
  #     --volume "$(pwd)/src/main/go:/data/project" \
  #     --volume "$(pwd)/target:/data/results" \
  #     --volume "$(pwd)/target/cache:/data/cache" \
  #     jetbrains/qodana-go:2022.3-eap --save-report --user "$(id -u):$(id -g)"
  # )
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


if [ ! -f go.mod ]; then
  readonly MODULE="github.com/sebastian-sommerfeld-io/jiracli"

  LOG_INFO "Initialize $MODULE"
  go mod init "$MODULE"
  go mod tidy

  go get -u github.com/spf13/cobra@latest
fi


format
linters
test
run "$@"
