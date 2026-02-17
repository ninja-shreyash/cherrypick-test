package main

import "fmt"

// Config holds application configuration.
type Config struct {
	AppName    string
	AppVersion string
	DebugMode  bool
	LogLevel   string
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	return &Config{
		AppName:    "myapp",
		AppVersion: "1.0.1-release",
		DebugMode:  false,
		LogLevel:   "info",
	}
}

// PrintConfig outputs the configuration details.
func PrintConfig(c *Config) {
	fmt.Printf("App: %s v%s (debug=%v, log=%s)\n", c.AppName, c.AppVersion, c.DebugMode, c.LogLevel)
}

func main() {
	c := NewConfig()
	PrintConfig(c)
}
