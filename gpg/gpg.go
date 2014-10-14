package gpg

import (
        "os"
        "os/exec"
        "path"
        "strings"
)



func Decrypt(filePath string) (string, error) {


    decryptArgs  := []string {"--decrypt", "--armor", "--batch", "--yes", filePath}


    cmd := exec.Command("gpg", decryptArgs...)

    out, err := cmd.Output()

    if err != nil {
        return "", err
    }

    return string(out), nil
}



func Encrypt(filePath string, text string, recipients []string) error {

    if err := os.MkdirAll(path.Dir(filePath), 0777); err != nil {
        return err
    }

    encryptArgs  := []string {"--encrypt", "--armor", "--batch", "--yes", "--output", filePath}

    for _, recipient := range recipients {
        encryptArgs = append(encryptArgs, "--recipient")
        encryptArgs = append(encryptArgs, recipient)
    }

    cmd := exec.Command("gpg", encryptArgs...)
    cmd.Stdin = strings.NewReader(text)
    _, err := cmd.CombinedOutput()

    if err != nil {
        return err
    }
    return nil
}

func EncryptFile(filePath string, sourceFile string, recipients []string) error {

    if err := os.MkdirAll(path.Dir(filePath), 0777); err != nil {
        return err
    }

    encryptArgs  := []string {"--encrypt", "--armor", "--batch", "--yes", "--output", filePath}

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
