package programs

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Retrieve list of available programs to install
func GetPrograms() ([]*packaging.Program, error) {

	var programs []*packaging.Program
	var available []*packaging.Program

	// Retrieve all possible programs
	programs = append(programs, AmazonGames())
	programs = append(programs, Azahar())
	programs = append(programs, BattleNet())
	programs = append(programs, BraveBrowser())
	programs = append(programs, Cemu())
	programs = append(programs, ChiakiNG())
	programs = append(programs, Discord())
	programs = append(programs, Dolphin())
	programs = append(programs, DuckStation())
	programs = append(programs, EAApp())
	programs = append(programs, Eden())
	programs = append(programs, EpicGames())
	programs = append(programs, ESDE())
	programs = append(programs, Firefox())
	programs = append(programs, Flycast())
	programs = append(programs, GeForceNow())
	programs = append(programs, GOGGalaxy())
	programs = append(programs, GoogleChrome())
	programs = append(programs, Gopher64())
	programs = append(programs, HeroicGamesLauncher())
	programs = append(programs, MelonDS())
	programs = append(programs, MGBA())
	programs = append(programs, MicrosoftEdge())
	programs = append(programs, MoonlightGameStreaming())
	programs = append(programs, NiceDeck())
	programs = append(programs, PCSX2())
	programs = append(programs, PPSSPP())
	programs = append(programs, ProtonPlus())
	programs = append(programs, Redream())
	programs = append(programs, RockstarGamesLauncher())
	programs = append(programs, RPCS3())
	programs = append(programs, Ryujinx())
	programs = append(programs, ShadPS4())
	programs = append(programs, Steam())
	programs = append(programs, UbisoftConnect())
	programs = append(programs, Vita3K())
	programs = append(programs, XboxCloudGaming())
	programs = append(programs, Xemu())
	programs = append(programs, Xenia())

	// Filter to return only available programs
	for _, program := range programs {
		if program.Package.Available() {
			available = append(available, program)
		}
	}

	// Flag installed programs and runtime
	for _, program := range available {
		runtime := fmt.Sprintf("--%s", program.Package.Runtime())
		program.Flags = append(program.Flags, runtime)

		installed, err := program.Package.Installed()
		if err != nil {
			return available, err
		}
		if installed {
			program.Flags = append(program.Flags, "--installed")
		}
	}

	return available, nil
}

// Retrieve program with given ID
func GetProgramByID(id string) (*packaging.Program, error) {

	programs, err := GetPrograms()
	notFound := &packaging.Program{
		Package: &packaging.Missing{},
	}

	if err != nil {
		return notFound, err
	}

	for _, program := range programs {
		if id == program.ID {
			return program, nil
		}
	}

	return notFound, nil
}
