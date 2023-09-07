package igniteshift

import "time"

type Manager struct {
	latestToolingVersion    string    // latestToolingVersion is the latest version of the tooling obtained from https://github.com/AndrewDonelson/IgniteShift/releases
	currentToolingVersion   string    // currentToolingVersion is the current version of the tooling
	avaialbleIgniteVersions []string  // avaialbleIgniteVersions is a list of available Ignite versions from https://github.com/ignite/cli/releases
	currentIgniteVersion    string    // currentIgniteVersion is the current version of Ignite
	lastCheckedDate         time.Time // lastCheckedDate is the last time the tooling version was checked (only check once per day)
}

var (
	manager *Manager
)

func init() {
	manager = NewManager()
	if checkForInternetConnection() {
		manager.CheckForUpdates()
		manager.CheckForToolingUpdates()
		manager.CheckForIgniteUpdates()
	}
}

// NewManager creates a new Manager
func NewManager() *Manager {
	return &Manager{}
}

func GetManager() *Manager {
	return manager
}

func (m *Manager) Welcome() {
	println("Welcome to IgniteShift")
}

func (m *Manager) CheckForUpdates() error {
	return nil
}

func (m *Manager) CheckForToolingUpdates() error {
	return nil
}

func (m *Manager) CheckForIgniteUpdates() error {
	return nil
}

func (m *Manager) CheckForIgniteVersion(version string) error {
	return nil
}

func (m *Manager) CheckForToolingVersion(version string) error {
	return nil
}

func (m *Manager) GetLatestToolingVersion() string {
	return ""
}

func (m *Manager) GetCurrentToolingVersion() string {
	return ""
}
