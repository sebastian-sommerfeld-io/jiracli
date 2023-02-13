#!/bin/bash
# @file integration-tests.sh
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


which jiracli

jiracli user view admin.admin
jiracli user view admin.admin --username
jiracli user view admin@localhost --email

jiracli user view jim.panse
jiracli user view jim.panse --username
jiracli user view jim.panse@localhost --email

jiracli user view homer.simpson
jiracli user view homer.simpson --username
jiracli user view homer.simpson@localhost --email
