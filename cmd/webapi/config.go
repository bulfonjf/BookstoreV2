package main

const (
	// DefaultDSN is the default datasource name.
	DefaultDSN = "sqlite::memory:"
)

type Config struct {
	DB struct {
		DSN string
	}

	HTTP struct {
		Addr     string
		Domain   string
	}
}

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	config.DB.DSN = DefaultDSN
	config.HTTP.Addr = ":3000"
	return config
}
