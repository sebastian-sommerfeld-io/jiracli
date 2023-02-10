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


readonly PIPELINE_MODE="--save-report"
MODE="--save-report"
if [ "$1" = "$PIPELINE_MODE" ]; then
  MODE="$1"
fi
readonly MODE


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
  readonly PORT="9080"
  LOG_HEADER "Run jetbrains/qodana"
  mkdir -p "$TARGET_DIR"
  mkdir -p "$TARGET_DIR/cache"

  if [ "$MODE" = "$PIPELINE_MODE" ]; then
    LOG_INFO "Run in pipeline mode"

    docker run --rm \
      --user "$(id -u):$(id -g)" \
      --volume "$(pwd):/data/project" \
      --volume "$(pwd)/$TARGET_DIR:/data/results" \
      --volume "$(pwd)/$TARGET_DIR/cache:/data/cache" \
      "$IMAGE" --save-report \
        --property=idea.suppressed.plugins.id=com.intellij.gradle
  else
    LOG_INFO "Run locally (http://localhost:$PORT)"

    docker run --rm -it -p "$PORT:8080" \
      --user "$(id -u):$(id -g)" \
      --volume "$(pwd):/data/project" \
      --volume "$(pwd)/$TARGET_DIR:/data/results" \
      --volume "$(pwd)/$TARGET_DIR/cache:/data/cache" \
      "$IMAGE" --show-report \
        --property=idea.suppressed.plugins.id=com.intellij.gradle
  fi
)
