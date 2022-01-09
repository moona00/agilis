package agilis

import "os"

// Delete deletes a database
func (db Database) Delete() error {
	path, err := db.FolderPath()
	if err != nil {
		return err
	}

	return os.RemoveAll(path)
}
