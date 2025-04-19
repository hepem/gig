package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hepem/gig/constants"
)

func CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func GenerateSHA1(name string) string {
	data := time.Now().Format(time.RFC3339) + name
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])

}

func FileToSHA1(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func Deflate(data []byte) ([]byte, error) {
	compressed := bytes.NewBuffer(nil)
	writer := zlib.NewWriter(compressed)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	writer.Close()
	return compressed.Bytes(), nil
}

func CreateDirFromSha(sha string) {
	err := CreateDir(fmt.Sprintf("%s/%s", constants.ObjectDir, sha[0:2]))
	if err != nil {
		fmt.Println("Error creating object directory:", err)
		return
	}
}
