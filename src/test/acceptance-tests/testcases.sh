#!/bin/bash
# @file testcases.sh
# @brief Integration tests triggered from Docker Compose.
#
# @description This script contains intergration test cases that are run against the Jira
# instance defined in the docker-compose.yml from this test setup. The ``local/jiracli:dev``
# Docker image is started in a container and executes the testcases from this script against
# a running Jira instance.
#
# Before running the tests, the script checks if Jira is up-and-running and triggers a re-index
# to make sure Jira exposes all data as intended.
#
# IMPORTANT: This script is intended to run inside a ``jiracli`` container from Docker Compose.
# Don't run this script directly!
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

readonly FLAG_BASE_URL="--baseUrl=$BASE_URL"
readonly FLAG_USER="--user=$USER"
readonly FLAG_PASS="--pass=$PASS"


# @description Wrapper function to sleep for a given interval.
#
# @example
#    waitFor 30s
#    waitFor 1m
#
# @arg $@ int The interval to sleep - Mandatory
#
# @exitcode 8 If interval param is missing
function waitFor() {
  if [ -z "$1" ]; then
    LOG_ERROR "No interval passed"
    LOG_ERROR "exit" && exit 8
  fi

  echo "sleep $1"
  sleep "$1"
}


# @description This function checks whether Jira is up-and-running and responding to requests.
#
# @example
#    isStartupComplete
function isStartupComplete() {
  waitFor "1m"

  while true; do
    httpCode=$(curl --location --request GET "$BASE_URL/rest/api/2/serverInfo?doHealthCheck=true" \
      -s -o /dev/null -w "%{http_code}" \
      --user "$USER:$PASS" \
      --header 'Content-Type: application/json')
    
    echo "Jira answered with $httpCode"
    if [[ "$httpCode" == "200" ]]; then
      echo "Jira is up-and-running"
      break
    fi
    
    waitFor "10s"
  done
}


# @description When starting the Jira instance it is provided with a h2db and is populated with a
# license, users and projects so there is no need re-run the setup wizard all the time. This data
# defines a baseline which is the same everytime this test-stack starts.
#
# The only issue is, that there is no data visible in the UI after starting up because  there is
# no (or at least no uncorrupted) index. A new index must be created in the foreground triggered
# through the link:https://docs.atlassian.com/software/jira/docs/api/REST/8.22.6[Jira REST API].
#
# @example
#    buildIndex
function buildIndex() {
  echo "Trigger re-index"
  response=$(curl -s --location --request POST "$BASE_URL/rest/api/2/reindex?type=FOREGROUND" \
    --user "$USER:$PASS" \
    --header 'Accept: application/json')
  echo "$response"
  
  sleep "10s"

  echo "Check re-index"
  response=$(curl -s --location --request GET "$BASE_URL/rest/api/2/reindex" \
    --user "$USER:$PASS" \
    --header 'Content-Type: application/json')
  echo "$response"
}


# @description Wrapper function to not repeat the mandatory flags when running ``jiracli``
# commands.
#
# @example
#    JIRACLI --version
#
# @arg $@ String The ``jiracli`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``jiracli`` command is missing
function JIRACLI() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the jiracli app"
    LOG_ERROR "exit" && exit 8
  fi

  jiracli "$@"
}


isStartupComplete
buildIndex


JIRACLI --version

JIRACLI user view admin.admin "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view admin.admin --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view admin@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

JIRACLI user view jim.panse "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view jim.panse --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view jim.panse@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

JIRACLI user view homer.simpson "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view homer.simpson --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
JIRACLI user view homer.simpson@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

JIRACLI license view "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
