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
# The script does not accept any parameters.
#
# === Script Example
#
# [source, bash]
# ```
# ./analyze.sh
# ```


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

  LOG_HEADER "Run yamllint"
  docker run -it --rm \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    cytopia/yamllint:latest .

  readonly TARGET_DIR="target/qodana"
  readonly PORT="9080"
  LOG_HEADER "Run jetbrains/qodana on http://localhost:$PORT"
  mkdir -p "$TARGET_DIR"
  mkdir -p "$TARGET_DIR/cache"

  docker run --rm -it -p 9080:8080 \
    --user "$(id -u):$(id -g)" \
    --volume "$(pwd):/data/project" \
    --volume "$(pwd)/$TARGET_DIR:/data/results" \
    --volume "$(pwd)/$TARGET_DIR/cache:/data/cache" \
    "$IMAGE" --show-report
)
