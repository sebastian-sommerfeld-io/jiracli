#!/bin/bash
# @file run-local.sh
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
# ./run-local.sh --help
# ```


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../util/bash-modules/log.sh"
source "../util/bash-modules/go-wrapper.sh"


# @description Format go source code.
function format() {
  LOG_HEADER "Format code"
  (
    cd go || exit
    go fmt ./...
  )
}


# @description Run all test cases and security scanner.
function test() {
  LOG_HEADER "Run tests"
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
}


# @description Build Docker image and run the app locally.
#
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
function build_and_run() {
  LOG_HEADER "Build Docker image and run app locally"
  readonly IMAGE="local/jiracli:dev"

  docker build -t "$IMAGE" .
  docker run --rm --network=host "$IMAGE" "$@"
}


(
  cd go || exit
  if [ ! -f go.mod ]; then
    readonly MODULE="github.com/sebastian-sommerfeld-io/jiracli"

    LOG_INFO "Initialize $MODULE"
    go mod init "$MODULE"
    go mod tidy

    go get -u github.com/spf13/cobra@latest
  fi
)


format
test
build_and_run "$@"
