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

	//Find a cleaner solution
	fileName := program.DownloadFromUrl(params)
	if strings.Contains(fileName, ".zip") {
		program.Unzip(fileName)
	} else if strings.Contains(fileName, ".jar") {
		params.Command = fmt.Sprintf("java -jar %s", fileName)
	}

	docker.BuildDockerImage(params, fileName)

	docker.TagDockerImages(params)

	docker.PushDockerImages(params)

}
