#!/bin/bash
# @file go-wrapper.sh.sh
# @brief Bash module which provides the wrapper function to run Go inside a Docker container.
#
# @description The script is Bash module which provides the wrapper function to run Go inside
# a Docker container. All go commands run in the link:https://hub.docker.com/_/golang[golang]
# Docker image.
#
# CAUTION: This script is a module an is not intended to run on its own. Include in script and
# use its functions!.
#
# === Script Arguments
#
# The script does not accept any parameters.


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


# @description Wrapper function to encapsulate ``go`` in a docker container. The current working
# directory is mounted into the container and selected as working directory so that all file are
# available to ``go``. Paths are preserved. The container runs with the current user.
#
# @example
#    go version
#
# @arg $@ String The ``go`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``go`` command is missing
function go() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the go container"
    LOG_ERROR "exit" && exit 8
  fi

  mkdir -p "/tmp/$USER/.cache"

  docker run --rm \
    --volume /etc/passwd:/etc/passwd:ro \
    --volume /etc/group:/etc/group:ro \
    --user "$(id -u):$(id -g)" \
    --volume "/tmp/$USER/.cache:/home/$USER/.cache" \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    golang:1.20-rc-alpine go "$@"
}
