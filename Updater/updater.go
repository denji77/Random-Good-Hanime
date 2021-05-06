package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

var (
	filename    string
	downloadUrl string
)

func main() {
	checkOs()
	downloadUpdate()
}

func checkOs() {
	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-windows-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "Random-Good-Hanime-windows-arm"
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-linux-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "Random-Good-Hanime-linux-arm"
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "amd64" {
			filename = "Random-Good-Hanime-darwin-amd64"
		}
		if runtime.GOARCH == "arm" {
			fmt.Println("Error, Darwin arm not supported")
			os.Exit(1)
		}
	}

	fmt.Println("Currently running on: " + runtime.GOOS + " " + runtime.GOARCH)

	filename = filename + ".zip"
}

func downloadUpdate() {
	api := "https://api.github.com/repos/superredstone/random-good-hanime/releases"

	response, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data []Githubapi
	json.Unmarshal([]byte(responseData), &data)

	for i := 0; i < len(data[0].Assets); i++ {
		if data[0].Assets[i].Name == filename {
			fmt.Println(data[0].Assets[i])
			downloadUrl = data[0].Assets[i].Url

			break
		}
	}

	err = DownloadFile(filename, downloadUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded new version: ", filename)

	fmt.Println("Unzip it and replace the old one, if there are problems with the config file delete it")
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

type Githubapi struct {
	Assets []Assets `json:"assets"`
}

type Assets struct {
	Url  string `json:"browser_download_url"`
	Name string `json:"name"`
}
