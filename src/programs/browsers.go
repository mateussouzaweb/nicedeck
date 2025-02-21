package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
)

// Installer for Brave Browser
func BraveBrowser() *Program {
	return &Program{
		ID:          "brave-browser",
		Name:        "Brave Browser",
		Description: "Web browser",
		Category:    "Utilities",
		Tags:        []string{"Utilities"},
		Folders:     []string{},
		Website:     "https://brave.com",
		IconURL:     assets.Icon("192d80a88b27b3e4115e1a45a782fe1b.png"),
		LogoURL:     assets.Logo("85b79607444cc565f0214d12c05cc5eb.png"),
		CoverURL:    assets.Cover("01a0ed0f07ddea7687fefaedb0f32a7b.png"),
		BannerURL:   assets.Banner("5ac7b3d023885d0d49e05a32f16c3d54.png"),
		HeroURL:     assets.Hero("776c2a30d4402b8c5126edd7ad111c5e.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.brave.Browser",
		}, &macos.Homebrew{
			AppID:   "brave-browser",
			AppName: "Brave Browser.app",
		}, &windows.WinGet{
			AppID:  "Brave.Brave",
			AppExe: "$APPDATA\\Local\\BraveSoftware\\Brave-Browser\\Application\\brave.exe",
		}),
	}
}

// Installer for Firefox
func Firefox() *Program {
	return &Program{
		ID:          "firefox",
		Name:        "Firefox",
		Description: "Web browser",
		Category:    "Utilities",
		Tags:        []string{"Utilities"},
		Website:     "https://www.mozilla.org/en-US/firefox",
		IconURL:     assets.Icon("b59bb2585ce93f60e21e1fab71cbf4ad.png"),
		LogoURL:     assets.Logo("43285a8b542fcdc35377439e05dcb04f.png"),
		CoverURL:    assets.Cover("4529f985441a035ae4a107b8862ba4dd.png"),
		BannerURL:   assets.Banner("9384fe92aef7ea0128be2c916ed07cea.png"),
		HeroURL:     assets.Hero("a318166b8539611449bf21ddc297a783.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "org.mozilla.firefox",
		}, &macos.Homebrew{
			AppID:   "firefox",
			AppName: "Firefox.app",
		}, &windows.WinGet{
			AppID:  "Mozilla.Firefox",
			AppExe: "$PROGRAMS\\Mozilla Firefox\\firefox.exe",
		}),
	}
}

// Installer for Google Chrome
func GoogleChrome() *Program {
	return &Program{
		ID:          "google-chrome",
		Name:        "Google Chrome",
		Description: "Web browser",
		Category:    "Utilities",
		Tags:        []string{"Utilities"},
		Website:     "https://www.google.com/intl/en_us/chrome",
		IconURL:     assets.Icon("09b45ae46393da4adbd9b0bdb977d1aa.png"),
		LogoURL:     assets.Logo("3b049d0f6cbf5421d399f156807b8657.png"),
		CoverURL:    assets.Cover("d45c26607db83f6f14b09dd70123913b.png"),
		BannerURL:   assets.Banner("d40c243072a2d2957b3484e775f1f925.png"),
		HeroURL:     assets.Hero("cae83cfcb1d8a2a4bb17bd1446fb1cee.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.google.Chrome",
		}, &macos.Homebrew{
			AppID:   "google-chrome",
			AppName: "Google Chrome.app",
		}, &windows.WinGet{
			AppID:  "Google.Chrome",
			AppExe: "$PROGRAMS\\Google\\Chrome\\Application\\chrome.exe",
		}),
	}
}

// Installer for Microsoft Edge
func MicrosoftEdge() *Program {
	return &Program{
		ID:          "microsoft-edge",
		Name:        "Microsoft Edge",
		Description: "Web browser",
		Category:    "Utilities",
		Tags:        []string{"Utilities"},
		Website:     "https://www.microsoft.com/en-us/edge",
		IconURL:     assets.Icon("714cb7478d98b1cb51d1f5f515f060c7.png"),
		LogoURL:     assets.Logo("cb88c85733fd8241b9190750318f1e59.png"),
		CoverURL:    assets.Cover("ca0dadd4ae381d26d4771208c1aa4408.png"),
		BannerURL:   assets.Banner("0656137651272c4bc984747f7a3e8c2d.png"),
		HeroURL:     assets.Hero("2c81a094d632c8b510c6c676eec4c358.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.microsoft.Edge",
		}, &macos.Homebrew{
			AppID:   "microsoft-edge",
			AppName: "Microsoft Edge.app",
		}, &windows.WinGet{
			AppID:  "Microsoft.Edge",
			AppExe: "$PROGRAMS_X86\\Microsoft\\Edge\\Application\\msedge.exe",
		}),
	}
}
