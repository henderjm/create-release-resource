package fakes

import "github.com/henderjm/create-release-resource/concourse"

type GithubClient struct {
	VersionsToReturn []string
	// HeadCommitSha    []string
	// ErrorToBeReturned       error
}

func (fakeGithubClient *GithubClient) GetListCommits(version string) (*concourse.VersionList, error) {
	return &concourse.VersionList{
		Number: fakeGithubClient.VersionsToReturn,
	}, nil
}
