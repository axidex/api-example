package config_provider

var (
	// buildVersion is the application build version, example: v1.0.0.
	buildVersion = "N/A"

	// buildDate is the application build date, example: 01.02.2025.
	buildDate = "N/A"

	// buildCommit is the application build commit, example: abcd1234.
	buildCommit = "N/A"
)

type Version interface {
	Version() string
	Date() string
	Commit() string
}

// BuildVersion is provides application build version details.
type BuildVersion struct {
	BuildVersion string
	BuildDate    string
	BuildCommit  string
}

// NewVersion creates a new version instance.
func NewVersion() Version {
	return &BuildVersion{
		BuildVersion: buildVersion,
		BuildDate:    buildDate,
		BuildCommit:  buildCommit,
	}
}

// Version returns the application build version.
func (v *BuildVersion) Version() string {
	return v.BuildVersion
}

// Date returns the application build date.
func (v *BuildVersion) Date() string {
	return v.BuildDate
}

// Commit returns the application build commit.
func (v *BuildVersion) Commit() string {
	return v.BuildCommit
}
