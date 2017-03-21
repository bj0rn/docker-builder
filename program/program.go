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

//http://localhost:8081/nexus/service/local/artifact/maven/redirect?r=snapshots&g=github.com.bj0rn&a=spring-openshift-demo&v=LATEST&c=bin&p=zip
func prepareUrl(parameters *parameters.Parameters) string {
	baseURL := parameters.Url
	groupID := parameters.GroupId
	artifactID := parameters.ArtifactId
	version := parameters.Version
	distributionType := parameters.DistributionType

	if parameters.DistributionManager == "nexus" {
		if strings.Contains(parameters.Version, "SNAPSHOT") {
			return fmt.Sprintf("%s/service/local/artifact/maven/redirect?r=snapshots&g=%s&a=%s&v=%s&c=bin&p=%s", baseURL, groupID, artifactID, version, distributionType)
		} else {
			return fmt.Sprintf("%s/service/local/artifact/maven/redirect?r=releases&g=%s&a=%s&v=%s&c=bin&p=%s", baseURL, groupID, artifactID, version, distributionType)
		}
	}
	return parameters.Url
}

func DownloadFromUrl(parameters *parameters.Parameters, filename string) {
	url := prepareUrl(parameters)
	fmt.Println("Downloading", url, " to", filename)

	output, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error while creating", filename, "-", err)
		return
	}

	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading file", filename, "-", err)
		return
	}

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while donwloading file", filename, "-", err)
		return
	}

	fmt.Println("Downloaded", n, "bytes")

}
