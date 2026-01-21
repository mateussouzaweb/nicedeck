package macos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Basic MacOS .app structure
// AppName.app/
// └── Contents/
//    ├── Info.plist
//    ├── Resources/
//    │   └── AppIcon.icns
//    └── MacOS/
//        └── launcher

// Bundle represents the basic structure in a .app directory:
type Bundle struct {
	AppName          string   `json:"appName"`
	BundleID         string   `json:"bundleID"`
	Launcher         string   `json:"launcher"`
	IconPath         string   `json:"iconPath"`
	WorkingDirectory string   `json:"workingDirectory"`
	Executable       string   `json:"executable"`
	Arguments        []string `json:"arguments"`
	Environment      []string `json:"environment"`
}

// Write the bundled application into destination folder
func WriteBundle(destination string, bundle *Bundle) error {

	contents := filepath.Join(destination, "Contents")
	macOS := filepath.Join(destination, "Contents", "MacOS")
	resources := filepath.Join(destination, "Contents", "Resources")
	appIconPath := filepath.Join(resources, "AppIcon.icns")
	launcherPath := filepath.Join(macOS, bundle.Launcher)
	infoPlistPath := filepath.Join(contents, "Info.plist")

	// Create necessary directories
	err := os.MkdirAll(contents, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(macOS, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(resources, os.ModePerm)
	if err != nil {
		return err
	}

	// Copy .icns to location when file already exists
	if filepath.Ext(bundle.IconPath) == ".icns" {
		err = fs.CopyFile(bundle.IconPath, appIconPath, true)
		if err != nil {
			return err
		}
	}

	// Convert .png file to .icns if necessary
	if filepath.Ext(bundle.IconPath) == ".png" {

		iconSetPath := fmt.Sprintf("%s/icon.iconset", resources)
		convertScript := fmt.Sprintf(``+
			`mkdir "%s";`+
			`sips -z 16 16 "%s" --out "%s/icon_16x16.png" > /dev/null;`+
			`sips -z 32 32 "%s" --out "%s/icon_16x16@2x.png" > /dev/null;`+
			`sips -z 32 32 "%s" --out "%s/icon_32x32.png" > /dev/null;`+
			`sips -z 64 64 "%s" --out "%s/icon_32x32@2x.png" > /dev/null;`+
			`sips -z 128 128 "%s" --out "%s/icon_128x128.png" > /dev/null;`+
			`sips -z 256 256 "%s" --out "%s/icon_128x128@2x.png" > /dev/null;`+
			`sips -z 256 256 "%s" --out "%s/icon_256x256.png" > /dev/null;`+
			`sips -z 512 512 "%s" --out "%s/icon_256x256@2x.png" > /dev/null;`+
			`sips -z 512 512 "%s" --out "%s/icon_512x512.png" > /dev/null;`+
			`sips -z 1024 1024 "%s" --out "%s/icon_512x512@2x.png" > /dev/null;`+
			`iconutil -c icns "%s" -o "%s" > /dev/null;`+
			`rm -Rf "%s"`,
			iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			bundle.IconPath, iconSetPath,
			iconSetPath, appIconPath,
			iconSetPath,
		)

		command := cli.Command(convertScript)
		err = cli.Run(command)
		if err != nil {
			return err
		}
	}

	// Create launcher script for the appBundle
	// Please note that this launcher script expect to call another .app
	launcherScript := fmt.Sprintf(``+
		`#!/bin/bash`+"\n"+
		`cd "%s" && %s open -n "%s" --args %s`,
		cli.Unquote(bundle.WorkingDirectory),
		strings.Join(bundle.Environment, " "),
		cli.Unquote(bundle.Executable),
		strings.Join(bundle.Arguments, " "),
	)

	// Write launcher script file
	err = fs.WriteFile(launcherPath, launcherScript)
	if err != nil {
		return err
	}

	// Make launcher script executable
	err = os.Chmod(launcherPath, 0755)
	if err != nil {
		return err
	}

	// Create Info.plist file
	infoPlist := fmt.Sprintf(``+
		`<?xml version="1.0" encoding="UTF-8"?>`+"\n"+
		`<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">`+"\n"+
		`<plist version="1.0">`+"\n"+
		`  <dict>`+"\n"+
		`    <key>CFBundleName</key>`+"\n"+
		`    <string>%s</string>`+"\n"+
		`    <key>CFBundleDisplayName</key>`+"\n"+
		`    <string>%s</string>`+"\n"+
		`    <key>CFBundleIdentifier</key>`+"\n"+
		`    <string>%s</string>`+"\n"+
		`    <key>CFBundleExecutable</key>`+"\n"+
		`    <string>%s</string>`+"\n"+
		`    <key>CFBundleIconFile</key>`+"\n"+
		`    <string>AppIcon</string>`+"\n"+
		`    <key>CFBundlePackageType</key>`+"\n"+
		`    <string>APPL</string>`+"\n"+
		`    <key>LSUIElement</key>`+"\n"+
		`    <false/>`+"\n"+
		`  </dict>`+"\n"+
		`</plist>`,
		bundle.AppName,
		bundle.AppName,
		bundle.BundleID,
		bundle.Launcher,
	)

	// Write Info.plist file
	err = fs.WriteFile(infoPlistPath, infoPlist)
	if err != nil {
		return err
	}

	return nil
}
