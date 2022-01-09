package agilis

import (
	"io/ioutil"
	"os"
)

// New creates a new database with the given name and data and returns it
func New(name string, data map[string]interface{}) (db *Database, err error) {
	db = &Database{
		Name: name,
		Data: data,
	}
	path, err := db.DefaultPath()
	if err != nil {
		return
	}

	if fileExists(path) {
		err = os.ErrExist
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	if os.Chdir(dir) != nil {
		return
	}

	fp, err := db.FolderPath()
	if err != nil {
		return
	}

	if os.Mkdir(fp, filePerms) != nil ||
		os.Chmod(fp, filePerms) != nil ||
		ioutil.WriteFile(path, []byte(getDataToWrite(db.Data)), filePerms) != nil {

		return
	}

	err = db.minify()
	return
}
