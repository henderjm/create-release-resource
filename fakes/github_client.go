package fakes

import "github.com/henderjm/create-release-resource/concourse"

type GithubClient struct {
	VersionsToReturn string
	// HeadCommitSha    []string
	// ErrorToBeReturned       error
}

func (fakeGithubClient *GithubClient) GetHead(version string) (*concourse.Version, error) {
	return &concourse.Version{
		Number: fakeGithubClient.VersionsToReturn,
	}, nil
}
