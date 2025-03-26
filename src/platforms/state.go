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

	// Azahar
	states = append(states, &State{
		Platform: "3DS",
		Emulator: "Azahar",
		Type:     "folder",
		Path:     "$STATE/Azahar/azahar-emu/sdmc",
		Source: []string{
			"$SHARE/azahar-emu/sdmc",
			"$CONFIG/azahar-emu/sdmc",
			"$CONFIG/Azahar/sdmc",
			"$CONFIG\\Azahar\\sdmc",
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Azahar",
		Type:     "folder",
		Path:     "$STATE/Azahar/azahar-emu/states",
		Source: []string{
			"$SHARE/azahar-emu/states",
			"$CONFIG/azahar-emu/states",
			"$CONFIG/Azahar/states",
			"$CONFIG\\Azahar\\states",
		},
	})

	// Lime3DS
	states = append(states, &State{
		Platform: "3DS",
		Emulator: "Lime3DS",
		Type:     "folder",
		Path:     "$STATE/Lime3DS/lime3ds-emu/sdmc",
		Source: []string{
			"$VAR/io.github.lime3ds.Lime3DS/data/lime3ds-emu/sdmc",
			"$SHARE/lime3ds-emu/sdmc",
			"$CONFIG/lime3ds-emu/sdmc",
			"$CONFIG/Lime3DS/sdmc",
			"$CONFIG\\Lime3DS\\sdmc",
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Lime3DS",
		Type:     "folder",
		Path:     "$STATE/Lime3DS/lime3ds-emu/states",
		Source: []string{
			"$VAR/io.github.lime3ds.Lime3DS/data/lime3ds-emu/states",
			"$SHARE/lime3ds-emu/states",
			"$CONFIG/lime3ds-emu/states",
			"$CONFIG/Lime3DS/states",
			"$CONFIG\\Lime3DS\\states",
		},
	})

	// Flycast
	states = append(states, &State{
		Platform: "DC",
		Emulator: "Flycast",
		Type:     "folder",
		Path:     "$STATE/Flycast/flycast",
		Source: []string{
			"$VAR/org.flycast.Flycast/data/flycast",
			"$SHARE/flycast",
			"$CONFIG/flycast/data",
			"$CONFIG/Flycast/data",
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
			"$SHARE/Redream/saves",
			"$CONFIG/Redream/saves",
			"$EMULATORS/Redream/saves",
			"$EMULATORS\\Redream\\saves",
		},
	})

	// MGBA requires configuration to work:
	// - Go to Tools > Settings > Paths
	// - Set save games location as $VAR/io.mgba.mGBA/save
	// - Set save states as $VAR/io.mgba.mGBA/states
	// - Make sure that "same directory as the ROM" in both options is unchecked
	states = append(states, &State{
		Platform: "GBA",
		Emulator: "MGBA",
		Type:     "folder",
		Path:     "$STATE/MGBA/save",
		Source: []string{
			"$VAR/io.mgba.mGBA/save",
			"$SHARE/mGBA/save",
			"$CONFIG/mGBA/save",
			"$CONFIG\\mGBA\\save",
			"$EMULATORS\\MGBA\\save",
		},
	}, &State{
		Platform: "GBA",
		Emulator: "MGBA",
		Type:     "folder",
		Path:     "$STATE/MGBA/states",
		Source: []string{
			"$VAR/io.mgba.mGBA/states",
			"$SHARE/mGBA/states",
			"$CONFIG/mGBA/states",
			"$CONFIG\\mGBA\\states",
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
			"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/GC",
			"$SHARE/dolphin-emu/GC",
			"$CONFIG/dolphin-emu/GC",
			"$CONFIG/Dolphin/GC",
			"$CONFIG\\Dolphin Emulator\\GC",
		},
	}, &State{
		Platform: "GC",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
			"$SHARE/dolphin-emu/StateSaves",
			"$CONFIG/dolphin-emu/StateSaves",
			"$CONFIG/Dolphin/StateSaves",
			"$CONFIG\\Dolphin Emulator\\StateSaves",
		},
	})

	// MelonDS requires configuration to work:
	// - Go to Config > Path Settings
	// - Set save files path as $VAR/net.kuribo64.melonDS/save
	// - Set save states path as $VAR/net.kuribo64.melonDS/states
	states = append(states, &State{
		Platform: "NDS",
		Emulator: "MelonDS",
		Type:     "folder",
		Path:     "$STATE/MelonDS/save",
		Source: []string{
			"$VAR/net.kuribo64.melonDS/save",
			"$SHARE/melonDS/save",
			"$CONFIG/melonDS/save",
			"$CONFIG\\melonDS\\save",
			"$EMULATORS\\MelonDS\\save",
		},
	}, &State{
		Platform: "NDS",
		Emulator: "MelonDS",
		Type:     "folder",
		Path:     "$STATE/MelonDS/states",
		Source: []string{
			"$VAR/net.kuribo64.melonDS/states",
			"$SHARE/melonDS/states",
			"$CONFIG/melonDS/states",
			"$CONFIG\\melonDS\\states",
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
			"$VAR/org.duckstation.DuckStation/config/duckstation/memcards",
			"$SHARE/duckstation/memcards",
			"$CONFIG/duckstation/memcards",
			"$CONFIG/DuckStation/memcards",
			"$DOCUMENTS\\DuckStation\\memcards",
		},
	}, &State{
		Platform: "PS1",
		Emulator: "DuckStation",
		Type:     "folder",
		Path:     "$STATE/DuckStation/duckstation/savestates",
		Source: []string{
			"$VAR/org.duckstation.DuckStation/config/duckstation/savestates",
			"$SHARE/duckstation/savestates",
			"$CONFIG/duckstation/savestates",
			"$CONFIG/DuckStation/savestates",
			"$DOCUMENTS\\DuckStation\\savestates",
		},
	})

	// PCSX2
	states = append(states, &State{
		Platform: "PS2",
		Emulator: "PCSX2",
		Type:     "folder",
		Path:     "$STATE/PCSX2/memcards",
		Source: []string{
			"$VAR/net.pcsx2.PCSX2/config/PCSX2/memcards",
			"$SHARE/PCSX2/memcards",
			"$CONFIG/PCSX2/memcards",
			"$DOCUMENTS\\PCSX2\\memcards",
		},
	}, &State{
		Platform: "PS2",
		Emulator: "PCSX2",
		Type:     "folder",
		Path:     "$STATE/PCSX2/sstates",
		Source: []string{
			"$VAR/net.pcsx2.PCSX2/config/PCSX2/sstates",
			"$SHARE/PCSX2/sstates",
			"$CONFIG/PCSX2/sstates",
			"$DOCUMENTS\\PCSX2\\sstates",
		},
	})

	// RPCS3
	states = append(states, &State{
		Platform: "PS3",
		Emulator: "RPCS3",
		Type:     "folder",
		Path:     "$STATE/RPCS3/rpcs3/dev_hdd0/home/00000001/savedata",
		Source: []string{
			"$VAR/net.rpcs3.RPCS3/config/rpcs3/dev_hdd0/home/00000001/savedata",
			"$SHARE/rpcs3/dev_hdd0/home/00000001/savedata",
			"$CONFIG/rpcs3/dev_hdd0/home/00000001/savedata",
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
			"$SHARE/shadps4/saves",
			"$CONFIG/shadps4/saves",
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
			"$VAR/org.ppsspp.PPSSPP/config/ppsspp/PSP/SAVEDATA",
			"$SHARE/ppsspp/PSP/SAVEDATA",
			"$CONFIG/ppsspp/PSP/SAVEDATA",
			"$CONFIG/PPSSPP/PSP/SAVEDATA",
			"$EMULATORS\\PPSSPP\\memstick\\PSP\\SAVEDATA",
		},
	}, &State{
		Platform: "PSP",
		Emulator: "PPSSPP",
		Type:     "folder",
		Path:     "$STATE/PPSSPP/ppsspp/PSP/PPSSPP_STATE",
		Source: []string{
			"$VAR/org.ppsspp.PPSSPP/config/ppsspp/PSP/PPSSPP_STATE",
			"$SHARE/ppsspp/PSP/PPSSPP_STATE",
			"$CONFIG/ppsspp/PSP/PPSSPP_STATE",
			"$CONFIG/PPSSPP/PSP/PPSSPP_STATE",
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
			"$SHARE/Vita3K/Vita3K/ux0/user/00/savedata",
			"$CONFIG/Vita3K/Vita3K/ux0/user/00/savedata",
			"$CONFIG/Vita3K/Vita3K/fs/ux0/user/00/savedata",
			"$CONFIG\\Vita3K\\Vita3k\\ux0\\user\\00\\savedata",
		},
	})

	// Citron
	states = append(states, &State{
		Platform: "SWITCH",
		Emulator: "Citron",
		Type:     "folder",
		Path:     "$STATE/Citron/nand/user/save",
		Source: []string{
			"$SHARE/citron/nand/user/save",
			"$CONFIG/citron/nand/user/save",
			"$EMULATORS\\Citron\\user\\nand\\user\\save",
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
			"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/save",
			"$SHARE/Ryujinx/bis/user/save",
			"$CONFIG/Ryujinx/bis/user/save",
			"$CONFIG\\Ryujinx\\bis\\user\\save",
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "folder",
		Path:     "$STATE/Ryujinx/bis/user/saveMeta",
		Source: []string{
			"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/saveMeta",
			"$SHARE/Ryujinx/bis/user/saveMeta",
			"$CONFIG/Ryujinx/bis/user/saveMeta",
			"$CONFIG\\Ryujinx\\bis\\user\\saveMeta",
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "file",
		Path:     "$STATE/Ryujinx/system/Profiles.json",
		Source: []string{
			"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/system/Profiles.json",
			"$SHARE/Ryujinx/system/Profiles.json",
			"$CONFIG/Ryujinx/system/Profiles.json",
			"$CONFIG\\Ryujinx\\system\\Profiles.json",
		},
	})

	// Cemu
	states = append(states, &State{
		Platform: "WIIU",
		Emulator: "Cemu",
		Type:     "folder",
		Path:     "$STATE/Cemu/mlc01/usr/save",
		Source: []string{
			"$VAR/info.cemu.Cemu/data/Cemu/mlc01/usr/save",
			"$SHARE/Cemu/mlc01/usr/save",
			"$CONFIG/Cemu/mlc01/usr/save",
			"$CONFIG\\Cemu\\mlc01\\usr\\save",
		},
	})

	// Dolphin for Wii state
	states = append(states, &State{
		Platform: "WII",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/Wii",
		Source: []string{
			"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/Wii",
			"$SHARE/dolphin-emu/Wii",
			"$CONFIG/dolphin-emu/Wii",
			"$CONFIG/Dolphin/Wii",
			"$CONFIG\\Dolphin Emulator\\Wii",
		},
	}, &State{
		Platform: "WII",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
			"$SHARE/dolphin-emu/StateSaves",
			"$CONFIG/dolphin-emu/StateSaves",
			"$CONFIG/Dolphin/StateSaves",
			"$CONFIG\\Dolphin Emulator\\StateSaves",
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
			"$VAR/app.xemu.xemu/data/xemu/xemu/eeprom.bin",
			"$SHARE/xemu/xemu/eeprom.bin",
			"$CONFIG/xemu/xemu/eeprom.bin",
			"$CONFIG\\xemu\\xemu\\eeprom.bin",
		},
	}, &State{
		Platform: "XBOX",
		Emulator: "Xemu",
		Type:     "file",
		Path:     "$STATE/Xemu/xemu/xemu/xbox_hdd.qcow2",
		Source: []string{
			"$VAR/app.xemu.xemu/data/xemu/xemu/xbox_hdd.qcow2",
			"$SHARE/xemu/xemu/xbox_hdd.qcow2",
			"$CONFIG/xemu/xemu/xbox_hdd.qcow2",
			"$CONFIG\\xemu\\xemu\\xbox_hdd.qcow2",
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
			"$DOCUMENTS\\xenia\\content",
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
