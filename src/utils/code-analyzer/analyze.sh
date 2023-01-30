#!/bin/bash
# @file analyze.sh
# @brief Analyze this projects source code locally by running some linters and analyzers.
#
# @description This script analyzes this projects source code by running some linters and
# analyzers. The script is intended to run on your local machine.
#
# TODO * Sonarlint
# * link:https://www.jetbrains.com/de-de/qodana[Qodana] -> Docker image `link:https://hub.docker.com/r/jetbrains/qodana[jetbrains/qodana]`
#
# === Script Arguments
#
# The script does not accept any parameters.
# * *$1* (string): Some var
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


(
  cd ../../../ || exit

  LOG_HEADER "Run yamllint"
  docker run -it --rm \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    cytopia/yamllint:latest .


  LOG_HEADER "Run jetbrains/qodana"
  readonly TARGET="target/qodana"
  
  mkdir -p "$TARGET"

#   docker run --rm -it -p 8080:8080 \
#     --volume "$(pwd):/data/project" \
#     --volume "$(pwd)/$TARGET:/data/results" \
#     --volume "$(pwd)/$TARGET/cache:/data/cache" \
#     jetbrains/qodana:2022.2-eap --show-report
  docker run --rm -it -p 8080:8080 \
    --volume "$(pwd):/data/project" \
    --volume "$(pwd)/$TARGET:/data/results" \
    --volume "$(pwd)/$TARGET/cache:/data/cache" \
    jetbrains/qodana-go:2022.3-eap --show-report
)
