package program

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/bj0rn/docker-builder/parameters"
)

func Unzip(filename string) {
	_, err := exec.Command("unzip", "-d", "app", filename).Output()
	if err != nil {
		fmt.Println("Failed to unzip file", filename)
		return
	}
	fmt.Println("Unzipped file", filename)
}

func prepareUrl(parameters *parameters.Parameters) string {

	return parameters.Url
}

func DownloadFromUrl(parameters *parameters.Parameters) string {
	url := prepareUrl(parameters)
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, " to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return ""
	}

	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading file", fileName, "-", err)
		return ""
	}

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while donwloading file", fileName, "-", err)
		return ""
	}

	fmt.Println("Downloaded", n, "bytes")

	return fileName
}
