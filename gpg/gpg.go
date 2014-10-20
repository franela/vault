package gpg

import (
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
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

func ReEncryptFile(src, dst string, recipients []string) error {
	decryptArgs := []string{"--decrypt", "--armor", "--batch", "--yes", src}
	encryptArgs := []string{"--encrypt", "--armor", "--batch", "--yes", "--output", dst}

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	decryptCmd := exec.Command("gpg", decryptArgs...)
	encryptCmd := exec.Command("gpg", encryptArgs...)

	r, w := io.Pipe()
	decryptCmd.Stdout = w
	encryptCmd.Stdin = r
	encryptCmd.Stderr = os.Stderr
	encryptCmd.Stdout = os.Stderr

	err1 := decryptCmd.Start()
	err2 := encryptCmd.Start()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err1
	}

	err1 = decryptCmd.Wait()
	w.Close()
	err2 = encryptCmd.Wait()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err1
	}
	return nil
}
