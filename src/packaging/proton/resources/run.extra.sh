#!/bin/bash

# Extra scripting support to Proton run.sh wrapper. 
# You can use this file to:
# - Add custom environment variables with export keyword
# - Modify the "$COMMAND" and "$ARGUMENTS" arrays 
# - Set wrapper arguments by modifying the "$@" array
# - Others

############################################
# Activate GameMode
# Debug: gamemoded -s && gamemodelist
# if [[ "$1" != "wine" ]]; then
#   ARGUMENTS=("gamemoderun" "${ARGUMENTS[@]}")
# fi

############################################
# Replace Proton Experimental with Proton-CachyOS Latest
# PROTON_SEARCH="steamapps/common/Proton - Experimental"
# PROTON_REPLACE="compatibilitytools.d/Proton-CachyOS Latest"
# COMMAND=("${COMMAND[@]/$PROTON_SEARCH/$PROTON_REPLACE}")
# ARGUMENTS=("${ARGUMENTS[@]/$PROTON_SEARCH/$PROTON_REPLACE}")

############################################
# Fix PATH and STEAM_COMPAT_MOUNTS for Proton compatibility
# Required to make xdg-open works with KDE system
# Please note that this is not required when running Steam flatpak
# if [[ "$INSTALL_TYPE" != "flatpak" ]]; then
#  
#   export PATH="$DATA_PATH/bin:$PATH"
#   export STEAM_COMPAT_MOUNTS="$DATA_PATH/bin"
#
#   mkdir -p "$DATA_PATH/bin"
#   touch "$DATA_PATH/bin/xdg-open"
#   chmod +x "$DATA_PATH/bin/xdg-open"
#   cat > "$DATA_PATH/bin/xdg-open" << EOF
#   #!/bin/bash
#   case "\$1" in
#     http://*|https://*)
#       exec gdbus call --session \
#         --dest=org.freedesktop.portal.Desktop \
#         --object-path /org/freedesktop/portal/desktop \
#         --method org.freedesktop.portal.OpenURI.OpenURI \
#         "" "\$1" "{}"
#       ;;
#     *)
#       exec /usr/bin/xdg-open "\$@"
#       ;;
#   esac
# EOF
#
# fi