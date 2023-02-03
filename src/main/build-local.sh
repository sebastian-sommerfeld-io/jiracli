#!/bin/bash
# @file build-local.sh
# @brief Build Jira CLI locally.
#
# @description The script builds the ``jiracli`` locally. Before building, first run the app by
# invoking ``xref:AUTO-GENERATED:bash-docs/src/main/go/jiracli-sh.adoc[src/main/go/jiracli.sh]``.
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

IMAGE_NAME="local/jiracli:dev"

# Include local bash modules
source "../util/bash-modules/log.sh"

# @description Build the app inside a Docker image.
function buildDockerImage() {
  LOG_HEADER "Build Docker image"
  docker build -t "$IMAGE_NAME" .
}

buildDockerImage

LOG_HEADER "Test Docker image"
docker run --rm "$IMAGE_NAME" user view admin.admin
