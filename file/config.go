package file

// ConfigDict holds the basic configuration information for the gorrent-server application,
// including basic information on how many workers to spin up, certain timing information,
// and similar properties (that would usually be hardcoded).
type ConfigDict struct {
	Workers WorkersDict
}

// WorkersDict holds configuration values specifically for concurrent worker functions, including
// how many to spin up at a time, how many should always be available, and how many to allocate
// at max capacity.
type WorkersDict struct {
	MinCount int `yaml:"min_count"`
	MaxCount int `yaml:"max_count"`
}
