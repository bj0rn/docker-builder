package docker

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/bj0rn/docker-builder/parameters"
)

type DockerFileVariables struct {
	FromImage    string
	ProgramFiles string
	Command      string
}

func generateDockerFile(fromImage, command, fileName string) {
	dockerFileVariables := DockerFileVariables{
		FromImage:    fromImage,
		ProgramFiles: fileName,
		Command:      command,
	}
	template := template.New("Dockerfile")

	template, _ = template.Parse(`FROM {{.FromImage}}
    COPY {{.ProgramFiles}} /app
    ENTRYPOINT ['{{.Command}}']`)

	output, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}

	defer output.Close()

	template.Execute(output, dockerFileVariables)

}

func BuildDockerImage(parameters *parameters.Parameters, fileName string) {
	generateDockerFile(parameters.BaseImage, parameters.Command, fileName)
	src := fmt.Sprintf("%s/%s", parameters.OutputRegistry, parameters.OutputImage)

	res, err := exec.Command("docker", "build", "-t", src, ".").Output()
	if err != nil {
		fmt.Println("Docker build failed with msg", err)
		return
	}
	fmt.Println("Docker build suceeded with msg", res)
}

func TagDockerImages(parameters *parameters.Parameters) {
	for _, tag := range parameters.Tags {
		src := fmt.Sprintf("%s/%s", parameters.OutputRegistry, parameters.OutputImage)
		dest := fmt.Sprintf("%s:%s", src, tag)

		res, err := exec.Command("docker", "tag", src, dest).Output()
		if err != nil {
			fmt.Println("Failed to tag image", dest)
			return
		}
		fmt.Println("Failed to tag image", res)
	}
}

func PushDockerImages(parameters *parameters.Parameters) {
	for _, tag := range parameters.Tags {
		image := fmt.Sprintf("%s/%s:%s", parameters.OutputRegistry, parameters.OutputImage, tag)
		res, err := exec.Command("docker", "push", image).Output()
		if err != nil {
			fmt.Println("Failed to push image", image)
			return
		}
		fmt.Println("Pushed image with msg:", res)
	}
}
