package windows

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Run a program with given set of arguments
func RunProcess(executable string, args []string) error {
	if len(args) > 0 {
		return cli.Start(fmt.Sprintf(``+
			`$Arguments = '%s';`+
			`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait -ArgumentList $Arguments`,
			strings.Join(args, " "),
			filepath.Dir(executable),
			executable,
		))
	}

	return cli.Start(fmt.Sprintf(
		`Start-Process -WorkingDirectory "%s" -FilePath "%s" -PassThru -Wait`,
		filepath.Dir(executable),
		executable,
	))
}
