package steam

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Download steam artworks for given app id
func DownloadArtworks(appId string, icon string, logo string, cover string, banner string, hero string) error {

	// Retrieve the user data path
	path, err := GetUserDataPath()
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
		# Make sure folder exist
		mkdir -p %s/config/grid

		# Download images
		[ "%s" != "" ] && wget -q -O %s/config/grid/%s.ico %s
		[ "%s" != "" ] && wget -q -O %s/config/grid/%s_logo.png %s
		[ "%s" != "" ] && wget -q -O %s/config/grid/%sp.png %s
		[ "%s" != "" ] && wget -q -O %s/config/grid/%s.png %s
		[ "%s" != "" ] && wget -q -O %s/config/grid/%s_hero.png %s
		`,
		path,
		icon, path, appId, icon,
		logo, path, appId, logo,
		cover, path, appId, cover,
		banner, path, appId, banner,
		hero, path, appId, hero,
	)).Run()

	return err
}
