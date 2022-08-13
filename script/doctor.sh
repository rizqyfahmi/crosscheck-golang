#!/usr/bin/env bash

ERR_COUNT=0

main() {
  # Check Programming Runtimes
  check_cmd "go"
  check_cmd "migrate" "golang-migrate"
  check_cmd "ginkgo"
  check_cmd "make"
  check_cmd "git"
  check_optional_cmd "brew"
  check_optional_cmd "docker"

  echo

  if [[ ${ERR_COUNT} -gt 0 ]]; then
    echo "Result: error count: ${ERR_COUNT}"
    exit 1
  fi

  echo "Result: all requirement is ready"
}

check_cmd() {
  # Get arguments
  CMD_NAME=$1

  if ! [[ -z "$2" ]]; then
    CMD_NAME=$2
  fi

  echo -e "Checking command ${CMD_NAME}: \c"

  if ! [[ -x "$(command -v "$1")" ]]; then
    ERR_COUNT=$((ERR_COUNT + 1))
    echo "NOT FOUND"

  else
    echo "OK"
  fi
}

check_optional_cmd() {
  # Get arguments
  CMD_NAME=$1

  echo -e "Checking optional command ${CMD_NAME}: \c"

  if ! [[ -x "$(command -v "${CMD_NAME}")" ]]; then
    echo "NOT FOUND"

  else
    echo "OK"
  fi
}

check_script() {
  # Get arguments
  SCRIPT_FILE=$1

  echo -e "Checking script ${SCRIPT_FILE}: \c"

  if ! [[ -f ${SCRIPT_FILE} && -x ${SCRIPT_FILE} ]]; then
    ERR_COUNT=$((ERR_COUNT + 1))
    echo "NON EXECUTABLE"
  else
    echo "OK"
  fi
}

main
