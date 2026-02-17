package main

import (
	"fmt"
	"os"
)

// Config holds application configuration.
type Config struct {
	Name    string
	Version string
	Debug   bool
	Env     string
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	return &Config{
		Name:    "app",
		Version: "2.0.0",
		Debug:   true,
		Env:     env,
	}
}

// Print outputs the configuration.
func Print(c *Config) {
	fmt.Printf("Name: %s, Version: %s, Env: %s, Debug: %v\n", c.Name, c.Version, c.Env, c.Debug)
}

func main() {
	c := NewConfig()
	Print(c)
	fmt.Printf("Running in %s mode\n", c.Env)
}
