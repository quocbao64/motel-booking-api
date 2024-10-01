package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"path/filepath"
	"time"
)

func SaveFile(base64Str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("failed to decode base64 string")
	}

	fileName := time.Now().Format("20060102150405") + ".pdf"
	filePath := filepath.Join("files", fileName)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", errors.New("failed to write file")
	}

	return filePath, nil
}

func HashFileBase64ToSHA256(fileBase64 string) (string, error) {
	fileBytes, err := base64.StdEncoding.DecodeString(fileBase64)
	if err != nil {
		return "", errors.New("failed to decode base64 string")
	}

	hash := sha256.New()
	hash.Write(fileBytes)
	hashBytes := hash.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)

	return hashHex, nil
}

func ConvertFileToBase64(filePath string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("failed to read file")
	}

	fileBase64 := base64.StdEncoding.EncodeToString(fileBytes)

	return fileBase64, nil
}
