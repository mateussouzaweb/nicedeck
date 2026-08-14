package state

import "github.com/mateussouzaweb/nicedeck/src/fs"

// State struct
type State struct {
	Platform    string  `json:"platform"`
	Emulator    string  `json:"emulator"`
	Type        string  `json:"type"`
	Destination string  `json:"destination"`
	Source      *Source `json:"source"`
}

// Custom state struct
type CustomState struct {
	Platform    string `json:"platform"`
	Emulator    string `json:"emulator"`
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

// Retrieve save state of each platform
func GetStates(options *Options) ([]*State, error) {

	// The following emulators store saves and states on ROMs directory:
	// - MGBA (user can leave at it is or configure emulator)
	// - MelonDS (user can leave at it is or configure emulator)
	states := []*State{}

	// Azahar
	states = append(states, &State{
		Platform:    "3DS",
		Emulator:    "Azahar",
		Type:        "folder",
		Destination: "$STATE/Azahar/azahar-emu/sdmc",
		Source: &Source{
			Linux:   []string{"$SHARE/azahar-emu/sdmc"},
			MacOS:   []string{"$CONFIG/azahar-emu/sdmc"},
			Windows: []string{"$CONFIG/Azahar/sdmc"},
		},
	}, &State{
		Platform:    "3DS",
		Emulator:    "Azahar",
		Type:        "folder",
		Destination: "$STATE/Azahar/azahar-emu/states",
		Source: &Source{
			Linux:   []string{"$SHARE/azahar-emu/states"},
			MacOS:   []string{"$CONFIG/azahar-emu/states"},
			Windows: []string{"$CONFIG/Azahar/states"},
		},
	})

	// Flycast
	states = append(states, &State{
		Platform:    "DC",
		Emulator:    "Flycast",
		Type:        "folder",
		Destination: "$STATE/Flycast/flycast",
		Source: &Source{
			Linux: []string{
				"$VAR/org.flycast.Flycast/data/flycast",
				"$SHARE/flycast",
			},
			MacOS: []string{"$CONFIG/flycast/data"},
			Windows: []string{
				"$CONFIG/Flycast/data",
				"$EMULATORS/Flycast/data",
			},
		},
	})

	// Redream
	states = append(states, &State{
		Platform:    "DC",
		Emulator:    "Redream",
		Type:        "folder",
		Destination: "$STATE/Redream/saves",
		Source: &Source{
			Linux:   []string{"$SHARE/Redream/saves"},
			MacOS:   []string{"$CONFIG/Redream/saves"},
			Windows: []string{"$EMULATORS/Redream/saves"},
		},
	})

	// MGBA requires configuration to work:
	// - Go to Tools > Settings > Paths
	// - Set save games location as $VAR/io.mgba.mGBA/save
	// - Set save states as $VAR/io.mgba.mGBA/states
	// - Make sure that "same directory as the ROM" in both options is unchecked
	states = append(states, &State{
		Platform:    "GBA",
		Emulator:    "MGBA",
		Type:        "folder",
		Destination: "$STATE/MGBA/save",
		Source: &Source{
			Linux: []string{
				"$VAR/io.mgba.mGBA/save",
				"$SHARE/mGBA/save",
			},
			MacOS:   []string{"$CONFIG/mGBA/save"},
			Windows: []string{"$EMULATORS/MGBA/save"},
		},
	}, &State{
		Platform:    "GBA",
		Emulator:    "MGBA",
		Type:        "folder",
		Destination: "$STATE/MGBA/states",
		Source: &Source{
			Linux: []string{
				"$VAR/io.mgba.mGBA/states",
				"$SHARE/mGBA/states",
			},
			MacOS:   []string{"$CONFIG/mGBA/states"},
			Windows: []string{"$EMULATORS/MGBA/states"},
		},
	})

	// Dolphin for GameCube state
	states = append(states, &State{
		Platform:    "GC",
		Emulator:    "Dolphin",
		Type:        "folder",
		Destination: "$STATE/Dolphin/dolphin-emu/GC",
		Source: &Source{
			Linux: []string{
				"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/GC",
				"$SHARE/dolphin-emu/GC",
				"$CONFIG/dolphin-emu/GC",
			},
			MacOS:   []string{"$CONFIG/Dolphin/GC"},
			Windows: []string{"$CONFIG/Dolphin Emulator/GC"},
		},
	}, &State{
		Platform:    "GC",
		Emulator:    "Dolphin",
		Type:        "folder",
		Destination: "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: &Source{
			Linux: []string{
				"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
				"$SHARE/dolphin-emu/StateSaves",
				"$CONFIG/dolphin-emu/StateSaves",
			},
			MacOS:   []string{"$CONFIG/Dolphin/StateSaves"},
			Windows: []string{"$CONFIG/Dolphin Emulator/StateSaves"},
		},
	})

	// MelonDS requires configuration to work:
	// - Go to Config > Path Settings
	// - Set save files path as $VAR/net.kuribo64.melonDS/save
	// - Set save states path as $VAR/net.kuribo64.melonDS/states
	states = append(states, &State{
		Platform:    "NDS",
		Emulator:    "MelonDS",
		Type:        "folder",
		Destination: "$STATE/MelonDS/save",
		Source: &Source{
			Linux: []string{
				"$VAR/net.kuribo64.melonDS/save",
				"$SHARE/melonDS/save",
			},
			MacOS:   []string{"$CONFIG/melonDS/save"},
			Windows: []string{"$EMULATORS/MelonDS/save"},
		},
	}, &State{
		Platform:    "NDS",
		Emulator:    "MelonDS",
		Type:        "folder",
		Destination: "$STATE/MelonDS/states",
		Source: &Source{
			Linux: []string{
				"$VAR/net.kuribo64.melonDS/states",
				"$SHARE/melonDS/states",
			},
			MacOS:   []string{"$CONFIG/melonDS/states"},
			Windows: []string{"$EMULATORS/MelonDS/states"},
		},
	})

	// Gopher64
	states = append(states, &State{
		Platform:    "N64",
		Emulator:    "Gopher64",
		Type:        "folder",
		Destination: "$STATE/Gopher64/saves",
		Source: &Source{
			Linux:   []string{"$VAR/io.github.gopher64.gopher64/data/gopher64/saves"},
			MacOS:   []string{"$HOME/Library/Containers/io.github.gopher64.gopher64/Data/Library/Application Support/gopher64/saves"},
			Windows: []string{"$APPDATA/Roaming/gopher64/saves"},
		},
	}, &State{
		Platform:    "N64",
		Emulator:    "Gopher64",
		Type:        "folder",
		Destination: "$STATE/Gopher64/states",
		Source: &Source{
			Linux:   []string{"$VAR/io.github.gopher64.gopher64/data/gopher64/states"},
			MacOS:   []string{"$HOME/Library/Containers/io.github.gopher64.gopher64/Data/Library/Application Support/gopher64/states"},
			Windows: []string{"$APPDATA/Roaming/gopher64/states"},
		},
	})

	// DuckStation
	states = append(states, &State{
		Platform:    "PS1",
		Emulator:    "DuckStation",
		Type:        "folder",
		Destination: "$STATE/DuckStation/duckstation/memcards",
		Source: &Source{
			Linux: []string{
				"$VAR/org.duckstation.DuckStation/config/duckstation/memcards",
				"$SHARE/duckstation/memcards",
				"$CONFIG/duckstation/memcards",
			},
			MacOS:   []string{"$CONFIG/DuckStation/memcards"},
			Windows: []string{"$DOCUMENTS/DuckStation/memcards"},
		},
	}, &State{
		Platform:    "PS1",
		Emulator:    "DuckStation",
		Type:        "folder",
		Destination: "$STATE/DuckStation/duckstation/savestates",
		Source: &Source{
			Linux: []string{
				"$VAR/org.duckstation.DuckStation/config/duckstation/savestates",
				"$SHARE/duckstation/savestates",
				"$CONFIG/duckstation/savestates",
			},
			MacOS:   []string{"$CONFIG/DuckStation/savestates"},
			Windows: []string{"$DOCUMENTS/DuckStation/savestates"},
		},
	})

	// PCSX2
	states = append(states, &State{
		Platform:    "PS2",
		Emulator:    "PCSX2",
		Type:        "folder",
		Destination: "$STATE/PCSX2/memcards",
		Source: &Source{
			Linux: []string{
				"$VAR/net.pcsx2.PCSX2/config/PCSX2/memcards",
				"$SHARE/PCSX2/memcards",
			},
			MacOS:   []string{"$CONFIG/PCSX2/memcards"},
			Windows: []string{"$DOCUMENTS/PCSX2/memcards"},
		},
	}, &State{
		Platform:    "PS2",
		Emulator:    "PCSX2",
		Type:        "folder",
		Destination: "$STATE/PCSX2/sstates",
		Source: &Source{
			Linux: []string{
				"$VAR/net.pcsx2.PCSX2/config/PCSX2/sstates",
				"$SHARE/PCSX2/sstates",
			},
			MacOS:   []string{"$CONFIG/PCSX2/sstates"},
			Windows: []string{"$DOCUMENTS/PCSX2/sstates"},
		},
	})

	// RPCS3
	states = append(states, &State{
		Platform:    "PS3",
		Emulator:    "RPCS3",
		Type:        "folder",
		Destination: "$STATE/RPCS3/rpcs3/dev_hdd0/home/00000001/savedata",
		Source: &Source{
			Linux: []string{
				"$VAR/net.rpcs3.RPCS3/config/rpcs3/dev_hdd0/home/00000001/savedata",
				"$SHARE/rpcs3/dev_hdd0/home/00000001/savedata",
			},
			MacOS:   []string{"$CONFIG/rpcs3/dev_hdd0/home/00000001/savedata"},
			Windows: []string{"$EMULATORS/RPCS3/dev_hdd0/home/00000001/savedata"},
		},
	})

	// ShadPS4
	states = append(states, &State{
		Platform:    "PS4",
		Emulator:    "ShadPS4",
		Type:        "folder",
		Destination: "$STATE/ShadPS4/saves",
		Source: &Source{
			Linux: []string{
				"$VAR/net.shadps4.shadPS4/data/shadps4/saves",
				"$SHARE/shadps4/saves",
			},
			MacOS:   []string{"$CONFIG/shadps4/saves"},
			Windows: []string{"$EMULATORS/ShadPS4/user/saves"},
		},
	})

	// PPSSPP
	states = append(states, &State{
		Platform:    "PSP",
		Emulator:    "PPSSPP",
		Type:        "folder",
		Destination: "$STATE/PPSSPP/ppsspp/PSP/SAVEDATA",
		Source: &Source{
			Linux: []string{
				"$VAR/org.ppsspp.PPSSPP/config/ppsspp/PSP/SAVEDATA",
				"$SHARE/ppsspp/PSP/SAVEDATA",
			},
			MacOS:   []string{"$CONFIG/ppsspp/PSP/SAVEDATA"},
			Windows: []string{"$EMULATORS/PPSSPP/memstick/PSP/SAVEDATA"},
		},
	}, &State{
		Platform:    "PSP",
		Emulator:    "PPSSPP",
		Type:        "folder",
		Destination: "$STATE/PPSSPP/ppsspp/PSP/PPSSPP_STATE",
		Source: &Source{
			Linux: []string{
				"$VAR/org.ppsspp.PPSSPP/config/ppsspp/PSP/PPSSPP_STATE",
				"$SHARE/ppsspp/PSP/PPSSPP_STATE",
			},
			MacOS:   []string{"$CONFIG/ppsspp/PSP/PPSSPP_STATE"},
			Windows: []string{"$EMULATORS/PPSSPP/memstick/PSP/PPSSPP_STATE"},
		},
	})

	// Vita3K
	states = append(states, &State{
		Platform:    "PSVITA",
		Emulator:    "Vita3K",
		Type:        "folder",
		Destination: "$STATE/Vita3K/ux0/user/00/savedata",
		Source: &Source{
			Linux:   []string{"$SHARE/Vita3K/Vita3K/ux0/user/00/savedata"},
			MacOS:   []string{"$CONFIG/Vita3K/Vita3K/ux0/user/00/savedata"},
			Windows: []string{"$CONFIG/Vita3K/Vita3K/fs/ux0/user/00/savedata"},
		},
	})

	// Eden
	states = append(states, &State{
		Platform:    "SWITCH",
		Emulator:    "Eden",
		Type:        "folder",
		Destination: "$STATE/Eden/nand/user/save",
		Source: &Source{
			Linux: []string{"$SHARE/eden/nand/user/save"},
			MacOS: []string{"$CONFIG/eden/nand/user/save"},
			Windows: []string{
				"$CONFIG/eden/nand/user/save",
				"$EMULATORS/Eden/user/nand/user/save",
			},
		},
	})

	// Ryujinx save state in two folders
	// We also sync the profiles.json file to avoid losing user reference
	states = append(states, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Destination: "$STATE/Ryujinx/bis/user/save",
		Source: &Source{
			Linux: []string{
				"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/save",
				"$SHARE/Ryujinx/bis/user/save",
			},
			MacOS:   []string{"$CONFIG/Ryujinx/bis/user/save"},
			Windows: []string{"$CONFIG/Ryujinx/bis/user/save"},
		},
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "folder",
		Destination: "$STATE/Ryujinx/bis/user/saveMeta",
		Source: &Source{
			Linux: []string{
				"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/bis/user/saveMeta",
				"$SHARE/Ryujinx/bis/user/saveMeta",
			},
			MacOS:   []string{"$CONFIG/Ryujinx/bis/user/saveMeta"},
			Windows: []string{"$CONFIG/Ryujinx/bis/user/saveMeta"},
		},
	}, &State{
		Platform:    "SWITCH",
		Emulator:    "Ryujinx",
		Type:        "file",
		Destination: "$STATE/Ryujinx/system/Profiles.json",
		Source: &Source{
			Linux: []string{
				"$VAR/org.ryujinx.Ryujinx/config/Ryujinx/system/Profiles.json",
				"$SHARE/Ryujinx/system/Profiles.json",
			},
			MacOS:   []string{"$CONFIG/Ryujinx/system/Profiles.json"},
			Windows: []string{"$CONFIG/Ryujinx/system/Profiles.json"},
		},
	})

	// Cemu
	states = append(states, &State{
		Platform:    "WIIU",
		Emulator:    "Cemu",
		Type:        "folder",
		Destination: "$STATE/Cemu/mlc01/usr/save",
		Source: &Source{
			Linux: []string{
				"$VAR/info.cemu.Cemu/data/Cemu/mlc01/usr/save",
				"$SHARE/Cemu/mlc01/usr/save",
			},
			MacOS:   []string{"$CONFIG/Cemu/mlc01/usr/save"},
			Windows: []string{"$CONFIG/Cemu/mlc01/usr/save"},
		},
	})

	// Dolphin for Wii state
	states = append(states, &State{
		Platform:    "WII",
		Emulator:    "Dolphin",
		Type:        "folder",
		Destination: "$STATE/Dolphin/dolphin-emu/Wii",
		Source: &Source{
			Linux: []string{
				"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/Wii",
				"$SHARE/dolphin-emu/Wii",
				"$CONFIG/dolphin-emu/Wii",
			},
			MacOS:   []string{"$CONFIG/Dolphin/Wii"},
			Windows: []string{"$CONFIG/Dolphin Emulator/Wii"},
		},
	}, &State{
		Platform:    "WII",
		Emulator:    "Dolphin",
		Type:        "folder",
		Destination: "$STATE/Dolphin/dolphin-emu/StateSaves",
		Source: &Source{
			Linux: []string{
				"$VAR/org.DolphinEmu.dolphin-emu/data/dolphin-emu/StateSaves",
				"$SHARE/dolphin-emu/StateSaves",
				"$CONFIG/dolphin-emu/StateSaves",
			},
			MacOS:   []string{"$CONFIG/Dolphin/StateSaves"},
			Windows: []string{"$CONFIG/Dolphin Emulator/StateSaves"},
		},
	})

	// Xemu requires two files to transfer states
	// The HD file is expected to be in the $BIOS/XBOX/xbox_hdd.qcow2 location by default
	// However, we also put it here in case of the user put this file in xemu/ folder
	states = append(states, &State{
		Platform:    "XBOX",
		Emulator:    "Xemu",
		Type:        "file",
		Destination: "$STATE/Xemu/xemu/xemu/eeprom.bin",
		Source: &Source{
			Linux: []string{
				"$VAR/app.xemu.xemu/data/xemu/xemu/eeprom.bin",
				"$SHARE/xemu/xemu/eeprom.bin",
			},
			MacOS:   []string{"$CONFIG/xemu/xemu/eeprom.bin"},
			Windows: []string{"$CONFIG/xemu/xemu/eeprom.bin"},
		},
	}, &State{
		Platform:    "XBOX",
		Emulator:    "Xemu",
		Type:        "file",
		Destination: "$STATE/Xemu/xemu/xemu/xbox_hdd.qcow2",
		Source: &Source{
			Linux: []string{
				"$VAR/app.xemu.xemu/data/xemu/xemu/xbox_hdd.qcow2",
				"$SHARE/xemu/xemu/xbox_hdd.qcow2",
			},
			MacOS:   []string{"$CONFIG/xemu/xemu/xbox_hdd.qcow2"},
			Windows: []string{"$CONFIG/xemu/xemu/xbox_hdd.qcow2"},
		},
	})

	// Xenia
	// Save file is expected to be in the Documents folder
	// However, user can also put this file in Xenia/ folder
	states = append(states, &State{
		Platform:    "X360",
		Emulator:    "Xenia",
		Type:        "file",
		Destination: "$STATE/Xenia/content",
		Source: &Source{
			Linux: []string{"$EMULATORS/Xenia/content"},
			Windows: []string{
				"$EMULATORS/Xenia/content",
				"$DOCUMENTS/xenia/content",
			},
		},
	})

	// Read custom states from configuration file
	customFile := fs.ExpandPath("$APPLICATIONS/NiceDeck/custom/states.json")
	customStates := make([]CustomState, 0)
	err := fs.ReadJSON(customFile, &customStates)
	if err != nil {
		return states, err
	}

	// Merge custom states with the built-in states
	for _, customState := range customStates {
		state := &State{
			Platform:    customState.Platform,
			Emulator:    customState.Emulator,
			Type:        customState.Type,
			Destination: customState.Destination,
			Source: &Source{
				Linux:   []string{customState.Source},
				MacOS:   []string{customState.Source},
				Windows: []string{customState.Source},
			},
		}

		states = append(states, state)
	}

	return states, nil
}
