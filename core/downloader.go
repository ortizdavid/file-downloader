package core

import (
	"net/http"
	"os"
	"io"
	
)

type Downloader struct{}


func (dl *Downloader) DownloadFile(filePath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}