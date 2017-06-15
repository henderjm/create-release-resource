package check

import (
	"fmt"

	"github.com/henderjm/create-release-resource/github"
)

type CheckCommand struct {
	client github.GithubClient
}

type Version struct {
	CommitSha string `json:"commit_sha"`
}

type CheckResponse []Version

func NewVersion(commitsha string) Version {
	return Version{
		CommitSha: commitsha,
	}
}

func NewCheckCommand(client github.GithubClient) CheckCommand {
	return CheckCommand{
		client: client,
	}
}

func (c CheckCommand) Execute(checkRequest CheckRequest) (CheckResponse, error) {
	output := []Version{}
	version := NewVersion("222222")
	fmt.Println("version is " + version.CommitSha)
	output = append(output, version)
	return output, nil
}
