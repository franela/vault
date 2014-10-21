package testutils

import (
	"io/ioutil"
	"os"
	"path"
)

func GetProjectDir() string {
	return os.Getenv("PROJECTDIR")
}

func SetTestGPGHome(name string) {
	os.Setenv("GNUPGHOME", path.Join(GetProjectDir(), "testdata", name))
}

func RemoveTemporaryHomeDir(homedir string) {
	err := os.RemoveAll(homedir)
	if err != nil {
		panic(err)
	}
}

func GetTemporaryHomeDir() string {
	tempPath := path.Join(os.Getenv("PROJECTDIR"), "tmp")

	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		if err := os.MkdirAll(tempPath, 0777); err != nil {
			panic(err)
		}
	}

	name, err := ioutil.TempDir(tempPath, "vault")

	if err != nil {
		panic(err)
	}

	return name
}
