package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Installer for Brave Browser
func BraveBrowser() *Program {
	return &Program{
		ID:          "brave-browser",
		Name:        "Brave Browser",
		Description: "Web browser",
		Category:    "Utilities",
		Tags:        []string{"Utilities"},
		Website:     "https://brave.com",
		IconURL:     "https://cdn2.steamgriddb.com/icon/192d80a88b27b3e4115e1a45a782fe1b.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/85b79607444cc565f0214d12c05cc5eb.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/01a0ed0f07ddea7687fefaedb0f32a7b.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/5ac7b3d023885d0d49e05a32f16c3d54.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/776c2a30d4402b8c5126edd7ad111c5e.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.brave.Browser",
		}, &packaging.Homebrew{
			AppID:   "brave-browser",
			AppName: "Brave Browser.app",
		}, &packaging.WinGet{
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
		IconURL:     "https://cdn2.steamgriddb.com/icon/f968fdc88852a4a3a27a81fe3f57bfc5.ico",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/43285a8b542fcdc35377439e05dcb04f.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/4529f985441a035ae4a107b8862ba4dd.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/9384fe92aef7ea0128be2c916ed07cea.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/a318166b8539611449bf21ddc297a783.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "org.mozilla.firefox",
		}, &packaging.Homebrew{
			AppID:   "firefox",
			AppName: "Firefox.app",
		}, &packaging.WinGet{
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
		IconURL:     "https://cdn2.steamgriddb.com/icon/3941c4358616274ac2436eacf67fae05.ico",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/3b049d0f6cbf5421d399f156807b8657.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/d45c26607db83f6f14b09dd70123913b.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/d40c243072a2d2957b3484e775f1f925.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/cae83cfcb1d8a2a4bb17bd1446fb1cee.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.google.Chrome",
		}, &packaging.Homebrew{
			AppID:   "google-chrome",
			AppName: "Google Chrome.app",
		}, &packaging.WinGet{
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
		IconURL:     "https://cdn2.steamgriddb.com/icon/714cb7478d98b1cb51d1f5f515f060c7.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/cb88c85733fd8241b9190750318f1e59.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/ca0dadd4ae381d26d4771208c1aa4408.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/0656137651272c4bc984747f7a3e8c2d.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/2c81a094d632c8b510c6c676eec4c358.png",
		Package: packaging.Best(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.microsoft.Edge",
		}, &packaging.Homebrew{
			AppID:   "microsoft-edge",
			AppName: "Microsoft Edge.app",
		}, &packaging.WinGet{
			AppID:  "Microsoft.Edge",
			AppExe: "$PROGRAMS_X86\\Microsoft\\Edge\\Application\\msedge.exe",
		}),
	}
}
