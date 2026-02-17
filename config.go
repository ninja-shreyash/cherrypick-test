package main

import (
	"fmt"
	"os"
)

// Config holds application configuration.
type Config struct {
<<<<<<< HEAD
	AppName    string
	AppVersion string
	DebugMode  bool
	LogLevel   string
=======
	Name    string
	Version string
	Debug   bool
	Env     string
>>>>>>> 4ca566c (Add environment support and email validation)
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	return &Config{
<<<<<<< HEAD
		AppName:    "myapp",
		AppVersion: "1.0.1-release",
		DebugMode:  false,
		LogLevel:   "info",
	}
}

// PrintConfig outputs the configuration details.
func PrintConfig(c *Config) {
	fmt.Printf("App: %s v%s (debug=%v, log=%s)\n", c.AppName, c.AppVersion, c.DebugMode, c.LogLevel)
=======
		Name:    "app",
		Version: "2.0.0",
		Debug:   true,
		Env:     env,
	}
}

// Print outputs the configuration.
func Print(c *Config) {
	fmt.Printf("Name: %s, Version: %s, Env: %s, Debug: %v\n", c.Name, c.Version, c.Env, c.Debug)
>>>>>>> 4ca566c (Add environment support and email validation)
}

func main() {
	c := NewConfig()
<<<<<<< HEAD
	PrintConfig(c)
=======
	Print(c)
	fmt.Printf("Running in %s mode\n", c.Env)
>>>>>>> 4ca566c (Add environment support and email validation)
}
