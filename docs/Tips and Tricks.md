# Tips and Tricks

Here are a few tips and tricks to expand the features of NiceDeck in your device:

## KDE Context Service Menu

If you are using Linux, you can use the script below to create a new KDE context service menu to directly add a Windows native application to the library or run it with Proton:

```bash
DESTINATION="$HOME/.local/share/kio/servicemenus/nicedeck.proton.desktop"
ICON="$HOME/.local/share/icons/nicedeck.png"
NICEDECK_BIN="$HOME/Games/Applications/NiceDeck/nicedeck"
PROTON_RUN="$HOME/Games/Proton/run.sh"

mkdir -p $(dirname $DESTINATION)
touch $DESTINATION
chmod +x $DESTINATION
cat > $DESTINATION << EOF
[Desktop Entry]
Type=Service
MimeType=application/vnd.microsoft.portable-executable;application/x-msdownload;application/x-msi;application/x-msdos-program;application/x-dosexec;application/x-bat;application/bat;application/octet-stream;
Actions=addToLibrary;runWithProton;
Terminal=true
X-KDE-Priority=TopLevel

[Desktop Action addToLibrary]
Name=Add to Library
Icon=$ICON
Exec=sh -c '$NICEDECK_BIN create --path="%u"'

[Desktop Action runWithProton]
Name=Run with Proton
Icon=$ICON
Exec=$PROTON_RUN "%u"
EOF
```

If you need to remove the service menu in the future:

```bash
DESTINATION="$HOME/.local/share/kio/servicemenus/nicedeck.proton.desktop"
rm $DESTINATION
```

## Running Games with Proton

- Some games will crash if ``NVAPI`` is enabled. You can tweak the launch options in the shortcut to disable ``NVAPI`` or made other tweaks.
- Disabling ``NVAPI`` is required in GRID racing series for example (``Race Drive GRID``, ``GRID 2``, ``GRID AutoSports``, ...)
- To disable NVAPI, edit the shortcut from NiceDeck and set launch options like: ``PROTON_DISABLE_NVAPI=1 %command% ...``.
- Additional options for more tweaks are provided on the [Proton repository](https://github.com/ValveSoftware/Proton?tab=readme-ov-file#runtime-config-options).
