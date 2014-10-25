package gpg

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func getGPGHomeDir() []string {
	if len(os.Getenv("GNUPGHOME")) > 0 {
		return []string{"--homedir", os.Getenv("GNUPGHOME")}
	}
	return []string{}
}

func Decrypt(filePath string) (string, error) {
	decryptArgs := append(getGPGHomeDir(), "--decrypt", "--armor", "--batch", "--yes", filePath)

	log.Printf("Running: gpg %s\n", strings.Join(decryptArgs, " "))
	cmd := exec.Command("gpg", decryptArgs...)
	cmd.Env = nil
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func DecryptFile(outputFile, filePath string) error {

	decryptArgs := append(getGPGHomeDir(), "--decrypt", "--armor", "--batch", "--yes", "--output", outputFile, filePath)

	log.Printf("Running: gpg %s\n", strings.Join(decryptArgs, " "))
	cmd := exec.Command("gpg", decryptArgs...)
	cmd.Env = nil
	cmd.Stderr = os.Stderr

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

	encryptArgs := append(getGPGHomeDir(), "--encrypt", "--armor", "--batch", "--yes", "--output", filePath)

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	log.Printf("Running: gpg %s\n", strings.Join(encryptArgs, " "))
	cmd := exec.Command("gpg", encryptArgs...)
	cmd.Env = nil
	cmd.Stderr = os.Stderr
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

	encryptArgs := append(getGPGHomeDir(), "--encrypt", "--armor", "--batch", "--yes", "--output", filePath)

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	encryptArgs = append(encryptArgs, sourceFile)

	log.Printf("Running: gpg %s\n", strings.Join(encryptArgs, " "))
	cmd := exec.Command("gpg", encryptArgs...)
	cmd.Env = nil
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

func ReEncryptFile(src, dst string, recipients []string) error {
	decryptArgs := append(getGPGHomeDir(), "--decrypt", "--armor", "--batch", "--yes", src)
	encryptArgs := append(getGPGHomeDir(), "--encrypt", "--armor", "--batch", "--yes", "--output", dst)

	for _, recipient := range recipients {
		encryptArgs = append(encryptArgs, "--recipient")
		encryptArgs = append(encryptArgs, recipient)
	}

	decryptCmd := exec.Command("gpg", decryptArgs...)
	decryptCmd.Env = nil
	encryptCmd := exec.Command("gpg", encryptArgs...)
	encryptCmd.Env = nil

	srcFile, fileErr := os.Open(src)
	if fileErr != nil {
		return fileErr
	}

	stat, statErr := srcFile.Stat()
	if statErr != nil {
		return statErr
	}

	r, w := io.Pipe()
	bufferedStdout := bufio.NewWriterSize(w, int(stat.Size()))
	decryptCmd.Stdout = bufferedStdout
	decryptCmd.Stderr = os.Stderr

	encryptCmd.Stderr = os.Stderr
	encryptCmd.Stdin = r

	log.Printf("Running: gpg %s | gpg %s\n", strings.Join(decryptArgs, " "), strings.Join(encryptArgs, " "))
	err1 := decryptCmd.Run()
	if err1 != nil {
		return err1
	}
	w.Close()

	err2 := encryptCmd.Run()
	if err2 != nil {
		return err2
	}

	return nil
}
