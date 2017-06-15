package concourse

type Version struct {
	Number string `json:"number"`
}

type Source struct {
	Branch   string `json:"branch"`
	RepoName string `json:"repo_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
