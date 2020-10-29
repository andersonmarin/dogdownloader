package usecases

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var downloadDirectory string

func init() {
	flag.StringVar(&downloadDirectory, "d", "downloads", "download directory")
	flag.Parse()

	err := os.MkdirAll(downloadDirectory, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func DownloadImage(src string) error {
	res, err := http.Get(src)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	filename := fmt.Sprintf("%s/%d.jpg", downloadDirectory, time.Now().UnixNano())
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, res.Body)

	return err
}
