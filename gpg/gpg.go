package gpg

import (
	"os"
	"os/exec"
	"path"
	"strings"
  "code.google.com/p/go.crypto/openpgp"
)

func Decrypt(filePath string) (string, error) {
	decryptArgs := []string{"--decrypt", "--armor", "--batch", "--yes", filePath}

	cmd := exec.Command("gpg", decryptArgs...)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func DecryptFile(outputFile, filePath string) error {

	decryptArgs := []string{"--decrypt", "--armor", "--batch", "--yes", "--output", outputFile, filePath}

	cmd := exec.Command("gpg", decryptArgs...)

	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func Encrypt(filePath string, text string, recipients []string) error {
	if err := os.MkdirAll(path.Dir(filePath), 0777); err != nil {
		return err
	}

	encryptArgs := []string{"--encrypt", "--armor", "--batch", "--yes", "--output", filePath}

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	cmd := exec.Command("gpg", encryptArgs...)
	cmd.Stdin = strings.NewReader(text)
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

func EncryptFile(filePath string, sourceFile string, recipients []string) error {

	if err := os.MkdirAll(path.Dir(filePath), 0777); err != nil {
		return err
	}

	encryptArgs := []string{"--encrypt", "--armor", "--batch", "--yes", "--output", filePath}

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	encryptArgs = append(encryptArgs, sourceFile)

	cmd := exec.Command("gpg", encryptArgs...)
	_, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}
	return nil
}

func GetRecipientsFromEncryptedFile(filePath string) []string {
    // Open the private key file

    

    keyringFileBuffer2, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer keyringFileBuffer2.Close()

    openpgp.ReadKeyRing()


  return []string {}
}




