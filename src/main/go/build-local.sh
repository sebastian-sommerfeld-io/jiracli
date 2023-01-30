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

# Include local bash modules
source "../../util/bash-modules/log.sh"
source "../../util/bash-modules/go-wrapper.sh"


# @description Build the app.
function buildGoApp() {
  local TARGET_DIR="../../../target"
  readonly TARGET_DIR

  LOG_HEADER "Build go app"
  go build .

  LOG_INFO "Move binary to target directory"
  mkdir -p "$TARGET_DIR"
  mv jiracli "$TARGET_DIR"
}


if [ ! -f go.mod ]; then
  LOG_ERROR "No go.mod file found. To generate mandatory files, first run the app locally!"
  LOG_ERROR "exit" && exit 8
fi


buildGoApp
