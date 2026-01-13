package macos

import (
	"fmt"
	"os"
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
	IconPath         string   `json:"iconPath"`
	WorkingDirectory string   `json:"workingDirectory"`
	Executable       string   `json:"executable"`
	Arguments        []string `json:"arguments"`
	Environment      []string `json:"environment"`
}

// Write the bundled application into destination folder
func WriteBundle(destination string, bundle *Bundle) error {

	contents := fmt.Sprintf("%s/Contents", destination)
	macOS := fmt.Sprintf("%s/Contents/MacOS", destination)
	resources := fmt.Sprintf("%s/Contents/Resources", destination)

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

	// Convert PNG file to .icns if provided
	if bundle.IconPath != "" {

		iconSetPath := fmt.Sprintf("%s/icon.iconset", resources)
		appIconPath := fmt.Sprintf("%s/AppIcon.icns", resources)
		script := fmt.Sprintf(``+
			`mkdir "%s";`+
			`sips -z 16 16 "%s" --out "%s/icon_16x16.png";`+
			`sips -z 32 32 "%s" --out "%s/icon_16x16@2x.png";`+
			`sips -z 32 32 "%s" --out "%s/icon_32x32.png";`+
			`sips -z 64 64 "%s" --out "%s/icon_32x32@2x.png";`+
			`sips -z 128 128 "%s" --out "%s/icon_128x128.png";`+
			`sips -z 256 256 "%s" --out "%s/icon_128x128@2x.png";`+
			`sips -z 256 256 "%s" --out "%s/icon_256x256.png";`+
			`sips -z 512 512 "%s" --out "%s/icon_256x256@2x.png";`+
			`sips -z 512 512 "%s" --out "%s/icon_512x512.png";`+
			`sips -z 1024 1024 "%s" --out "%s/icon_512x512@2x.png";`+
			`iconutil -c icns "%s" -o "%s";`+
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

		command := cli.Command(script)
		err = cli.Run(command)
		if err != nil {
			return err
		}
	}

	// Create launcher script
	launcherPath := fmt.Sprintf("%s/launcher", macOS)
	launcherScript := fmt.Sprintf(``+
		`#!/bin/bash`+"\n"+
		`cd "%s" && exec %s open -n "%s" --args %s`,
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
	infoPlistPath := fmt.Sprintf("%s/Info.plist", contents)
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
		`    <string>launcher</string>`+"\n"+
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
	)

	// Write Info.plist file
	err = fs.WriteFile(infoPlistPath, infoPlist)
	if err != nil {
		return err
	}

	return nil
}
