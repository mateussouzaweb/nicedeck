#!/bin/bash

# Variables for execution
export INSTALL_TYPE="${INSTALL_TYPE}"
export DRIVE_PATH=$(realpath "${DRIVE_PATH}")
export PROTON_RUNTIME=$(realpath "${PROTON_RUNTIME}")
export STEAM_COMPAT_CLIENT_INSTALL_PATH=$(realpath "${STEAM_CLIENT_PATH}")
export STEAM_COMPAT_DATA_PATH=$(realpath "${DATA_PATH}")
export STEAM_RUNTIME=$(realpath "${STEAM_RUNTIME}")

# Replace C: with driver path
set -- "${1/C:/$DRIVE_PATH}" "${@:2}"

# Go to target executable path
# This step is required for some games
export WORKING_DIRECTORY=$(dirname "$1")
cd "$WORKING_DIRECTORY"

# Run command based on install type
# In all cases, this path is relative to the environment

# Steam Flatpak
if [[ "$INSTALL_TYPE" -eq "flatpak" ]]; then
  exec /usr/bin/flatpak run \
    --branch=stable --file-forwarding \
    --cwd="$WORKING_DIRECTORY" --command="$STEAM_RUNTIME" \
    com.valvesoftware.Steam "$PROTON_RUNTIME" run "$@" 2>&1

# Steam Native
elif [[ "$INSTALL_TYPE" -eq "native" ]]; then
  exec "$STEAM_RUNTIME" "$PROTON_RUNTIME" run "$@" 2>&1
fi
