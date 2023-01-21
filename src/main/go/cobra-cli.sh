#!/bin/bash
# @file cobra-cli.sh
# @brief Use cobra-cli to generate some code blocks.
#
# @description The script runs the ``cobra-cli`` in Docker to generate some source code blocks
# like adding new cli-commands.
#
# Cobra provides its own program (``link:https://github.com/spf13/cobra-cli[cobra-cli]``) that
# will create your application and add any commands you want. It's the easiest way to incorporate
# Cobra into your application.
#
# === Script Arguments
#
# The script does not accept any parameters.
#
# === Script Example
#
# [source, bash]
# ```
# ./cobra-cli.sh
# ```
#
# === See Also
#
# * https://github.com/spf13/cobra-cli


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

# Include local bash modules
source "../utils/bash-modules/log.sh"
source "../utils/bash-modules/go-wrapper.sh"


readonly OPTION_INIT="init"
readonly OPTION_ADD_COMMAND="add_command"


# @description Wrapper function to encapsulate ``cobra-cli`` in a docker container.
# The current working directory is mounted into the container and selected as working
# directory so that all file are available to ``cobra-cli``. Paths are preserved.
# The container runs with the current user.
#
# @example
#    cobra-cli --help
#
# @arg $@ String The ``cobra-cli`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``cobra-cli`` command is missing
function cobra-cli() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the cobra-cli container"
    LOG_ERROR "exit" && exit 8
  fi

  docker run --rm \
    --volume /etc/passwd:/etc/passwd:ro \
    --volume /etc/group:/etc/group:ro \
    --user "$(id -u):$(id -g)" \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    local/cobra-cli:dev cobra-cli "$@"
}


# @description The ``cobra-cli init [app]`` command will create your initial application code for
# you. It is a very powerful application that will populate your program with the right structure
# so you can immediately enjoy all the benefits of Cobra. It can also apply the license you specify
# to your application.
function init() {
  LOG_HEADER "Initialize application"

  LOG_INFO "Enter application name ..."
  read -r name
  readonly name

  LOG_INFO "Initialize $name"
  # cobra-cli init "$name"
  cobra-cli --help
}


# @description Once a cobra application is initialized you can continue to use the Cobra generator
# to add additional commands to your application. The command to do this is ``cobra-cli add``.
function add_command() {
  LOG_HEADER "Add command"

  LOG_INFO "Enter command name to add to project - e.g."
  LOG_INFO "    - cobra-cli add [NAME]"
  LOG_INFO "    - cobra-cli add [NAME] -p '[PARENTNAME]Cmd'"
  read -r name
  readonly name

  LOG_INFO "Initialize $name"
  # cobra-cli add "$name"
  cobra-cli --help
}


LOG_HEADER "Build utility docker image for cobra-cli"
(
  cd "../utils/cobra-cli" || exit
  bash build.sh
)

LOG_HEADER "What do you want me to do?"
select task in "$OPTION_INIT" "$OPTION_ADD_COMMAND"; do
  case "$task" in
    "$OPTION_INIT" ) init ;;
    "$OPTION_ADD_COMMAND" ) add_command ;;
  esac

  break
done
