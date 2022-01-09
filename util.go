package agilis

import (
	"fmt"
	"os"
)

const filePerms = 0777

// FolderPath returns the folder path of a database
func (db Database) FolderPath() (p string, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	p = fmt.Sprintf("%s/agilis/%s/", dir, db.Name)
	return
}

// DefaultPath returns the path of a database (not minified)
func (db Database) DefaultPath() (p string, err error) {
	dir, err := db.FolderPath()
	if err != nil {
		return
	}

	p = dir + db.Name
	return
}

// MinifiedPath returns the path of a database (minified version)
func (db Database) MinifiedPath() (p string, err error) {
	dir, err := db.DefaultPath()
	if err != nil {
		return
	}

	p = dir + "__m"
	return
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
