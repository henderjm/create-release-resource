package fakes

import (
	"fmt"

	"github.com/henderjm/create-release-resource/concourse"
)

type GithubClient struct {
	VersionsToReturn string
	// HeadCommitSha    []string
	// ErrorToBeReturned       error
}

func (fakeGithubClient *GithubClient) GetHead(version string) (*concourse.Version, error) {
	fmt.Println(fakeGithubClient.VersionsToReturn)
	return &concourse.Version{
		Number: fakeGithubClient.VersionsToReturn,
	}, nil
}
