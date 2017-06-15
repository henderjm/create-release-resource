package fakes

type GithubClient struct {
	VersionsToReturn []string
	HeadCommitSha    []string
	// ErrorToBeReturned       error
}

func (fakeGithubClient *GithubClient) GetHead(version string) error {

	return nil
}
