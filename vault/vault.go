package vault

import (
	"os"
)

func GetHomeDir() (string, error) {
	env := os.Getenv("VAULTDIR")
	if len(env) == 0 {
		if wd, err := os.Getwd(); err != nil {
			return "", err
		} else {
			return wd, nil
		}
	} else {
		return env, nil
	}
}
