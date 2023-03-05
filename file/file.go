package file

import (
	"os"
	"path"
)

const (
	baseDir = "./.file"
	ext     = ".json"
)

func Create(dir, file string, data []byte, callback func(err error)) {
	name := path.Join(baseDir, dir, file) + ext
	callback(os.WriteFile(name, data, 0666))
}

func Read(dir, file string, callback func(data []byte, err error)) {
	name := path.Join(baseDir, dir, file) + ext
	callback(os.ReadFile(name))
}

func Update(dir, file string, data []byte, callback func(err error)) {
	name := path.Join(baseDir, dir, file) + ext
	f, err := os.OpenFile(name, os.O_TRUNC|os.O_WRONLY, 0666)
	if err == nil {
		_, err = f.Write(data)
		if err == nil {
			callback(f.Close())
		} else {
			callback(err)
		}
	} else {
		callback(err)
	}
}

func Delete(dir, file string, callback func(err error)) {
	name := path.Join(baseDir, dir, file) + ext
	callback(os.Remove(name))
}
