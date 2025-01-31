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
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Lime3DS",
		Type:     "folder",
		Path:     "$STATE/Lime3DS/lime3ds-emu/states",
		Source: []string{
			"$HOME/.var/app/io.github.lime3ds.Lime3DS/data/lime3ds-emu/states",
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
		},
	}, &State{
		Platform: "3DS",
		Emulator: "Citra",
		Type:     "folder",
		Path:     "$STATE/Citra/citra-emu/states",
		Source: []string{
			"$HOME/.var/app/org.citra_emu.citra/data/citra-emu/states",
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
		},
	}, &State{
		Platform: "GBA",
		Emulator: "MGBA",
		Type:     "folder",
		Path:     "$STATE/MGBA/states",
		Source: []string{
			"$HOME/.var/app/io.mgba.mGBA/states",
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
		},
	}, &State{
		Platform: "GC",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
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
		},
	}, &State{
		Platform: "NDS",
		Emulator: "MelonDS",
		Type:     "folder",
		Path:     "$STATE/MelonDS/states",
		Source: []string{
			"$HOME/.var/app/net.kuribo64.melonDS/states",
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
		},
	}, &State{
		Platform: "PS1",
		Emulator: "DuckStation",
		Type:     "folder",
		Path:     "$STATE/DuckStation/duckstation/savestates",
		Source: []string{
			"$HOME/.var/app/org.duckstation.DuckStation/config/duckstation/savestates",
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
		},
	}, &State{
		Platform: "PS2",
		Emulator: "PCSX2",
		Type:     "folder",
		Path:     "$STATE/PCSX2/sstates",
		Source: []string{
			"$HOME/.var/app/net.pcsx2.PCSX2/config/PCSX2/sstates",
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
		},
	}, &State{
		Platform: "PSP",
		Emulator: "PPSSPP",
		Type:     "folder",
		Path:     "$STATE/PPSSPP/ppsspp/PSP/PPSSPP_STATE",
		Source: []string{
			"$HOME/.var/app/org.ppsspp.PPSSPP/config/ppsspp/PSP/PPSSPP_STATE",
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
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "folder",
		Path:     "$STATE/Ryujinx/bis/user/saveMeta",
		Source: []string{
			"$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/saveMeta",
		},
	}, &State{
		Platform: "SWITCH",
		Emulator: "Ryujinx",
		Type:     "file",
		Path:     "$STATE/Ryujinx/system/Profiles.json",
		Source: []string{
			"$HOME/.var/app/org.ryujinx.Ryujinx/config/Ryujinx/system/Profiles.json",
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
		},
	}, &State{
		Platform: "WII",
		Emulator: "Dolphin",
		Type:     "folder",
		Path:     "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: []string{
			"$HOME/.var/app/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
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
		},
	}, &State{
		Platform: "XBOX",
		Emulator: "Xemu",
		Type:     "file",
		Path:     "$STATE/Xemu/xemu/xemu/xbox_hdd.qcow2",
		Source: []string{
			"$HOME/.var/app/app.xemu.xemu/data/xemu/xemu/xbox_hdd.qcow2",
		},
	})

	// Xenia
	// Save file is expected to be in the Documents folder
	// However, user can also put this file in Xenia/ folder
	states = append(states, &State{
		Platform: "XBOX360",
		Emulator: "Xenia",
		Type:     "file",
		Path:     "$STATE/Xenia/content",
		Source: []string{
			"$USERHOME\\Documents\\xenia\\content",
			"$EMULATORS\\Xenia\\content",
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
