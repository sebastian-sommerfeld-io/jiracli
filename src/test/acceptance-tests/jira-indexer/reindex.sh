#!/bin/bash
# @file reindex.sh
# @brief Create a Jira index in the foreground after starting Jira.
#
# @description When starting the Jira instance that acts as the test environment the Jira
# instance is provided with a h2db and is populated with a license, users and projects so
# there is no need  re-run the setup wizard all the time.
#
# The only issue is, that there is no data visible  in the UI after starting up because 
# there is no (or at least no uncorrupted) index. A new index must be created in the foreground.
# To avoid manual interaction, this is done in a dedicated container running this script. This
# script calls the Jira API to trigger a re-index. 
#
# IMPORTANT: This script is intended to run inside a ``jira-indexer`` container from Docker
# Compose. Don't run this script directly!
#
# === Script Arguments
#
# The script does not accept any parameters.


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


readonly BASE_URL="http://jira:8080"
readonly USER="admin"
readonly PASS="admin"


sleep 5m


curl --location --request POST "$BASE_URL/rest/api/2/reindex?type=FOREGROUND" \
  --user "$USER:$PASS" \
  --header 'Accept: application/json'


sleep 10s


curl --location --request GET "$BASE_URL/rest/api/2/reindex" \
  --user "$USER:$PASS" \
  --header 'Content-Type: application/json'
  