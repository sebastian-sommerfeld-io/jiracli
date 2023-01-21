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


# Avoid 'unbound variable' errors in pipeline
readonly LOG_ERROR="[\e[1;31mERROR\e[0m]"
readonly LOG_INFO="[\e[34mINFO\e[0m]"

readonly TARGET_DIR="../../../target"

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


echo -e "$LOG_INFO Run $0"


# @description Wrapper function to encapsulate ``go`` in a docker container. The current working
# directory is mounted into the container and selected as working directory so that all file are
# available to ``go``. Paths are preserved. The container runs with the current user.
#
# @example
#    go version
#
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with terraform command is missing
function go() {
  if [ -z "$1" ]; then
    echo -e "$LOG_ERROR No command passed to the go container"
    echo -e "$LOG_ERROR exit" && exit 8
  fi

  mkdir -p "/tmp/$USER/.cache"

  docker run --rm \
    --volume /etc/passwd:/etc/passwd:ro \
    --volume /etc/group:/etc/group:ro \
    --user "$(id -u):$(id -g)" \
    --volume "/tmp/$USER/.cache:/home/$USER/.cache" \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    golang:1.20-rc-alpine go "$@"
}


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
