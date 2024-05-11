package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Installer function
type Installer func(shortcut *shortcuts.Shortcut) error

// Program struct
type Program struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Category         string    `json:"category"`
	Tags             []string  `json:"tags"`
	RequiredFolders  []string  `json:"requiredFolders"`
	FlatpakAppID     string    `json:"flatpakAppId"`
	FlatpakOverrides []string  `json:"flatpakOverrides"`
	FlatpakArguments []string  `json:"FlatpakArguments"`
	IconURL          string    `json:"iconUrl"`
	LogoURL          string    `json:"logoUrl"`
	CoverURL         string    `json:"coverUrl"`
	BannerURL        string    `json:"bannerUrl"`
	HeroURL          string    `json:"heroUrl"`
	Installer        Installer `json:"-"`
}

// Retrieve list of available programs to install
func GetPrograms() []*Program {

	var programs []*Program

	programs = append(programs, BraveBrowser())
	programs = append(programs, Bottles())
	programs = append(programs, Cemu())
	programs = append(programs, Chiaki())
	programs = append(programs, Citra())
	programs = append(programs, Dolphin())
	programs = append(programs, DuckStation())
	programs = append(programs, EmulationStationDE())
	programs = append(programs, Firefox())
	programs = append(programs, Flycast())
	programs = append(programs, GeForceNow())
	programs = append(programs, GoogleChrome())
	programs = append(programs, HeroicGamesLauncher())
	programs = append(programs, JellyfinMediaPlayer())
	programs = append(programs, Lutris())
	programs = append(programs, MelonDS())
	programs = append(programs, MGBA())
	programs = append(programs, MicrosoftEdge())
	programs = append(programs, MoonlightGameStreaming())
	programs = append(programs, PCSX2())
	programs = append(programs, PPSSPP())
	programs = append(programs, ProtonPlus())
	programs = append(programs, RPCS3())
	programs = append(programs, Ryujinx())
	programs = append(programs, Simple64())
	programs = append(programs, XboxCloudGaming())
	programs = append(programs, Xemu())
	programs = append(programs, Yuzu())

	return programs
}

// Retrieve program with given ID
func GetProgramByID(id string) *Program {

	for _, program := range GetPrograms() {
		if id == program.ID {
			return program
		}
	}

	return &Program{}
}

// Install program with given ID
func Install(id string) error {

	program := GetProgramByID(id)

	// Program not found
	if program.ID == "" {
		return fmt.Errorf("program not found: %s", id)
	}

	// Print step message
	cli.Printf(cli.ColorNotice, "Installing %s...\n", program.Name)

	// Fill basic shortcut information
	shortcut := &shortcuts.Shortcut{
		AppName:   program.Name,
		Tags:      program.Tags,
		IconURL:   program.IconURL,
		LogoURL:   program.LogoURL,
		CoverURL:  program.CoverURL,
		BannerURL: program.BannerURL,
		HeroURL:   program.HeroURL,
	}

	// Make sure required folders exist
	if len(program.RequiredFolders) > 0 {
		for _, folder := range program.RequiredFolders {
			err := os.MkdirAll(os.ExpandEnv(folder), 0755)
			if err != nil {
				return err
			}
		}
	}

	// Program is flatpak
	if program.FlatpakAppID != "" {

		// Install from flatpak
		script := fmt.Sprintf(
			"flatpak install --or-update --assumeyes --noninteractive flathub %s",
			program.FlatpakAppID,
		)

		err := cli.Command(script).Run()
		if err != nil {
			return err
		}

		// Apply flatpak overrides
		if len(program.FlatpakOverrides) > 0 {
			for _, override := range program.FlatpakOverrides {
				script := fmt.Sprintf("flatpak override --user %s %s", override, program.FlatpakAppID)
				err := cli.Command(script).Run()
				if err != nil {
					return err
				}
			}
		}

		// Fill shortcut information for flatpak app
		shortcut.StartDir = "/var/lib/flatpak/exports/bin/"
		shortcut.Exe = "/var/lib/flatpak/exports/bin/" + program.FlatpakAppID
		shortcut.ShortcutPath = "/var/lib/flatpak/exports/share/applications/" + program.FlatpakAppID + ".desktop"
		shortcut.LaunchOptions = ""

		// Append shortcut launch arguments
		if len(program.FlatpakArguments) > 0 {
			shortcut.LaunchOptions = strings.Join(program.FlatpakArguments, " ")
		}

	}

	// Program has custom installer script
	if program.Installer != nil {
		err := program.Installer(shortcut)
		if err != nil {
			return err
		}
	}

	// Add to shortcuts list
	err := library.AddToShortcuts(shortcut, false)
	if err != nil {
		return err
	}

	// Print success message
	cli.Printf(cli.ColorSuccess, "%s installed!\n", program.Name)
	return nil
}
