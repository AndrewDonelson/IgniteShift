package igniteshift

import (
	"testing"
)

// TestgetIgniteReleases tests the getIgniteReleases function.
func TestGetIgniteReleases(t *testing.T) {

	releases, err := getIgniteReleases()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(releases) != 30 { // Adjust this based on the actual number of releases in ignite_releases.json
		t.Fatalf("expected 2 releases, got %d", len(releases))
	}
}

// TestgettoolingReleases tests the gettoolingReleases function.
func TestGetToolingReleases(t *testing.T) {

	releases, err := getToolingReleases()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(releases) != 0 { // Adjust this based on the actual number of releases in ignite_releases.json
		t.Fatalf("expected 0 releases, got %d", len(releases))
	}
}

// TestCurrentOSArch tests the CurrentOSArch function.
func TestCurrentOSArch(t *testing.T) {
	os, arch := CurrentOSArch()
	if os == "" || arch == "" {
		t.Fatalf("expected non-empty os and arch, got os: %s, arch: %s", os, arch)
	}
}
