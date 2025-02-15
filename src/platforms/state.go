package platforms

import (
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// State struct
type State struct {
	Platform string   `json:"platform"`
	Emulator string   `json:"emulator"`
	Type     string   `json:"type"`
	Path     string   `json:"path"`
	Source   []string `json:"source"`
}

// Retrieve save state of each platform
func GetStates(options *Options) []*State {

	// The following emulators store saves and states on ROMs directory:
	// - MGBA (user can leave at it is or configure emulator)
	// - MelonDS (user can leave at it is or configure emulator)
	// - Simple64 (cannot be configured as it does not provide options)
	states := []*State{}

	// Lime3DS
	states = append(states, &State{
		Platform: "3DS",
		Emulator: "Lime3DS",
		Type:     "folder",
		Path:     "$STATE/Lime3DS/lime3ds-emu/sdmc",
		Source: []string{
			"$HOME/.var/app/io.github.lime3ds.Lime3DS/data/lime3ds-emu/sdmc",
			"$HOME/.local/share/lime3ds-emu/sdmc",
			"$HOME/.config/lime3ds-emu/sdmc",
			"$HOME/Library/Application Support/Lime3DS/sdmc",
			"$APPDATA\\Roaming\\Lime3DS\\sdmc",
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Lime3DS",
		Type:     "folder",
		Path:     "$STATE/Lime3DS/lime3ds-emu/states",
		Source: []string{
			"$HOME/.var/app/io.github.lime3ds.Lime3DS/data/lime3ds-emu/states",
			"$HOME/.local/share/lime3ds-emu/states",
			"$HOME/.config/lime3ds-emu/states",
			"$HOME/Library/Application Support/Lime3DS/states",
			"$APPDATA\\Roaming\\Lime3DS\\states",
		},
	})

	// Citra
	states = append(states, &State{
		Platform: "3DS",
		Emulator: "Citra",
		Type:     "folder",
		Path:     "$STATE/Citra/citra-emu/sdmc",
		Source: []string{
			"$HOME/.var/app/org.citra_emu.citra/data/citra-emu/sdmc",
			"$HOME/.local/share/citra-emu/sdmc",
			"$HOME/.config/citra-emu/sdmc",
			"$HOME/Library/Application Support/Citra/sdmc",
			"$APPDATA\\Roaming\\Citra\\sdmc",
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Citra",
		Type:     "folder",
		Path:     "$STATE/Citra/citra-emu/states",
		Source: []string{
			"$HOME/.var/app/org.citra_emu.citra/data/citra-emu/states",
			"$HOME/.local/share/citra-emu/states",
			"$HOME/.config/citra-emu/states",
			"$HOME/Library/Application Support/Citra/states",
			"$APPDATA\\Roaming\\Citra\\states",
		},
	})

	// Flycast
	states = append(states, &State{
		Platform: "DC",
		Emulator: "Flycast",
		Type:     "folder",
		Path:     "$STATE/Flycast/flycast",
		Source: []string{
			"$HOME/.var/app/org.flycast.Flycast/data/flycast",
			"$HOME/.local/share/flycast",
			"$HOME/.config/flycast",
			"$HOME/Library/Application Support/Flycast/data",
			"$EMULATORS\\Flycast\\data",
		},
	})

	// Redream
	states = append(states, &State{
		Platform: "DC",
		Emulator: "Redream",
		Type:     "folder",
		Path:     "$STATE/Redream/saves",
		Source: []string{
			"$EMULATORS/Redream/saves",
			"$HOME/.local/share/Redream/saves",
			"$HOME/.config/Redream/saves",
			"$HOME/Library/Application Support/Redream/saves",
			"$EMULATORS\\Redream\\saves",
		},
	})

	// MGBA requires configuration to work:
	// - Go to Tools > Settings > Paths
	// - Set save games location as $HOME/.var/app/io.mgba.mGBA/save
	// - Set save states as $HOME/.var/app/io.mgba.mGBA/states
	// - Make sure that "same directory as the ROM" in both options is unchecked
	states = append(states, &State{
		Platform: "GBA",
		Emulator: "MGBA",
		Type:     "folder",
		Path:     "$STATE/MGBA/save",
		Source: []string{
			"$HOME/.var/app/io.mgba.mGBA/save",
			"$HOME/.local/share/mGBA/save",
			"$HOME/.config/mGBA/save",
			"$HOME/Library/Application Support/mGBA/save",
			"$APPDATA\\Roaming\\mGBA\\save",
			"$EMULATORS\\MGBA\\save",
		},
	}, &State{
		Platform: "GBA",
		Emulator: "MGBA",
		Type:     "folder",
		Path:     "$STATE/MGBA/states",
		Source: []string{
			"$HOME/.var/app/io.mgba.mGBA/states",
			"$HOME/.local/share/mGBA/states",
			"$HOME/.config/mGBA/states",
			"$HOME/Library/Application Support/mGBA/states",
			"$APPDATA\\Roaming\\mGBA\\states",
			"$EMULATORS\\MGBA\\states",
		},
	})

	// Dolphin for GameCube state
	states = append(states, &State{
		Platform: "GC",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/GC",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/GC",
			"$HOME/.local/share/dolphin-emu/GC",
			"$HOME/.config/dolphin-emu/GC",
			"$HOME/Library/Application Support/Dolphin/GC",
			"$APPDATA\\Roaming\\Dolphin Emulator\\GC",
		},
	}, &State{
		Platform: "GC",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
			"$HOME/.local/share/dolphin-emu/StateSaves",
			"$HOME/.config/dolphin-emu/StateSaves",
			"$HOME/Library/Application Support/Dolphin/StateSaves",
			"$APPDATA\\Roaming\\Dolphin Emulator\\StateSaves",
		},
	})

	// MelonDS requires configuration to work:
	// - Go to Config > Path Settings
	// - Set save files path as $HOME/.var/app/net.kuribo64.melonDS/save
	// - Set save states path as $HOME/.var/app/net.kuribo64.melonDS/states
	states = append(states, &State{
		Platform: "NDS",
		Emulator: "MelonDS",
		Type:     "folder",
		Path:     "$STATE/MelonDS/save",
		Source: []string{
			"$HOME/.var/app/net.kuribo64.melonDS/save",
			"$HOME/.local/share/melonDS/save",
			"$HOME/.config/melonDS/save",
			"$HOME/Library/Application Support/melonDS/save",
			"$APPDATA\\Roaming\\melonDS\\save",
			"$EMULATORS\\MelonDS\\save",
		},
	}, &State{
		Platform: "NDS",
		Emulator: "MelonDS",
		Type:     "folder",
		Path:     "$STATE/MelonDS/states",
		Source: []string{
			"$HOME/.var/app/net.kuribo64.melonDS/states",
			"$HOME/.local/share/melonDS/states",
			"$HOME/.config/melonDS/states",
			"$HOME/Library/Application Support/melonDS/states",
			"$APPDATA\\Roaming\\melonDS\\states",
			"$EMULATORS\\MelonDS\\states",
		},
	})

	// DuckStation
	states = append(states, &State{
		Platform: "PS1",
		Emulator: "DuckStation",
		Type:     "folder",
		Path:     "$STATE/DuckStation/duckstation/memcards",
		Source: []string{
			"$HOME/.var/app/org.duckstation.DuckStation/config/duckstation/memcards",
			"$HOME/.local/share/duckstation/memcards",
			"$HOME/.config/duckstation/memcards",
			"$HOME/Library/Application Support/DuckStation/memcards",
			"$HOME\\Documents\\DuckStation\\memcards",
		},
	}, &State{
		Platform: "PS1",
		Emulator: "DuckStation",
		Type:     "folder",
		Path:     "$STATE/DuckStation/duckstation/savestates",
		Source: []string{
			"$HOME/.var/app/org.duckstation.DuckStation/config/duckstation/savestates",
			"$HOME/.local/share/duckstation/savestates",
			"$HOME/.config/duckstation/savestates",
			"$HOME/Library/Application Support/DuckStation/savestates",
			"$HOME\\Documents\\DuckStation\\savestates",
		},
	})

	// PCSX2
	states = append(states, &State{
		Platform: "PS2",
		Emulator: "PCSX2",
		Type:     "folder",
		Path:     "$STATE/PCSX2/memcards",
		Source: []string{
			"$HOME/.var/app/net.pcsx2.PCSX2/config/PCSX2/memcards",
			"$HOME/.local/share/PCSX2/memcards",
			"$HOME/.config/PCSX2/memcards",
			"$HOME/Library/Application Support/PCSX2/memcards",
			"$HOME\\Documents\\PCSX2\\memcards",
		},
	}, &State{
		Platform: "PS2",
		Emulator: "PCSX2",
		Type:     "folder",
		Path:     "$STATE/PCSX2/sstates",
		Source: []string{
			"$HOME/.var/app/net.pcsx2.PCSX2/config/PCSX2/sstates",
			"$HOME/.local/share/PCSX2/sstates",
			"$HOME/.config/PCSX2/sstates",
			"$HOME/Library/Application Support/PCSX2/sstates",
			"$HOME\\Documents\\PCSX2\\sstates",
		},
	})

	// RPCS3
	states = append(states, &State{
		Platform: "PS3",
		Emulator: "RPCS3",
		Type:     "folder",
		Path:     "$STATE/RPCS3/rpcs3/dev_hdd0/home/00000001/savedata",
		Source: []string{
			"$HOME/.var/app/net.rpcs3.RPCS3/config/rpcs3/dev_hdd0/home/00000001/savedata",
			"$HOME/.local/share/rpcs3/dev_hdd0/home/00000001/savedata",
			"$HOME/.config/rpcs3/dev_hdd0/home/00000001/savedata",
			"$HOME/Library/Application Support/rpcs3/dev_hdd0/home/00000001/savedata",
			"$EMULATORS\\RPCS3\\dev_hdd0\\home\\00000001\\savedata",
		},
	})

	// ShadPS4
	states = append(states, &State{
		Platform: "PS4",
		Emulator: "ShadPS4",
		Type:     "folder",
		Path:     "$STATE/ShadPS4/saves",
		Source: []string{
			"$HOME/.local/share/shadps4/saves",
			"$HOME/.config/shadps4/saves",
			"$HOME/Library/Application Support/shadPS4/saves",
			"$EMULATORS\\ShadPS4\\user\\saves",
		},
	})

	// PPSSPP
	states = append(states, &State{
		Platform: "PSP",
		Emulator: "PPSSPP",
		Type:     "folder",
		Path:     "$STATE/PPSSPP/ppsspp/PSP/SAVEDATA",
		Source: []string{
			"$HOME/.var/app/org.ppsspp.PPSSPP/config/ppsspp/PSP/SAVEDATA",
			"$HOME/.local/share/ppsspp/PSP/SAVEDATA",
			"$HOME/.config/ppsspp/PSP/SAVEDATA",
			"$HOME/Library/Application Support/PPSSPP/PSP/SAVEDATA",
			"$EMULATORS\\PPSSPP\\memstick\\PSP\\SAVEDATA",
		},
	}, &State{
		Platform: "PSP",
		Emulator: "PPSSPP",
		Type:     "folder",
		Path:     "$STATE/PPSSPP/ppsspp/PSP/PPSSPP_STATE",
		Source: []string{
			"$HOME/.var/app/org.ppsspp.PPSSPP/config/ppsspp/PSP/PPSSPP_STATE",
			"$HOME/.local/share/ppsspp/PSP/PPSSPP_STATE",
			"$HOME/.config/ppsspp/PSP/PPSSPP_STATE",
			"$HOME/Library/Application Support/PPSSPP/PSP/PPSSPP_STATE",
			"$EMULATORS\\PPSSPP\\memstick\\PSP\\PPSSPP_STATE",
		},
	})

	// Vita3K
	states = append(states, &State{
		Platform: "PSVITA",
		Emulator: "Vita3K",
		Type:     "folder",
		Path:     "$STATE/Vita3K/ux0/user/00/savedata",
		Source: []string{
			"$HOME/.local/share/Vita3K/Vita3K/ux0/user/00/savedata",
			"$HOME/.config/Vita3K/Vita3K/ux0/user/00/savedata",
			"$HOME/Library/Application Support/Vita3K/Vita3K/fs/ux0/user/00/savedata",
			"$APPDATA\\Roaming\\Vita3K\\Vita3k\\ux0\\user\\00\\savedata",
		},
	})

	// Citron
	states = append(states, &State{
		Platform: "SWITCH",
		Emulator: "Citron",
		Type:     "folder",
		Path:     "$STATE/Citron/nand/user/save",
		Source: []string{
			"$HOME/.local/share/citron/nand/user/save",
			"$HOME/.config/citron/nand/user/save",
			"$EMULATORS\\Citron\\user\\nand\\user\\save",
		},
	})

	// Yuzu
	states = append(states, &State{
		Platform: "SWITCH",
		Emulator: "Yuzu",
		Type:     "folder",
		Path:     "$STATE/Yuzu/nand/user/save",
		Source: []string{
			"$HOME/.var/app/org.yuzu_emu.yuzu/data/nand/user/save",
			"$HOME/.local/share/yuzu/nand/user/save",
			"$HOME/.config/yuzu/nand/user/save",
			"$APPDATA\\Roaming\\Yuzu\\nand\\user\\save",
		},
	})

	// Ryujinx save state in two folders
	// We also sync the profiles.json file to avoid losing user reference
	states = append(states, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "folder",
		Path:     "$STATE/Ryujinx/bis/user/save",
		Source: []string{
			"$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/save",
			"$HOME/.local/share/Ryujinx/bis/user/save",
			"$HOME/.config/Ryujinx/bis/user/save",
			"$HOME/Library/Application Support/Ryujinx/bis/user/save",
			"$APPDATA\\Roaming\\Ryujinx\\bis\\user\\save",
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "folder",
		Path:     "$STATE/Ryujinx/bis/user/saveMeta",
		Source: []string{
			"$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/saveMeta",
			"$HOME/.local/share/Ryujinx/bis/user/saveMeta",
			"$HOME/.config/Ryujinx/bis/user/saveMeta",
			"$HOME/Library/Application Support/Ryujinx/bis/user/saveMeta",
			"$APPDATA\\Roaming\\Ryujinx\\bis\\user\\saveMeta",
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "file",
		Path:     "$STATE/Ryujinx/system/Profiles.json",
		Source: []string{
			"$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/system/Profiles.json",
			"$HOME/.local/share/Ryujinx/system/Profiles.json",
			"$HOME/.config/Ryujinx/system/Profiles.json",
			"$HOME/Library/Application Support/Ryujinx/system/Profiles.json",
			"$APPDATA\\Roaming\\Ryujinx\\system\\Profiles.json",
		},
	})

	// Cemu
	states = append(states, &State{
		Platform: "WIIU",
		Emulator: "Cemu",
		Type:     "folder",
		Path:     "$STATE/Cemu/mlc01/usr/save",
		Source: []string{
			"$HOME/.var/app/info.cemu.Cemu/data/Cemu/mlc01/usr/save",
			"HOME/.local/share/Cemu/mlc01/usr/save",
			"HOME/.config/Cemu/mlc01/usr/save",
			"$HOME/Library/Application Support/Cemu/mlc01/usr/save",
			"$APPDATA\\Roaming\\Cemu\\mlc01\\usr\\save",
		},
	})

	// Dolphin for Wii state
	states = append(states, &State{
		Platform: "WII",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/Wii",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/Wii",
			"$HOME/.local/share/dolphin-emu/Wii",
			"$HOME/.config/dolphin-emu/Wii",
			"$HOME/Library/Application Support/Dolphin/Wii",
			"$APPDATA\\Roaming\\Dolphin Emulator\\Wii",
		},
	}, &State{
		Platform: "WII",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
			"$HOME/.local/share/dolphin-emu/StateSaves",
			"$HOME/.config/dolphin-emu/StateSaves",
			"$HOME/Library/Application Support/Dolphin/StateSaves",
			"$APPDATA\\Roaming\\Dolphin Emulator\\StateSaves",
		},
	})

	// Xemu requires two files to transfer states
	// The HD file is expected to be in the $BIOS/XBOX/xbox_hdd.qcow2 location by default
	// However, we also put it here in case of the user put this file in xemu/ folder
	states = append(states, &State{
		Platform: "XBOX",
		Emulator: "Xemu",
		Type:     "file",
		Path:     "$STATE/Xemu/xemu/xemu/eeprom.bin",
		Source: []string{
			"$HOME/.var/app/app.xemu.xemu/data/xemu/xemu/eeprom.bin",
			"$HOME/.local/share/xemu/xemu/eeprom.bin",
			"$HOME/.config/xemu/xemu/eeprom.bin",
			"$HOME/Library/Application Support/xemu/xemu/eeprom.bin",
			"$APPDATA\\Roaming\\xemu\\xemu\\eeprom.bin",
		},
	}, &State{
		Platform: "XBOX",
		Emulator: "Xemu",
		Type:     "file",
		Path:     "$STATE/Xemu/xemu/xemu/xbox_hdd.qcow2",
		Source: []string{
			"$HOME/.var/app/app.xemu.xemu/data/xemu/xemu/xbox_hdd.qcow2",
			"$HOME/.local/share/xemu/xemu/xbox_hdd.qcow2",
			"$HOME/.config/xemu/xemu/xbox_hdd.qcow2",
			"$HOME/Library/Application Support/xemu/xemu/xbox_hdd.qcow2",
			"$APPDATA\\Roaming\\xemu\\xemu\\xbox_hdd.qcow2",
		},
	})

	// Xenia
	// Save file is expected to be in the Documents folder
	// However, user can also put this file in Xenia/ folder
	states = append(states, &State{
		Platform: "X360",
		Emulator: "Xenia",
		Type:     "file",
		Path:     "$STATE/Xenia/content",
		Source: []string{
			"$EMULATORS/Xenia/content",
			"$EMULATORS\\Xenia\\content",
			"$HOME\\Documents\\xenia\\content",
		},
	})

	return states
}

// Sync state of each platform
func SyncState(options *Options) error {

	// Default action is copy from source to destination path as backup method
	// However, user can choose to restore state with optional preference
	restoreState := slices.Contains(options.Preferences, "restore-state")

	// Process each state
	for _, state := range GetStates(options) {

		// Check if should process this platform
		if !slices.Contains(options.Platforms, state.Platform) {
			continue
		}

		// Source are in multiple locations due to multiple runtimes and operating systems
		// To ensure compatibility, we process just the first valid location for source
		for _, source := range state.Source {

			// Fill source and destination information
			source := fs.ExpandPath(source)
			destination := fs.ExpandPath(state.Path)

			if restoreState {
				source = fs.ExpandPath(state.Path)
				destination = fs.ExpandPath(source)
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

			// Ensure that only the first valid result will be processed
			break
		}
	}

	cli.Printf(cli.ColorNotice, "State synchronized.\n")
	return nil
}
