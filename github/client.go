package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/henderjm/create-release-resource/concourse"
)

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
	gc := github.NewClient(nil)

	repos, _, err := gc.Repositories.ListCommits(context.Background(), "henderjm", "test-repository", nil)
	if err != nil {
		fmt.Println("Error getting repo")
		fmt.Println(err)
		os.Exit(1)
	}
	// for _, shas := range repos {
	// 	fmt.Println(*shas.SHA)
	// }
	latestVersion := *repos[0].SHA
	return &concourse.Version{
		Number: latestVersion}, nil
}
