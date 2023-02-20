#!/bin/bash
# @file build-and-run.sh
# @brief Build, test and run the  Jira CLI locally.
#
# @description The script builds, tests and runs the `jiracli` app locally. The script runs all
# testcases, builds a Docker container and runs the app inside the Docker container.
#
# === Script Arguments
#
# * *$@* (...): The same parameters as ``jiracli`` app.
#
# === Script Example
#
# [source, bash]
# ```
# ./build-and-run.sh --help
# ```

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../util/bash-modules/log.sh"

readonly JIRACLI_IMAGE="local/jiracli:dev"


# @description Wrapper function to encapsulate ``go`` in a Docker container (``go`` commands
# are delegated to the link:https://hub.docker.com/_/golang[golang] Docker image).
#
# The current working directory is mounted into the container and selected as working directory
# so all files are available to ``go``. Paths are not preserved. The working directory is placed
# in ``/app`` (in the container) to make sure paths to the go app are the same everywhere (Go
# wrapper container, Dev Container and all images built from ``src/main/Dockerfile``). Keep in
# mind that most functions in this script (which call this ``go`` wrapper function) first ``cd``
# into the ``go`` folder. So most of the time the current working direktory is not ``src/main``
# (where this script is placed) but ``src/main/go``.
#
# The go wrapper container runs with the current user.
#
# @example
#    go version
#
# @arg $@ String The ``jiracli`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``go`` command is missing
function go() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the go container"
    LOG_ERROR "exit" && exit 8
  fi

  mkdir -p "/tmp/$USER/.cache"

  docker run --rm \
    --volume /etc/passwd:/etc/passwd:ro \
    --volume /etc/group:/etc/group:ro \
    --user "$(id -u):$(id -g)" \
    --volume "/tmp/$USER/.cache:/home/$USER/.cache" \
    --volume "$(pwd):/app" \
    --workdir "/app" \
    --network host \
    golang:1.20-rc-alpine go "$@"
}


# @description Format go source code. Before formatting, the function ``cd``s into the
# ``go`` folder.
function format() {
  LOG_HEADER "Format code"
  (
    cd go || exit
    go fmt ./...
  )
}


# @description Run all test cases and security scanner. The function copies the
# ``src/main/app-info.yml`` into the ``go`` folder to make sure it is mounted to the
# same location for all containers (go wrapper container, dev container, jiracli images).
# The temporairy ``src/main/go/app-info.yml`` is deleted when the tests are finished.
#
# Before testing, the function ``cd``s into the ``go`` folder.
function test() {
  LOG_HEADER "Run tests"
  
  cp app-info.yml go/app-info.yml
  
  (
    cd go || exit
  
    local TARGET_DIR="../../../target"
    local COVERAGE_REPORT="go-code-coverage.out"
    mkdir -p "$TARGET_DIR"

    go test -coverprofile="./$COVERAGE_REPORT" ./...

    old='github.com/sebastian-sommerfeld-io/jiracli'
    new='src/main/go'
    sed -i "s|$old|$new|g" "$COVERAGE_REPORT"
    mv "$COVERAGE_REPORT" "$TARGET_DIR/$COVERAGE_REPORT"
  )

  rm go/app-info.yml
}


# @description Build ``local/jiracli:dev`` Docker image.
function build() {
  LOG_HEADER "Build $JIRACLI_IMAGE Docker image"
  docker build -t "$JIRACLI_IMAGE" .
}

# @description Run ``jiracli`` app in Docker container.
#
# @arg $@ String The ``jiracli`` commands (1-n arguments) - $1 is mandatory
function run() {
  LOG_HEADER "Run app in Docker container" "$@"
  docker run --rm --network=host "$JIRACLI_IMAGE" "$@"
  echo
}


# @description Initialize the go application in needed. Before initializing, the function
# ``cd``s into the ``go`` folder.
function init() {
  (
    cd go || exit
    if [ ! -f go.mod ]; then
      local MODULE="github.com/sebastian-sommerfeld-io/jiracli"
      readonly MODULE

      LOG_HEADER "Initialize $MODULE"
      go mod init "$MODULE"
      go mod tidy

      go get -u github.com/spf13/cobra@latest
    fi
  )
}


init
format
test
build
run "$@"
