package version

var (
	// Version shows the current notation-azure-kv version, optionally with pre-release.
	Version = "v0.5.0-rc.1"

	// BuildMetadata stores the build metadata.
	BuildMetadata = "unreleased"
)

// GetVersion returns the version string in SemVer 2.
func GetVersion() string {
	if BuildMetadata == "" {
		return Version
	}
	return Version + "+" + BuildMetadata
}
