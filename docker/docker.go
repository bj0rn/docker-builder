package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	template, _ = template.Parse(
		`FROM {{.FromImage}}
    COPY {{.ProgramFiles}} /app
    CMD {{.Command}}`)

	output, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}

	defer output.Close()

	template.Execute(output, dockerFileVariables)

}

func BuildDockerImage(parameters *parameters.Parameters, folderName string) {
	filePath := fmt.Sprintf("%s/%s-%s", folderName, parameters.ArtifactId, parameters.Version)

	generateDockerFile(parameters.BaseImage, parameters.Command, filePath)
	src := fmt.Sprintf("%s/%s", parameters.OutputRegistry, parameters.OutputImage)

	fmt.Println("Building the docker image: ", src)
	cmd := exec.Command("docker", "build", "-t", src, ".")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()

	fmt.Println("Docker build suceeded")
}

func TagDockerImages(parameters *parameters.Parameters) {
	tags := strings.Split(parameters.Tags, " ")
	for _, tag := range tags {
		src := fmt.Sprintf("%s/%s", parameters.OutputRegistry, parameters.OutputImage)
		dest := fmt.Sprintf("%v:%v", src, tag)

		cmd := exec.Command("docker", "tag", src, dest)
		fmt.Println("Tagged docker image: ", dest)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		cmd.Wait()
	}
}

func PushDockerImages(parameters *parameters.Parameters) {
	tags := strings.Split(parameters.Tags, " ")
	for _, tag := range tags {
		image := fmt.Sprintf("%v/%v:%v", parameters.OutputRegistry, parameters.OutputImage, tag)
		cmd := exec.Command("docker", "push", image)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		cmd.Wait()

	}
}
