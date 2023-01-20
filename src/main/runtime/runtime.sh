#!/bin/bash
# @file runtime.sh
# @brief Start or stop the docker stack.
#
# @description The script starts and stops the docker stack used for the local jira instance.
#
# === Script Arguments
#
# * *$1* (string): Command (``start``, ``logs``, ``stop``)
#
# === Script Example
#
# [source, bash]
# ```
# ./start.sh
# ```


readonly START="start"
readonly LOGS="logs"
readonly STOP="stop"

ARG="$1"
readonly ARG


echo -e "$LOG_INFO Run $0"


if [ -z "$ARG" ]; then
  echo -e "$LOG_ERROR Param missing: command (start | logs | stop)"
  echo -e "$LOG_ERROR exit" && exit 0
fi


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


case "$ARG" in
    "$START" )
        echo -e "$LOG_INFO Start stack"
        docker-compose up -d
    ;;
    "$LOGS" )
        echo -e "$LOG_INFO Show logs"
        docker-compose logs -f
    ;;
    "$STOP" )
        echo -e "$LOG_INFO Stop stack"
        docker-compose down -v --rmi all
    ;;
esac
