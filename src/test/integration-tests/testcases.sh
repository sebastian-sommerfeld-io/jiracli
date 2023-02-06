#!/bin/bash
# @file integration-tests.sh
# @brief Integration tests triggered from Docker Compose.
#
# @description Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
#
# IMPORTANT: This script is intended to run inside a ``jiracli`` container from Docker Compose.
# Don't run this script directly!
#
# === Script Arguments
#
# The script does not accept any parameters.

# todo ... use in container instead of single command ... this way I can do some assertions here
set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


which jiracli

jiracli user view admin.admin
jiracli user view jim.panse
