package main

import (
	"fmt"
	"os"
)

// Config holds application configuration.
type Config struct {
	AppName    string
	AppVersion string
	DebugMode  bool
	LogLevel   string
	Env        string
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	return &Config{
		AppName:    "myapp",
		AppVersion: "1.0.1-release",
		DebugMode:  false,
		LogLevel:   "info",
		Env:        env,
	}
}

// PrintConfig outputs the configuration details.
func PrintConfig(c *Config) {
	fmt.Printf("App: %s v%s (env=%s, debug=%v, log=%s)\n", c.AppName, c.AppVersion, c.Env, c.DebugMode, c.LogLevel)
}

func main() {
	c := NewConfig()
	PrintConfig(c)
	fmt.Printf("Running in %s mode\n", c.Env)
}
