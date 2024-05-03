package version

var (
	// -ldflags="-X 'main.Version=vX.X.X'"
	Version = "1.0.0"
)

func GetVersion() string {
	version := Version

	return version
}
