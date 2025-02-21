package assets

import "fmt"

// Return link for icon image
func Icon(id string) string {
	return fmt.Sprintf("https://cdn2.steamgriddb.com/icon/%s", id)
}

// Return link for logo image
func Logo(id string) string {
	return fmt.Sprintf("https://cdn2.steamgriddb.com/logo/%s", id)
}

// Return link for cover image
func Cover(id string) string {
	return fmt.Sprintf("https://cdn2.steamgriddb.com/grid/%s", id)
}

// Return link for banner image
func Banner(id string) string {
	return fmt.Sprintf("https://cdn2.steamgriddb.com/grid/%s", id)
}

// Return link for hero image
func Hero(id string) string {
	return fmt.Sprintf("https://cdn2.steamgriddb.com/hero/%s", id)
}
