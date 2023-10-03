package cli

import "os"

// Ensure current working directory is home
func EnsureHome() error {

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Move working directory to home folder
	err = os.Chdir(home)
	if err != nil {
		return err
	}

	return nil
}
