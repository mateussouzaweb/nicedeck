package steam

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Download steam artworks for given app id
func DownloadArtworks(appId string, icon string, logo string, cover string, banner string, hero string) error {

	// If you dont see the images on Steam
	// Use the following command to get updated application ids
	// grep -i "<game-title>" ~/.local/share/Steam/steamapps/appmanifest_*.acf

	// TODO list paths in folder
	path := "~/.local/share/Steam/userdata/${USER_ID}/config/grid"
	err := cli.Command(fmt.Sprintf(
		`mkdir -p %s`,
		path,
	)).Run()

	if err != nil {
		return err
	}

	// Required format
	// Icon: ${APPID}.ico
	// Logo: ${APPID}_logo.png
	// Cover: ${APPID}p.png
	// Banner: ${APPID}.png
	// Hero: ${APPID}_hero_.png
	err = cli.Command(fmt.Sprintf(`
		[ "%s" != "" ] && wget -q -O %s/%s.ico %s
		[ "%s" != "" ] && wget -q -O %s/%s_logo.png %s
		[ "%s" != "" ] && wget -q -O %s/%sp.png %s
		[ "%s" != "" ] && wget -q -O %s/%s.png %s
		[ "%s" != "" ] && wget -q -O %s/%s_hero.png %s
		`,
		icon, path, appId, icon,
		logo, path, appId, logo,
		cover, path, appId, cover,
		banner, path, appId, banner,
		hero, path, appId, hero,
	)).Run()

	return err
}
