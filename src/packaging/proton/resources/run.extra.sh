#!/bin/bash

# Extra scripting support to Proton run.sh wrapper.
# You can use this file to:
# - Add custom environment variables with export keyword
# - Modify the "$COMMAND" and "$ARGUMENTS" arrays
# - Set wrapper arguments by modifying the "$@" array
# - Others

# Optional global variable definitions 
# PROTON="10"
# GAMEMODE="1"

# Replace Proton Experimental with another version
# Example: PROTON="10" %command%
# Example: PROTON="Proton-GE" %command%
# Example: PROTON="Proton-CachyOS Latest" %command%
if [[ -n "$PROTON" ]]; then

  if [[ -d "$STEAM_PATH/steamapps/common/Proton $PROTON" ]]; then
    PROTON_REPLACE="steamapps/common/Proton $PROTON"
  elif [[ -d "$STEAM_PATH/steamapps/common/Proton $PROTON.0" ]]; then
    PROTON_REPLACE="steamapps/common/Proton $PROTON.0"
  elif [[ -d "$STEAM_PATH/steamapps/common/Proton - $PROTON" ]]; then
    PROTON_REPLACE="steamapps/common/Proton - $PROTON"
  elif [[ -d "$STEAM_PATH/compatibilitytools.d/$PROTON" ]]; then
    PROTON_REPLACE="compatibilitytools.d/$PROTON"
  fi

  if [[ -n "$PROTON_REPLACE" ]]; then
    PROTON_SEARCH="steamapps/common/Proton - Experimental"
    COMMAND=("${COMMAND[@]/"$PROTON_SEARCH"/$PROTON_REPLACE}")
    ARGUMENTS=("${ARGUMENTS[@]/"$PROTON_SEARCH"/$PROTON_REPLACE}")
  fi

fi

# Activate GameMode
# Example: GAMEMODE=1 %command%
# Debug: gamemoded -s && gamemodelist
if [[ "$GAMEMODE" == "1" ]] && [[ "$1" != "wine" ]]; then
  ARGUMENTS=("gamemoderun" "${ARGUMENTS[@]}")
fi