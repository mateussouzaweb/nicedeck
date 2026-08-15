# Extending NiceDeck

NiceDeck offers ways to extend its capabilities with custom configuration files for your specific use case. You can add support for extra console platforms, emulators, and states to parse ROMs and sync additional save data using the built-in features of NiceDeck.

| Feature          | Status            | Notes |
|------------------|-------------------|-------|
| Custom Platform  | **supported**     | Base for emulators and states |
| Custom Emulators | **supported**     | Used to process ROMs |
| Custom State     | **supported**     | Used to sync state |
| Custom Programs  | **not supported** | You can manually add programs as shortcuts from the UI |

## Custom Platforms

Platforms are the foundation for emulators and states. You should add new platforms when the console you want is not yet available in the official list.

File: `$HOME/Games/Applications/NiceDeck/custom/platforms.json`

```json
[{
    "name": "SNES",
    "console": "Super Nintendo",
    "folder": "SNES"
}]
```

## Custom Emulators

You can process ROMs with additional emulators for built-in platforms or custom-added platforms.
This feature is useful for adding support for emulators you have installed manually on your device.

File: `$HOME/Games/Applications/NiceDeck/custom/emulators.json`

```json
[{
    "name": "SuperZNES",
    "platform": "SNES",
    "program": "superznes",
    "executable": "$EMULATORS/SuperZNES/SUPERZNES",
    "extensions": ".bin .bml .bs .bsx .dx2 .fig .gd3 .gd7 .mgd .sfc .smc .st .swc .7z .zip",
    "launchOptions": "${ROM}"
}]
```

Note that the `executable` value supports environment variables (including NiceDeck built-in variables) to expand the path.
Additionally, `${ROM}` represents the full path of the parsed ROM file.

## Custom States

Declare additional rules to sync state. This can also be used to sync program configurations or saves from any game that you have.

File: `$HOME/Games/Applications/NiceDeck/custom/states.json`

```json
[{
    "platform": "SNES",
    "emulator": "SuperZNES",
    "type": "folder",
    "source": "$EMULATORS/SuperZNES/data",
    "destination": "$STATE/SuperZNES/data"
}]
```

Please note that the `source` and `destination` values support environment variables to expand the path, including NiceDeck built-in variables.
