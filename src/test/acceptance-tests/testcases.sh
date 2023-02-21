#!/bin/bash
# @file testcases.sh
# @brief Integration tests triggered from Docker Compose.
#
# @description This script contains intergration test cases that are run against the Jira
# instance defined in the docker-compose.yml from this test setup. The ``local/jiracli:dev``
# Docker image is started in a container and executes the testcases from this script against
# a running Jira instance.
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


readonly FLAG_BASE_URL="--baseUrl=http://jira:8080"
readonly FLAG_USER="--user=admin"
readonly FLAG_PASS="--pass=admin"


# @description Wrapper function to not repeat the three mandatory flags for the ``jiracli``
# app all the time.
#
# @example
#    _jiracli --version
#
# @arg $@ String The ``jiracli`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``jiracli`` command is missing
function _jiracli() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the jiracli app"
    LOG_ERROR "exit" && exit 8
  fi

  jiracli "$@"
}

_jiracli --version

_jiracli user view admin.admin "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view admin.admin --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view admin@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

_jiracli user view jim.panse "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view jim.panse --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view jim.panse@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

_jiracli user view homer.simpson "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view homer.simpson --by-username "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
_jiracli user view homer.simpson@localhost --by-email "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"

_jiracli license view "$FLAG_BASE_URL" "$FLAG_USER" "$FLAG_PASS"
