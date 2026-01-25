#!/bin/bash

# Variables for execution
INSTALL_TYPE="@{INSTALL_TYPE}"
FLATPAK_ID="@{FLATPAK_ID}"
DATA_PATH=$(realpath "@{DATA_PATH}")
WINE_PATH=$(realpath "@{WINE_PATH}")
DRIVE_PATH=$(realpath "@{DRIVE_PATH}")
STEAM_PATH=$(realpath "@{STEAM_PATH}")
STEAM_RUNTIME=$(realpath "@{STEAM_RUNTIME}")
PROTON_RUNTIME=$(realpath "@{PROTON_RUNTIME}")
WINE_BINARY=$(realpath "@{WINE_BINARY}")

# Determine environment variables, command and its arguments
if [[ "$1" == "wine" ]]; then
  export WINEPREFIX="$WINE_PATH"
  COMMAND=("$WINE_BINARY")
  ARGUMENTS=()
  shift
else
  export STEAM_COMPAT_CLIENT_INSTALL_PATH="$STEAM_PATH"
  export STEAM_COMPAT_DATA_PATH="$DATA_PATH"
  COMMAND=("$STEAM_RUNTIME")
  ARGUMENTS=("$PROTON_RUNTIME" "run")
fi

# Replace C: with driver path
if [[ "$1" =~ ^[Cc]: ]]; then
  set -- "${1/C:/$DRIVE_PATH}" "${@:2}"
  set -- "${1/c:/$DRIVE_PATH}" "${@:2}"
fi

# Go to target working directory based on executable path if defined
# This step is required for some games / applications
if [[ -n "$1" && "$1" == /* && -e "$1" ]]; then
  cd "$(dirname "$1")" || exit 1
fi

# Wrapper for Flatpak compatibility
if [[ "$INSTALL_TYPE" == "flatpak" ]]; then
  COMMAND=(
    /usr/bin/flatpak run
    --branch="stable"
    --file-forwarding
    --cwd="$PWD"
    --command="${COMMAND[0]}"
    "$FLATPAK_ID"
  )
fi

# Execute the command
exec "${COMMAND[@]}" "${ARGUMENTS[@]}" "$@" 2>&1