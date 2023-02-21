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


readonly BASE_URL="http://localhost:8080"
readonly USER="admin"
readonly PASS="admin"


# curl --location --request GET "$BASE_URL/rest/plugins/applications/1.0/installed/jira-software/license" \
#   --user "$USER:$PASS" \
#   --header 'Content-Type: application/json'

httpCode=$(curl --location --request GET "$BASE_URL/rest/api/2/serverInfo?doHealthCheck=true" \
  -s -o /dev/null -w "%{http_code}" \
  --user "$USER:$PASS" \
  --header 'Content-Type: application/json')
echo "$httpCode"
