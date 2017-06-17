package github

import "github.com/henderjm/create-release-resource/concourse"

type DeployParams struct {
	Final      bool
	Repository string
	User       string
	Password   string
}

type Client struct {
	username   string
	password   string
	repository string
	branch     string
}

func NewClient(uname, pwd, repo, gitBranch string) *Client {
	return &Client{username: uname, password: pwd, repository: repo,
		branch: gitBranch}
}

type GithubClient interface {
	GetHead(version string) (*concourse.Version, error)
}

func (client *Client) GetHead(version string) (*concourse.Version, error) {
	return nil, nil
}
