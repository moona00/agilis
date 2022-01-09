package agilis

import "io/ioutil"

// Read returns the database with the given name
func Read(name string) (*Database, error) {
	db := new(Database)
	db.Name = name

	path, err := db.MinifiedPath()
	if err != nil {
		return nil, err
	}

	dataStr, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	db.Data = parse(string(dataStr))
	return db, nil
}
