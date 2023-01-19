#!/bin/bash
# @file run.sh
# @brief Build and run Jira Ops CLI locally.
#
# @description The script builds and runs the ``jira-ops-cli`` locally. Testcases are run as well. Every thin go-related is delegated
# to the link:https://hub.docker.com/_/golang[golang] Docker image
#
# === Script Arguments
#
# The script does not accept any parameters.
#
# === Script Example
#
# [source, bash]
# ```
# ./run.sh
# ```


# Avoid 'unbound variable' errors in pipeline
# readonly LOG_ERROR="[\e[1;31mERROR\e[0m]"
# readonly LOG_INFO="[\e[34mINFO\e[0m]"


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
#    goversion
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


if [ ! -f go.mod ]; then
  readonly MODULE="sebastian-sommerfeld-io/jira-ops-cli"

  echo -e "$LOG_INFO Initialize $MODULE"
  go mod init "$MODULE"
  go mod tidy

  go get -u github.com/spf13/cobra@latest
fi

echo -e "$LOG_INFO Run tests"
go test ./...

echo -e "$LOG_INFO Run app"
#go build .
go run . "$@"