package check

import (
	"fmt"

	"github.com/henderjm/create-release-resource/concourse"
	"github.com/henderjm/create-release-resource/github"
)

type CheckCommand struct {
	client github.GithubClient
}

type Version struct {
	CommitSha string `json:"commit_sha"`
}

type CheckResponse []concourse.Version

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
	gitHeadSha, err := c.client.GetHead(checkRequest.Version.Number)
	if err != nil {
		fmt.Println("Here lies an error")
		return nil, err
	}
	output := CheckResponse{}

	output = append(output, concourse.Version{Number: gitHeadSha.Number})
	return output, nil
}
