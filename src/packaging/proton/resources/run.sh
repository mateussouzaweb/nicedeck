#!/bin/bash

# Variables for execution
INSTALL_TYPE="@{INSTALL_TYPE}"
FLATPAK_ID="@{FLATPAK_ID}"
DATA_PATH=$(realpath "@{DATA_PATH}")
DRIVE_PATH=$(realpath "@{DRIVE_PATH}")
STEAM_PATH=$(realpath "@{STEAM_PATH}")
STEAM_RUNTIME=$(realpath "@{STEAM_RUNTIME}")
PROTON_RUNTIME=$(realpath "@{PROTON_RUNTIME}")

# Replace C: with driver path
set -- "${1/C:/$DRIVE_PATH}" "${@:2}"

# Go to target executable path
# This step is required for some games
WORKING_DIRECTORY=$(dirname "$1")
cd "$WORKING_DIRECTORY"

# Steam Flatpak
if [[ "$INSTALL_TYPE" -eq "flatpak" ]]; then

  # When running with flatpak, need to use sandboxed paths 
  SEARCH="/.var/app/$FLATPAK_ID/"
  STEAM_PATH="${STEAM_PATH/$SEARCH//}"
  STEAM_RUNTIME="${STEAM_RUNTIME/$SEARCH//}"
  PROTON_RUNTIME="${PROTON_RUNTIME/$SEARCH//}"

  export STEAM_COMPAT_CLIENT_INSTALL_PATH="$STEAM_PATH"
  export STEAM_COMPAT_DATA_PATH="$DATA_PATH"
  exec /usr/bin/flatpak run \
    --branch=stable --file-forwarding \
    --cwd="$WORKING_DIRECTORY" --command="$STEAM_RUNTIME" \
    "$FLATPAK_ID" "$PROTON_RUNTIME" run "$@" 2>&1

# Steam Native
elif [[ "$INSTALL_TYPE" -eq "native" ]]; then

  export STEAM_COMPAT_CLIENT_INSTALL_PATH="$STEAM_PATH"
  export STEAM_COMPAT_DATA_PATH="$DATA_PATH"
  exec "$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$@" 2>&1

# Unknown
else
  echo "ERROR: Unknown installation type: $INSTALL_TYPE"
  exit 1
fi
