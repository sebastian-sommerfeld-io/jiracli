#!/bin/bash
# @file log.sh
# @brief Bash module which provides utility functions for logging.
#
# @description The script is bash module which provides utility functions for logging.
#
# CAUTION: This script is a module an is not intended to run on its own. Include in script and
# use its functions!.
#
# === Script Arguments
#
# The script does not accept any parameters.


# TODO function LOG_ERROR and LOG_INFO
# TODO    then remove vars from scripts
readonly LOG_INFO="[\e[34mINFO\e[0m]" # Avoid 'unbound variable' errors in pipeline

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


# @description Print log output in a header-style.
#
# @arg $@ String The line to print.
function LOG_HEADER() {
  echo -e "$LOG_INFO ------------------------------------------------------------------------"
  echo -e "$LOG_INFO $1"
  echo -e "$LOG_INFO ------------------------------------------------------------------------"
}
