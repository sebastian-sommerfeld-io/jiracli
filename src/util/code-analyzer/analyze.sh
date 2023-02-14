#!/bin/bash
# @file analyze.sh
# @brief Analyze this projects source code locally by running some linters and analyzers.
#
# @description This script analyzes this projects source code by running some linters and
# analyzers. The script is intended to run on your local machine.
#
# TODO * Sonarlint
# * link:https://www.jetbrains.com/de-de/qodana[Qodana] -> Docker image link:https://hub.docker.com/r/jetbrains/qodana[``jetbrains/qodana``]
#
# === Script Arguments
#
# * *$1* (string): Use ``--save-report`` to run in pipelines. When omitting this option a webserver starts at link:http://localhost:9080[localhost:9080].
#
# === Script Example
#
# [source, bash]
# ```
# ./analyze.sh
# ```


readonly PORT="9080"
readonly PIPELINE_MODE="--save-report"
MODE="--show-report"
FLAGS="--rm -it -p $PORT:8080"
if [ "$1" = "$PIPELINE_MODE" ]; then
  MODE="$1"
  FLAGS="--rm"
fi
readonly MODE
readonly FLAGS


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../bash-modules/log.sh"

readonly IMAGE="local/qodana-go:dev"


LOG_HEADER "Build analyzer image"
docker build -t "$IMAGE" .

(
  cd ../../../ || exit

  readonly TARGET_DIR="target/qodana"
  LOG_HEADER "Run jetbrains/qodana"
  mkdir -p "$TARGET_DIR"
  mkdir -p "$TARGET_DIR/cache"

  LOG_INFO "Run Qodana (mode = '$MODE', flags = '$FLAGS')"
  # shellcheck disable=SC2086
  docker run $FLAGS \
    --user "$(id -u):$(id -g)" \
    --volume "$(pwd):/data/project" \
    --volume "$(pwd)/$TARGET_DIR:/data/results" \
    --volume "$(pwd)/$TARGET_DIR/cache:/data/cache" \
    "$IMAGE" "$MODE" \
      --property=idea.suppressed.plugins.id=com.intellij.gradle
)
