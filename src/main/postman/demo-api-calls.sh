#!/bin/bash
# @file demo-api-calls.sh
# @brief Invoke some demo calls to the Jira API to learn about the Rest API.
#
# @description This script invokes some demo calls to the Jira API to learn about the Rest API.
#
# === Script Arguments
#
# The script does not accept any parameters.
#
# === Script Example
#
# [source, bash]
# ```
# ./demo-api-calls.sh
# ```


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


# Include local bash modules
source "../../util/bash-modules/log.sh"


# @description Wrapper function to not repeat myself when calling the Jira API lots of times.
#
# @example
#    callEndpoint /rest/plugins/applications/1.0/installed/jira-software/license
#
# @arg $@ String The API endpoint - mandatory
#
# @exitcode 8 If param API endpoint string is missing
function callEndpoint() {
  if [ -z "$1" ]; then
    LOG_ERROR "API endpoint missing"
    LOG_ERROR "exit" && exit 8
  fi

  readonly USER="admin"
  readonly PASS="admin"
  readonly BASE_URL="http://localhost:8080"

  LOG_HEADER "Call API endpoint $1"
  LOG_INFO "User = $USER"

  curl --location --request GET "${BASE_URL}${1}" \
    -u "$USER:$PASS" \
    --header 'Content-Type: application/json' # | json_pp
}

callEndpoint /rest/plugins/applications/1.0/installed/jira-software/license
