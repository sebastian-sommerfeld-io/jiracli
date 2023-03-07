#!/bin/bash
# @file tests.sh
# @brief Start the docker stack and run tests or stop the docker stack.
#
# @description The script starts and stops the docker compose stack used for the local jira
# instance. This stack is used for automated tests against a running Jira system. This stack
# also builds the ``local/jiracli:test`` image to run tests based on ``src/main/Dockerfile``.
#
# These tests are also triggered from the CI pipeline. The ``local/jiracli:test`` image which
# is built from this Docker Compose configuration is also deployed to DockerHub as
# link:https://hub.docker.com/r/sommerfeldio/jiracli[``sommerfeldio/jiracli:ci``]. This
# ``sommerfeldio/jiracli:ci`` image acts as the foundation for the CD pipeline and is re-tagged
# as ``sommerfeldio/jiracli:latest`` and ``sommerfeldio/jiracli:<version>``. This way it is
# ensured, that all releases are based on the image which was tested previously and that only
# one image build takes place.
#
# === Script Arguments
#
# * *$1* (string): Command (``start``, ``logs``, ``stop``)
#
# === Script Example
#
# [source, bash]
# ```
# ./tests.sh start
# ./tests.sh logs
# ./tests.sh stop
# ```


readonly START="start"
readonly LOGS="logs"
readonly STOP="stop"

ARG="$1"
readonly ARG

# Include local bash modules
source "../../util/bash-modules/log.sh"


if [ -z "$ARG" ]; then
  LOG_ERROR "Param missing: command (start | logs | stop)"
  LOG_ERROR "exit" && exit 0
fi


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


case "$ARG" in
  "$START" )
    LOG_HEADER "Start stack"
    docker-compose up --build -d
  ;;
  "$LOGS" )
    LOG_HEADER "Show logs"
    docker-compose logs
  ;;
  "$STOP" )
    LOG_HEADER "Stop stack"
    docker-compose down -v --rmi all --remove-orphans
  ;;
esac
