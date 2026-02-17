package main

import "fmt"

// Config holds application configuration.
type Config struct {
	Name    string
	Version string
	Debug   bool
}

// NewConfig creates a default configuration.
func NewConfig() *Config {
	return &Config{
		Name:    "app",
		Version: "1.0.0",
		Debug:   false,
	}
}

// Print outputs the configuration.
func Print(c *Config) {
	fmt.Printf("Name: %s, Version: %s, Debug: %v\n", c.Name, c.Version, c.Debug)
}

func main() {
	c := NewConfig()
	Print(c)
}
