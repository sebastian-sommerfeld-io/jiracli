#!/bin/bash
# @file build.sh
# @brief Build the utility image containing the cobra-cli.
#
# @description The script builds the utility image containing the
# ``link:https://github.com/spf13/cobra-cli[cobra-cli]``.
#
# === Script Arguments
#
# The script does not accept any parameters.
#
# === Script Example
#
# [source, bash]
# ```
# ./build.sh
# ```


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


IMAGE="local/cobra-cli:dev"


# docker image rm --force "$IMAGE"
docker build -t "$IMAGE" .
# docker run --rm "$IMAGE" cobra-cli --help
