# Know Issues

## ARM64 Specific Issues

These issues are generally present in Linux, MacOs and Windows:

- First, is important to notice that even if there is no oficial build for ARM64, NiceDeck will download the AMD64 version of the application that may run over compatibility layer.
- Unfortunately, many applications still does not provide oficial builds for the ARM64 architecture. That means that you may experience issues when trying to run such applications.
- Please remember that those issues are not related to NiceDeck, which acts only as a facilitator to install and run such programs.

## Linux Specific Issues

- Windows native programs provided with Proton implementation from NiceDeck always use "Proton - Experimental" provided by Steam. You should manually install this Proton version from Steam in your device before trying to use the NiceDeck Proton feature.
- You don't need or should configure the Proton version inside Steam shortcuts as this is automatically handled by NiceDeck internally. Making such action will result in error.
- If the game don't work with NiceDeck Proton implementation, remove the NiceDeck Proton shortcut of the game and try to use it natively with Steam, adding it to the library and manually selecting the Proton version.

## MacOS Specific Issues

- You may need to give "Full Disk Access" system permissions to the "Terminal" app before trying to use NiceDeck.
- You may need to install Rosetta 2 before trying to install or running non-native applications on Apple Silicon.
- Battle.net package is provided from homebrew, but it does not seems to work.
- GOG Galaxy package is provided from homebrew and requires sudo to be installed. If you try to install it from NiceDeck, the process will hang indefinitely.
- PCSX2 and ShadPS4 downloaded archive files has another archive inside it, so it cannot be properly extracted. Trying to install it will return an error.
- In general, MacOS gaming still is not a good experience. Do not expect that everything will work flawlessly.

## Windows Specific Issues

- When removing Discord, it does not fully remove itself and you will still see the app on Windows.
- Brave Browser removes automatically when requested, but return error on NiceDeck.
- Ubisoft Connect will remove automatically when requested, but the Windows start menu shortcut is not removed.
