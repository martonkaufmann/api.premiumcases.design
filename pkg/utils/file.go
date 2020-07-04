package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func DownloadRemoteFile(path, url string) (string, error) {
	fn := fmt.Sprintf("%x", md5.Sum([]byte(url)))

	if FileExists(fn) {
		return fn, nil
	}

	ext := filepath.Ext(url)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	out, err := os.Create(path + "/" + fn + ext)

	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return "", err
	}

	return fn, nil
}
