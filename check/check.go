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

type CheckResponse []concourse.Version

func (c CheckCommand) Execute(checkRequest CheckRequest) (CheckResponse, error) {
	response, err := c.client.GetListCommits(checkRequest.Version.Number)
	if err != nil {
		fmt.Println("Here lies an error")
		return nil, err
	}
	output := CheckResponse{}

	// check head
	if response.Number[0] == checkRequest.Version.Number {
		return output, nil
	}

	for _, num := range response.Number {
		output = append(output, concourse.Version{Number: num})
	}

	for indx, version := range output {
		if version.Number == checkRequest.Version.Number {
			// slice
			output = output[:indx]
		}
	}

	return output, nil
}
