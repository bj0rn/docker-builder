package main

import (
	"fmt"
	"strings"

	"github.com/bj0rn/docker-builder/docker"
	"github.com/bj0rn/docker-builder/parameters"
	"github.com/bj0rn/docker-builder/program"
)

func main() {
	//Gather input variables in struct
	params := parameters.Get()

	filename := fmt.Sprintf("%s.%s", "app", params.DistributionType)
	//Find a cleaner solution
	program.DownloadFromUrl(params, filename)
	if strings.Contains(filename, ".zip") {
		params.Command = fmt.Sprintf("java -jar app/%s-%s.jar", params.ArtifactId, params.Version)
		program.Unzip(filename)

	} else if strings.Contains(filename, ".jar") {
		params.Command = fmt.Sprintf("java -jar %s", filename)

	}

	docker.BuildDockerImage(params)
	docker.TagDockerImages(params)

	docker.PushDockerImages(params)

}
