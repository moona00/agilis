package agilis

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func getDataToWrite(data Data) (ret string) {
	for n, d := range data {
		t := reflect.TypeOf(d)
		v := reflect.ValueOf(d)
		ret += fmt.Sprintf("%s : %s = {\n", n, t)

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			ret += fmt.Sprintf("\t%s : %s = %v,\n", f.Name, f.Type, v.Field(i))
		}

		ret += "}\n"
	}

	return ret[:len(ret)-1]
}

func (db *Database) add(data Data, prepend bool) (err error) {
	path, err := db.DefaultPath()
	if err != nil {
		return
	}

	old, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	str := getDataToWrite(data)
	dataToWrite := []byte(getDataToWrite(data))
	if prepend {
		dataToWrite = append([]byte(str+"\n"), old...)
	} else {
		dataToWrite = append(old, []byte("\n"+str)...)
	}

	if err = ioutil.WriteFile(path, dataToWrite, filePerms); err != nil {
		return
	}

	if err = db.minify(); err != nil {
		return
	}

	for k, v := range data {
		db.Data[k] = v
	}

	return nil
}

// Prepend prepends stuff
func (db *Database) Prepend(data Data) error {
	return db.add(data, true)
}

// Append appends stuff
func (db *Database) Append(data Data) error {
	return db.add(data, false)
}
