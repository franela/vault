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
	name, err := ioutil.TempDir(path.Join(os.Getenv("PROJECTDIR"), "tmp"), "vault")
	if err != nil {
		panic(err)
	}

	return name
}
