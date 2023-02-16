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


# @description Wrapper function to not repeat the three mandatory flags for the ``jiracli``
# app all the time.
#
# @example
#    _jiracli user view admin.admin
#
# @arg $@ String The ``jiracli`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``jiracli`` command is missing
function _jiracli() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the jiracli app"
    LOG_ERROR "exit" && exit 8
  fi

  jiracli "$@" --baseUrl="http://localhost:8080" --user="admin" --pass="admin"
}


_jiracli user view admin.admin 
_jiracli user view admin.admin --username
_jiracli user view admin@localhost --email

_jiracli user view jim.panse
_jiracli user view jim.panse --username
_jiracli user view jim.panse@localhost --email

_jiracli user view homer.simpson
_jiracli user view homer.simpson --username
_jiracli user view homer.simpson@localhost --email

_jiracli license view
