#!/bin/bash
# @file .go-wrapper.sh.sh
# @brief Bash module which provides the wrapper function to run Go inside a Docker container.
#
# @description The script is a wrapper for the `jiracli` command. The script builds and runs
# the ``jiracli`` locally. Before building, all tests are run as well. Everything go-related
# is delegated to the link:https://hub.docker.com/_/golang[golang] Docker image.
#
# === Script Arguments
#
# * *$@* (...): The same parameters as ``jiracli`` app (see ``./jiracli.sh --help``)
#
# === Script Example
#
# [source, bash]
# ```
# ./jiracli.sh
# ```


readonly LOG_ERROR="[\e[1;31mERROR\e[0m]" # Avoid 'unbound variable' errors in pipeline

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
# @arg $@ String The go commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with terraform command is missing
function go() {
  if [ -z "$1" ]; then
    echo -e "$LOG_ERROR No command passed to the go container"
    echo -e "$LOG_ERROR exit" && exit 8
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
