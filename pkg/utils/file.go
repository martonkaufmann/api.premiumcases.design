package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadRemoteFile(path, url string) (string, error) {
	fn := fmt.Sprintf("%x", md5.Sum([]byte(url)))
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
