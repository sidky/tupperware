package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleConfig(t *testing.T) {
	config := `
configs:
- name: "Foo Bar"
  path: "foo/bar"
  githubOwner: "gown"
  githubRepoName: "repo"
  githubRepoHost: "host"
  baseBranch: "base"
`
	parsed, err := ParseConfigFromBytes([]byte(config))

	assert.NoError(t, err, "Unable to parse")

	if parsed == nil || parsed.RepoConfigs == nil || len(parsed.RepoConfigs) != 1 {
		t.Fatalf("Parsed config should contain 1 element, actual: %v", parsed)
	}

	pc := parsed.RepoConfigs[0]

	assert.Equalf(t, "Foo Bar", pc.RepoName, "Invalid repo name")
	assert.Equalf(t, "foo/bar", pc.RepoPath, "Invalid repo path")
	assert.Equalf(t, "gown", pc.GithubRepoOwner, "Invalid repo owner")
	assert.Equalf(t, "repo", pc.GithubRepoName, "Invalid repo name")
	assert.Equalf(t, "host", pc.GithubRepoHost, "Invalid repo host")
	assert.Equalf(t, "base", pc.BaseBranch, "Invalid base branch")

}

func TestMultipleConfig(t *testing.T) {
	config := `
configs:
- name: "Foo Bar"
  path: "foo/bar"
  githubOwner: "gown"
  githubRepoName: "repo"
  githubRepoHost: "host"
  baseBranch: "base"
- name: "Second Repo"
  path: "foo/baz"
  githubOwner: "gown"
  githubRepoName: "repo2"
  githubRepoHost: "host"
  baseBranch: "develop"
`
	parsed, err := ParseConfigFromBytes([]byte(config))

	assert.NoError(t, err, "Unable to parse")

	if parsed == nil || parsed.RepoConfigs == nil || len(parsed.RepoConfigs) != 2 {
		t.Fatalf("Parsed config should contain 2 element, actual: %v", parsed)
	}

	pc1 := parsed.RepoConfigs[0]

	assert.Equalf(t, "Foo Bar", pc1.RepoName, "Invalid repo name")
	assert.Equalf(t, "foo/bar", pc1.RepoPath, "Invalid repo path")
	assert.Equalf(t, "gown", pc1.GithubRepoOwner, "Invalid repo owner")
	assert.Equalf(t, "repo", pc1.GithubRepoName, "Invalid repo name")
	assert.Equalf(t, "host", pc1.GithubRepoHost, "Invalid repo host")
	assert.Equalf(t, "base", pc1.BaseBranch, "Invalid base branch")

	pc2 := parsed.RepoConfigs[1]

	assert.Equalf(t, "Second Repo", pc2.RepoName, "Invalid repo name")
	assert.Equalf(t, "foo/baz", pc2.RepoPath, "Invalid repo path")
	assert.Equalf(t, "gown", pc2.GithubRepoOwner, "Invalid repo owner")
	assert.Equalf(t, "repo2", pc2.GithubRepoName, "Invalid repo name")
	assert.Equalf(t, "host", pc2.GithubRepoHost, "Invalid repo host")
	assert.Equalf(t, "develop", pc2.BaseBranch, "Invalid base branch")
}
