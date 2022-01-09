package agilis

import (
	"io/ioutil"
	"strings"
)

func (db Database) minify() (err error) {
	path, err := db.DefaultPath()
	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	replacements := map[string]string{
		"{\n\t": "{",
		"\n}":   "}",
		"}\n":   "}",
		" : ":   ":",
		" = ":   "=",
	}
	for k, v := range replacements {
		data = []byte(strings.ReplaceAll(string(data), k, v))
	}

	mp, err := db.MinifiedPath()
	if err != nil {
		return
	}
	err = ioutil.WriteFile(mp, data, 0777)
	return
}
