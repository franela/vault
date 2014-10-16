package vault

import (
	"os"
)

var definedHomeDir string
var workingDir string

func init() {
	var err error
	workingDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func SetHomeDir(homedir string) {
	definedHomeDir = homedir
}

func UnsetHomeDir() string {
	dir := GetHomeDir()
	definedHomeDir = ""
	return dir
}

func GetHomeDir() string {
	if len(definedHomeDir) > 0 {
		return definedHomeDir
	}
	if vaultdir := os.Getenv("VAULTDIR"); len(vaultdir) > 0 {
		return vaultdir
	}
	return workingDir
}


func GetGPGHomeDir() string {
  return os.Getenv("")

}


