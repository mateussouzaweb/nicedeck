# Known Issues

## Linux Specific Issues

Luckily, there is not much here, but you should read the [Proton](./Proton.md) section for more insights about that specific implementation.

## Windows Specific Issues

Installing and removing programs:

- When removing Discord, it does not fully remove itself and you will still see the app on Windows.
- Brave Browser removes automatically when requested, but return error on NiceDeck.
- Ubisoft Connect will remove automatically when requested, but the Windows start menu shortcut is not removed.

Windows file and folder shortcuts:

- Windows shortcut ``.lnk`` files or junction points are not accepted or handled by NiceDeck.
- If you want to put additional content in another location with a symbolic link, you must create a valid directory shortcut from Terminal.
- For example, to create a directory shortcut for **ROMs** in another disk (*D:*), open the *Terminal* application with **administrator privilegies**, go to the ``$GAMES`` folder and run the following command to create the symbolic link: ``cmd /c mklink /d ".\ROMs" "D:\ROMs"``.

## MacOS Specific Issues

In general, MacOS gaming still is not a good experience. Do not expect that everything will work flawlessly:

- You may need to give "Full Disk Access" system permissions to the "Terminal" app before trying to use NiceDeck.
- You may need to install Rosetta 2 before trying to install or running non-native applications on Apple Silicon.
- Battle.net package is provided from homebrew, but it does not seems to work.
- GOG Galaxy package is provided from homebrew and requires sudo to be installed. If you try to install it from NiceDeck, the process will hang indefinitely.
- PCSX2 and ShadPS4 downloaded archive files has another archive inside it, so it cannot be properly extracted. Trying to install it will return an error.

## ARM64 Specific Issues

These issues are generally present in Linux, MacOs and Windows when your device uses an ARM64 chip:

- First, is important to notice that even if there is no oficial build for ARM64, NiceDeck will download the AMD64 version of the application that may run over compatibility layer.
- Unfortunately, many applications still does not provide oficial builds for the ARM64 architecture. That means that you may experience issues when trying to run such applications.
- Please remember that those issues are not related to NiceDeck, which acts only as a facilitator to install and run such programs.
