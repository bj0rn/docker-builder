package parameters

import (
	"os"
)

type Parameters struct {
	Url            string
	OutputRegistry string
	OutputImage    string
	Tags           string
	GroupId        string
	ArtifactId     string
	Version        string
	BaseImage      string
	Command        string
}

func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Get() *Parameters {
	baseImage := getEnvOrDefault("BASE_IMAGE", "")
	groupId := getEnvOrDefault("GROUP_ID", "")
	artifactId := getEnvOrDefault("ARTIFACT_ID", "")
	version := getEnvOrDefault("VERISON", "")
	tags := getEnvOrDefault("TAGS", "latest")
	outputRegistry := getEnvOrDefault("OUTPUT_REGISTRY", "")
	outputImage := getEnvOrDefault("OUTPUT_IMAGE", "")
	url := getEnvOrDefault("URL", "")
	command := getEnvOrDefault("COMMAND", "app/start.sh")

	return &Parameters{
		Url:            url,
		OutputRegistry: outputRegistry,
		OutputImage:    outputImage,
		Tags:           tags,
		GroupId:        groupId,
		ArtifactId:     artifactId,
		Version:        version,
		BaseImage:      baseImage,
		Command:        command,
	}

}
