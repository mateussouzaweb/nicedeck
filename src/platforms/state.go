package platforms

import (
	"os"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// State struct
type State struct {
	Platform    string `json:"platform"`
	Emulator    string `json:"emulator"`
	Type        string `json:"type"`
	Source      string `json:"path"`
	Destination string `json:"destination"`
}

// Retrieve save state of each platform
func GetStates(options *Options) []*State {

	// The following emulators store saves and states on ROMs directory:
	// - MGBA (user can leave at it is or configure emulator)
	// - MelonDS (user can leave at it is or configure emulator)
	// - Simple64 (cannot be configured as it does not provide options)
	states := []*State{}

	// Citra
	states = append(states, &State{
		Platform:    "3DS",
		Emulator:    "Citra",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.citra_emu.citra/data/citra-emu/sdmc",
		Destination: "$STATE/Citra/citra-emu/sdmc",
	}, &State{
		Platform:    "3DS",
		Emulator:    "Citra",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.citra_emu.citra/data/citra-emu/states",
		Destination: "$STATE/Citra/citra-emu/states",
	})

	// Flycast
	states = append(states, &State{
		Platform:    "DC",
		Emulator:    "Flycast",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.flycast.Flycast/data/flycast",
		Destination: "$STATE/Flycast/flycast",
	})

	// MGBA requires configuration to work:
	// - Go to Tools > Settings > Paths
	// - Set save games location as $HOME/.var/app/io.mgba.mGBA/save
	// - Set save states as $HOME/.var/app/io.mgba.mGBA/states
	// - Make sure that "same directory as the ROM" in both options is unchecked
	states = append(states, &State{
		Platform:    "GBA",
		Emulator:    "MGBA",
		Type:        "folder",
		Source:      "$HOME/.var/app/io.mgba.mGBA/save",
		Destination: "$STATE/MGBA/save",
	}, &State{
		Platform:    "GBA",
		Emulator:    "MGBA",
		Type:        "folder",
		Source:      "$HOME/.var/app/io.mgba.mGBA/states",
		Destination: "$STATE/MGBA/states",
	})

	// Dolphin for GameCube state
	states = append(states, &State{
		Platform:    "GC",
		Emulator:    "Dolphin",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/GC",
		Destination: "$STATE/Dolphin/dolphin-emu/GC",
	}, &State{
		Platform:    "GC",
		Emulator:    "Dolphin",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
		Destination: "$STATE/Dolphin/dolphin-emu/StateSaves",
	})

	// MelonDS requires configuration to work:
	// - Go to Config > Path Settings
	// - Set save files path as $HOME/.var/app/net.kuribo64.melonDS/save
	// - Set save states path as $HOME/.var/app/net.kuribo64.melonDS/states
	states = append(states, &State{
		Platform:    "NDS",
		Emulator:    "MelonDS",
		Type:        "folder",
		Source:      "$HOME/.var/app/net.kuribo64.melonDS/save",
		Destination: "$STATE/MelonDS/save",
	}, &State{
		Platform:    "NDS",
		Emulator:    "MelonDS",
		Type:        "folder",
		Source:      "$HOME/.var/app/net.kuribo64.melonDS/states",
		Destination: "$STATE/MelonDS/states",
	})

	// DuckStation
	states = append(states, &State{
		Platform:    "PS1",
		Emulator:    "DuckStation",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.duckstation.DuckStation/config/duckstation/memcards",
		Destination: "$STATE/DuckStation/duckstation/memcards",
	}, &State{
		Platform:    "PS1",
		Emulator:    "DuckStation",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.duckstation.DuckStation/config/duckstation/savestates",
		Destination: "$STATE/DuckStation/duckstation/savestates",
	})

	// PCSX2
	states = append(states, &State{
		Platform:    "PS2",
		Emulator:    "PCSX2",
		Type:        "folder",
		Source:      "$HOME/.var/app/net.pcsx2.PCSX2/config/PCSX2/memcards",
		Destination: "$STATE/PCSX2/memcards",
	}, &State{
		Platform:    "PS2",
		Emulator:    "PCSX2",
		Type:        "folder",
		Source:      "$HOME/.var/app/net.pcsx2.PCSX2/config/PCSX2/sstates",
		Destination: "$STATE/PCSX2/sstates",
	})

	// RPCS3
	states = append(states, &State{
		Platform:    "PS3",
		Emulator:    "RPCS3",
		Type:        "folder",
		Source:      "$HOME/.var/app/net.rpcs3.RPCS3/config/rpcs3/dev_hdd0/home/00000001/savedata",
		Destination: "$STATE/RPCS3/rpcs3/dev_hdd0/home/00000001/savedata",
	})

	// PPSSPP
	states = append(states, &State{
		Platform:    "PSP",
		Emulator:    "PPSSPP",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.ppsspp.PPSSPP/config/ppsspp/PSP/SAVEDATA",
		Destination: "$STATE/PPSSPP/ppsspp/PSP/SAVEDATA",
	}, &State{
		Platform:    "PSP",
		Emulator:    "PPSSPP",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.ppsspp.PPSSPP/config/ppsspp/PSP/PPSSPP_STATE",
		Destination: "$STATE/PPSSPP/ppsspp/PSP/PPSSPP_STATE",
	})

	// Yuzu
	states = append(states, &State{
		Platform:    "SWITCH",
		Emulator:    "Yuzu",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.yuzu_emu.yuzu/data/nand/user/save",
		Destination: "$STATE/Yuzu/nand/user/save",
	})

	// Ryujinx save state in two folders
	// We also sync the profiles.json file to avoid losing user reference
	states = append(states, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/save",
		Destination: "$STATE/Ryujinx/bis/user/save",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/saveMeta",
		Destination: "$STATE/Ryujinx/bis/user/saveMeta",
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "file",
		Source:      "$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/system/Profiles.json",
		Destination: "$STATE/Ryujinx/system/Profiles.json",
	})

	// Cemu
	states = append(states, &State{
		Platform:    "WIIU",
		Emulator:    "Cemu",
		Type:        "folder",
		Source:      "$HOME/.var/app/info.cemu.Cemu/data/Cemu/mlc01/usr/save",
		Destination: "$STATE/Cemu/mlc01/usr/save",
	})

	// Dolphin for Wii state
	states = append(states, &State{
		Platform:    "WII",
		Emulator:    "Dolphin",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/Wii",
		Destination: "$STATE/Dolphin/dolphin-emu/Wii",
	}, &State{
		Platform:    "WII",
		Emulator:    "Dolphin",
		Type:        "folder",
		Source:      "$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
		Destination: "$STATE/Dolphin/dolphin-emu/StateSaves",
	})

	// Xemu requires two files to transfer states
	// The HD file is expected to be in the $BIOS/XBOX/xbox_hdd.qcow2 location by default
	// However, we also put it here in case of the user put this file in xemu/ folder
	states = append(states, &State{
		Platform:    "XBOX",
		Emulator:    "Xemu",
		Type:        "file",
		Source:      "$HOME/.var/app/app.xemu.xemu/data/xemu/xemu/eeprom.bin",
		Destination: "$STATE/Xemu/xemu/xemu/eeprom.bin",
	}, &State{
		Platform:    "XBOX",
		Emulator:    "Xemu",
		Type:        "file",
		Source:      "$HOME/.var/app/app.xemu.xemu/data/xemu/xemu/xbox_hdd.qcow2",
		Destination: "$STATE/Xemu/xemu/xemu/xbox_hdd.qcow2",
	})

	return states
}

// Sync state of each platform
func SyncState(options *Options) error {

	// Default action is copy from source to destination as backup method
	// However, user can choose to restore state with optional preference
	restoreState := slices.Contains(options.Preferences, "restore-state")

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// Fill source and destination information
		source := os.ExpandEnv(state.Source)
		destination := os.ExpandEnv(state.Destination)

		if restoreState {
			source = os.ExpandEnv(state.Destination)
			destination = os.ExpandEnv(state.Source)
		}

		// Process file or folder state
		if state.Type == "file" {

			// Ensure file exist
			exist, err := fs.FileExist(source)
			if err != nil {
				return err
			} else if !exist {
				cli.Printf(cli.ColorNotice, "Skipping file not detected: %s\n", source)
				continue
			}

			// Copy file
			cli.Printf(cli.ColorNotice, "Synchronizing file from %s to %s...\n", source, destination)
			err = fs.CopyFile(source, destination)
			if err != nil {
				return err
			}

		} else if state.Type == "folder" {

			// Ensure folder exist
			exist, err := fs.DirectoryExist(source)
			if err != nil {
				return err
			} else if !exist {
				cli.Printf(cli.ColorNotice, "Skipping folder not detected: %s\n", source)
				continue
			}

			// Recursive copy content
			cli.Printf(cli.ColorNotice, "Synchronizing folder from %s to %s...\n", source, destination)
			err = fs.CopyDirectory(source, destination)
			if err != nil {
				return err
			}

		}
	}

	cli.Printf(cli.ColorNotice, "State synchronized.\n")
	return nil
}
