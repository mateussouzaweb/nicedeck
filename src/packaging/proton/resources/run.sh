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

# Replace C: with driver path on first argument
if [[ "$1" =~ ^[Cc]: ]]; then
  set -- "${1/#[Cc]:/$DRIVE_PATH}" "${@:2}"
fi

# Expand glob safely on first argument if contains an asterisk
if [[ "$1" == *"*"* ]] && [[ ! -e "$1" ]]; then
  MATCHED=$(compgen -G "$1" | head -n 1)
  if [[ -e "${MATCHED}" ]]; then
    set -- "${MATCHED}" "${@:2}"
  fi
fi

# Go to target working directory based on executable path
# NOTE: This step is required for some games / applications
if [[ -n "$1" && "$1" == /* ]] && [[ -e "$1" ]]; then
  cd "$(dirname "$1")"
fi

# Replace Proton Experimental with another version
# Example: PROTON="10" %command%
# Example: PROTON="Proton-GE" %command%
# Example: PROTON="Proton-CachyOS Latest" %command%
if [[ -n "$PROTON" ]]; then
  if [[ -d "$STEAM_PATH/steamapps/common/Proton $PROTON" ]]; then
    PROTON_RUNTIME="$STEAM_PATH/steamapps/common/Proton $PROTON/proton"
    WINE_BINARY="$STEAM_PATH/steamapps/common/Proton $PROTON/files/bin/wine"
  elif [[ -d "$STEAM_PATH/steamapps/common/Proton $PROTON.0" ]]; then
    PROTON_RUNTIME="$STEAM_PATH/steamapps/common/Proton $PROTON.0/proton"
    WINE_BINARY="$STEAM_PATH/steamapps/common/Proton $PROTON.0/files/bin/wine"
  elif [[ -d "$STEAM_PATH/steamapps/common/Proton - $PROTON" ]]; then
    PROTON_RUNTIME="$STEAM_PATH/steamapps/common/Proton - $PROTON/proton"
    WINE_BINARY="$STEAM_PATH/steamapps/common/Proton - $PROTON/files/bin/wine"
  elif [[ -d "$STEAM_PATH/compatibilitytools.d/$PROTON" ]]; then
    PROTON_RUNTIME="$STEAM_PATH/compatibilitytools.d/$PROTON/proton"
    WINE_BINARY="$STEAM_PATH/compatibilitytools.d/$PROTON/files/bin/wine"
  fi
fi

# Activate GameMode or use default process
# Example: GAMEMODE=1 %command%
# Debug: gamemoded -s && gamemodelist
if [[ "$GAMEMODE" == "1" ]]; then
  COMMAND=("gamemoderun")
elif [[ "$INSTALL_TYPE" == "flatpak" ]]; then
  COMMAND=("exec")
else
  COMMAND=()
fi

# Required changes for Flatpak compatibility
# Use wrapper command and paths inside sandbox environment
if [[ "$INSTALL_TYPE" == "flatpak" ]]; then
  SEARCH="/.var/app/$FLATPAK_ID/"
  STEAM_PATH="${STEAM_PATH/$SEARCH//}"
  STEAM_RUNTIME="${STEAM_RUNTIME/$SEARCH//}"
  PROTON_RUNTIME="${PROTON_RUNTIME/$SEARCH//}"
  WINE_BINARY="${WINE_BINARY/$SEARCH//}"

  COMMAND=(
    /usr/bin/flatpak run
    --branch="stable" --file-forwarding
    --cwd="$PWD" --command="${COMMAND[0]}"
    "$FLATPAK_ID"
  )
fi

# Determine environment variables and command arguments
if [[ "$1" == "wine" ]]; then
  export WINEPREFIX="$WINE_PATH"
  ARGUMENTS=("$WINE_BINARY")
  shift
else
  export STEAM_COMPAT_CLIENT_INSTALL_PATH="$STEAM_PATH"
  export STEAM_COMPAT_DATA_PATH="$DATA_PATH"
  ARGUMENTS=(
    "$STEAM_RUNTIME" "--verb=run" "--"
    "$PROTON_RUNTIME" "run"
  )
fi

# Extra scripting support
if [[ -f "$DATA_PATH/run.extra.sh" ]]; then
  source "$DATA_PATH/run.extra.sh"
fi

# Execute the command
exec "${COMMAND[@]}" "${ARGUMENTS[@]}" "$@" 2>&1