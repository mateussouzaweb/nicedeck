# Known Issues

## Linux Specific Issues

Luckily, there is not much here, but you should read the [Proton](./Proton.md) section for more insights about that specific implementation.

## Windows Specific Issues

Installing and removing programs:

- When removing Discord, it does not fully remove itself and you will still see the app on Windows.
- Brave Browser removes automatically when requested, but returns an error in NiceDeck.

Windows file and folder shortcuts:

- Windows shortcut ``.lnk`` files or junction points are not accepted or handled by NiceDeck.
- If you want to put additional content in another location with a symbolic link, you must create a valid directory shortcut from the Terminal.
- For example, to create a directory shortcut for **ROMs** on another disk (*D:*), open the *Terminal* application with **administrator privileges**, go to the ``$GAMES`` folder and run the following command to create the symbolic link: ``cmd /c mklink /d ".\ROMs" "D:\ROMs"``.

## MacOS Specific Issues

In general, MacOS gaming is still not a good experience. Do not expect that everything will work flawlessly:

- You may need to give "Full Disk Access" system permissions to the "Terminal" app before trying to use NiceDeck.
- You may need to install Rosetta 2 before trying to install or run non-native applications on Apple Silicon.
- The Battle.net package is provided from Homebrew, but it does not seem to work.
- The GOG Galaxy package is provided from Homebrew and requires sudo to be installed. If you try to install it from NiceDeck, the process will hang indefinitely.
- The PCSX2 and ShadPS4 downloaded archive files have another archive inside them, so they cannot be properly extracted. Trying to install them will return an error.

## ARM64 Specific Issues

These issues are generally present in Linux, MacOS and Windows when your device uses an ARM64 chip:

- First, it's important to notice that even if there is no official build for ARM64, NiceDeck will download the AMD64 version of the application, which may run via a compatibility layer.
- Unfortunately, many applications still do not provide official builds for the ARM64 architecture. This means that you may experience issues when trying to run such applications.
- Please remember that these issues are not related to NiceDeck, which acts only as a facilitator to install and run such programs.