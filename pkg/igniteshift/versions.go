package igniteshift

type Version struct {
	Releases GitHubRelease `json:"url"` // slice of urls via os/arch => url
}

func NewVersion(release *GitHubRelease) *Version {
	return &Version{
		Releases: *release,
	}
}

func (v *Version) IsNightly() bool {
	return false
}

func (v *Version) IsLatest() bool {
	return false
}
