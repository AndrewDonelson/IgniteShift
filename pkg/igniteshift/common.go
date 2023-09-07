package igniteshift

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// checkForInternetConnection checks if there is an active internet connection.
func checkForInternetConnection() bool {
	// Define a timeout duration for the request
	timeout := time.Duration(5 * time.Second)

	// Create a new HTTP client with the defined timeout
	client := http.Client{
		Timeout: timeout,
	}

	// Make a GET request to a well-known, reliable domain to check for an internet connection
	_, err := client.Get("https://www.google.com")

	// If the request completed without errors, there's an active internet connection
	return err == nil
}

// CurrentOSArch returns the operating system and architecture the tool is running on.
func CurrentOSArch() (os, arch string) {
	return runtime.GOOS, runtime.GOARCH
}

func isInTests() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.v=") || strings.HasSuffix(os.Args[0], ".test") || strings.HasSuffix(os.Args[0], "debug_bin") {
			return true
		}
	}
	return false
}

// getIgniteReleases returns a list of available Ignite versions from the GitHub repository.
func getIgniteReleases() (releases []GitHubRelease, err error) {

	releases, err = GetGitHubReleases(&MockReleaser{}, "ignite", "cli")
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// gettoolingReleases returns a list of available tooling versions from the GitHub repository.
func getToolingReleases() (releases []GitHubRelease, err error) {
	releases, err = GetGitHubReleases(&MockReleaser{}, "andrewdonelson", "igniteshift")
	if err != nil {
		return nil, err
	}

	return releases, nil

}
