package windows

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Convert PNG icon to ICO format for Windows shortcuts
func ConvertPngToIco(source string, destination string) error {

	// Using .NET's System.Drawing assembly to perform the conversion
	script := fmt.Sprintf(``+
		`Add-Type -AssemblyName System.Drawing;`+
		`$png = "%s";`+
		`$ico = "%s";`+
		`$bitmap = [System.Drawing.Bitmap]::FromFile($png);`+
		`$icon = [System.Drawing.Icon]::FromHandle($bitmap.GetHicon());`+
		`$fileStream = [System.IO.File]::Create($ico);`+
		`$icon.Save($fileStream);`+
		`$fileStream.Close();`+
		`$icon.Dispose();`+
		`$bitmap.Dispose();`,
		source,
		destination,
	)

	command := cli.Command(script)
	err := cli.Run(command)
	if err != nil {
		return err
	}

	return nil
}
