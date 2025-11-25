package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Provider string `json:"provider"`
}

// configDir determines the configuration directory path.
func configDir() (string, error) {
	home, err := os.UserHomeDir() // Get the user's home directory.
	if err != nil {
		return "", fmt.Errorf("could not determine user home: %w", err)
	}
	return filepath.Join(home, ".autocommenter"), nil // Join home directory with the config directory name.
}

// configPath determines the full path to the configuration file.
func configPath() (string, error) {
	dir, err := configDir() // Get the configuration directory.
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil // Join the config directory with the filename.
}

// Load reads the configuration from the file.
func Load() (*Config, error) {
	path, err := configPath() // Get the configuration file path.
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path) // Attempt to open the configuration file.
	if os.IsNotExist(err) {
		// default config
		return &Config{Provider: "gemini"}, nil // If the file doesn't exist, return a default configuration.
	}
	if err != nil {
		return nil, err
	}
	defer f.Close() // Ensure the file is closed after the function returns.

	var cfg Config
	dec := json.NewDecoder(f) // Create a JSON decoder for the file.
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	if cfg.Provider == "" {
		cfg.Provider = "gemini" // If provider is not set in the config, default to "gemini".
	}
	return &cfg, nil
}

// Save writes the configuration to the file.
func Save(cfg *Config) error {
	dir, err := configDir() // Get the configuration directory.
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil { // Create the configuration directory if it doesn't exist.
		return fmt.Errorf("could not create config dir: %w", err)
	}

	path := filepath.Join(dir, "config.json")
	tmp := path + ".tmp" // Use a temporary file for atomic writes.

	f, err := os.Create(tmp) // Create the temporary file.
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ") // Set indentation for pretty JSON output.
	if err := enc.Encode(cfg); err != nil {
		f.Close()            // Close the file before removing.
		os.Remove(tmp)       // Remove the temporary file on encoding error.
		return fmt.Errorf("failed to write config: %w", err)
	}
	if err := f.Close(); err != nil { // Close the temporary file.
		os.Remove(tmp) // Remove the temporary file if closing fails.
		return err
	}
	return os.Rename(tmp, path) // Atomically rename the temporary file to the final configuration file.
}

// SetProvider helper stores the provider name
func SetProvider(name string) error {
	if name == "" {
		return fmt.Errorf("provider name cannot be empty")
	}
	cfg := &Config{Provider: name}
	return Save(cfg)
}

// GetProvider returns the saved provider or default "gemini"
func GetProvider() (string, error) {
	cfg, err := Load() // Load the configuration.
	if err != nil {
		return "gemini", err
	}
	if cfg.Provider == "" {
		return "gemini", nil // Return default if provider is empty in loaded config.
	}
	return cfg.Provider, nil
}