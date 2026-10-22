## Game Launchers

In addition to [ROMs](ROMs.md), NiceDeck can detect installed games from third-party game launchers and automatically add them to your Steam Library. This feature offers a good balance for those who have games from other stores outside Steam.

On Linux, NiceDeck uses the Proton layer to run additional game launchers, making it feel like you are on Windows. See the [Proton](./Proton.md) section for more details. The MacOS experience is still experimental and requires more testing.

## Supported Launchers

The table below shows which launchers are supported on each operating system:

| Launchers     | Windows | Linux | macOS        |
|---------------|---------|-------|--------------|
| Amazon Games  | Yes     | Yes   | Untested     |
| Battle.net    | Yes     | Yes   | Untested     |
| EA Games      | Yes     | Yes   | Untested     |
| Epic Games    | Yes     | Yes   | Untested     |
| GOG           | Yes     | Yes   | Untested     |
| Ubisoft Games | Yes     | Yes   | Untested     |

## Tips

For the best possible experience, here are a few tips:

- Game detection does not require authentication inside NiceDeck. Just open the launcher, install your games, then run the scanner to import them.
- The scanner is faster but can only detect installed games; seeing your full library would require a complex authentication mechanism.
- Always keep the default installation path for games. This ensures the scanner can detect the installed games from the launcher.
- On Linux with Proton, you are encouraged to disable the *minimize to system tray* option inside launchers. The system tray feature keeps the launcher running in the background, which can cause issues.