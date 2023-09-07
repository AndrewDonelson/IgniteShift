package igniteshift

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	githubAPIURL = "https://api.github.com/repos/%s/%s/releases"
)

// GitHubUser represents a GitHub user or bot, associated with a release or asset.
type GitHubUser struct {
	Login     string `json:"login"`
	ID        int64  `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"url"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}

// GitHubAsset represents a downloadable asset of a release.
type GitHubAsset struct {
	URL                string     `json:"url"`
	ID                 int64      `json:"id"`
	NodeID             string     `json:"node_id"`
	Name               string     `json:"name"`
	Label              string     `json:"label"`
	Uploader           GitHubUser `json:"uploader"`
	ContentType        string     `json:"content_type"`
	State              string     `json:"state"`
	Size               int64      `json:"size"`
	DownloadCount      int        `json:"download_count"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	BrowserDownloadURL string     `json:"browser_download_url"`
}

// GitHubRelease represents a release in a GitHub repository.
type GitHubRelease struct {
	URL             string        `json:"url"`
	AssetsURL       string        `json:"assets_url"`
	UploadURL       string        `json:"upload_url"`
	HTMLURL         string        `json:"html_url"`
	ID              int64         `json:"id"`
	NodeID          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       string        `json:"created_at"`
	PublishedAt     string        `json:"published_at"`
	Assets          []GitHubAsset `json:"assets"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Body            string        `json:"body"`
	Author          GitHubUser    `json:"author"`
}

type GitHubReleaser interface {
	FetchReleases(owner, repo string) ([]byte, error)
}

type RealReleaser struct{}

func (r *RealReleaser) FetchReleases(owner, repo string) ([]byte, error) {
	// Actual network request to GitHub API
	resp, err := http.Get(fmt.Sprintf(githubAPIURL, owner, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch releases: %v", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

type MockReleaser struct{}

func (m *MockReleaser) FetchReleases(owner, repo string) ([]byte, error) {
	// Return the mock data from a file
	if strings.ToLower(owner) == "ignite" && strings.ToLower(repo) == "cli" {
		return os.ReadFile("../../test/ignite_releases.json")
	} else if strings.ToLower(owner) == "andrewdonelson" && strings.ToLower(repo) == "igniteshift" {
		return os.ReadFile("../../test/tooling_releases.json")
	} else {
		return nil, fmt.Errorf("unknown owner/repo: %s/%s", owner, repo)
	}
}

func GetGitHubReleases(releaser GitHubReleaser, owner, repo string) ([]GitHubRelease, error) {
	data, err := releaser.FetchReleases(owner, repo)
	if err != nil {
		return nil, err
	}

	var releases []GitHubRelease
	err = json.Unmarshal(data, &releases)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return releases, nil
}
