package version

// These variables are set at build time via ldflags
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// GetVersion returns the current version
func GetVersion() string {
	return Version
}

// GetCommit returns the git commit hash
func GetCommit() string {
	return Commit
}

// GetDate returns the build date
func GetDate() string {
	return Date
}

// GetFullVersion returns a formatted version string
func GetFullVersion() string {
	commitShort := Commit
	if len(Commit) > 7 {
		commitShort = Commit[:7]
	}
	return Version + " (" + commitShort + ") built on " + Date
}
