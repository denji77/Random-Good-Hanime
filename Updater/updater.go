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
	filename string
)

func main() {
	checkOs()
	downloadUpdate()
	checkHash()
}

func checkOs() {
	//Salvare il filename in base al sistema operativo

	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			filename = "RandomHentai-windows-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "RandomHentai-windows-arm"
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "amd64" {
			filename = "RandomHentai-linux-amd64"
		}
		if runtime.GOARCH == "arm" {
			filename = "RandomHentai-linux-arm"
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "amd64" {
			filename = "RandomHentai-darwin-amd64"
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

	fmt.Println(data)

	/*err = DownloadFile(filename, data.Assets.Url)
	if err != nil {
		fmt.Println(err)
	}*/

	fmt.Println("Downloaded new version: ", filename)
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

func checkHash() {
	// Controllare l'hash
}

type Githubapi struct {
	Assets []Assets `json:"assets"`
}

type Assets struct {
	Url  string `json:"browser_download_url"`
	Name string `json:"name"`
}