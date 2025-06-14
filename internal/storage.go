package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type StorageManager struct {
	configDir string
}

func NewStorageManager() (*StorageManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	
	configDir := filepath.Join(homeDir, ".gurlz")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}
	
	return &StorageManager{configDir: configDir}, nil
}

func (sm *StorageManager) GetRequestsPath() string {
	return filepath.Join(sm.configDir, "requests.yaml")
}

func (sm *StorageManager) GetConfigPath() string {
	return filepath.Join(sm.configDir, "config.yaml")
}

func (sm *StorageManager) LoadRequests() (*RequestStore, error) {
	requestsPath := sm.GetRequestsPath()
	
	// Create empty file if it doesn't exist
	if _, err := os.Stat(requestsPath); os.IsNotExist(err) {
		store := &RequestStore{Requests: []Request{}}
		return store, sm.SaveRequests(store)
	}
	
	data, err := os.ReadFile(requestsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read requests file: %w", err)
	}
	
	var store RequestStore
	if err := yaml.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("failed to parse requests file: %w", err)
	}
	
	return &store, nil
}

func (sm *StorageManager) SaveRequests(store *RequestStore) error {
	data, err := yaml.Marshal(store)
	if err != nil {
		return fmt.Errorf("failed to marshal requests: %w", err)
	}
	
	requestsPath := sm.GetRequestsPath()
	if err := os.WriteFile(requestsPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write requests file: %w", err)
	}
	
	return nil
}

func (sm *StorageManager) LoadConfig() (*Config, error) {
	configPath := sm.GetConfigPath()
	
	// Create default config if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := sm.DefaultConfig()
		return config, sm.SaveConfig(config)
	}
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	
	return &config, nil
}

func (sm *StorageManager) SaveConfig(config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	
	configPath := sm.GetConfigPath()
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

func (sm *StorageManager) DefaultConfig() *Config {
	return &Config{
		DefaultHeaders: map[string]string{
			"User-Agent": "gurlz/1.0.0",
		},
		Timeout:         "30s",
		FollowRedirect:  true,
		SaveResponses:   true,
		OutputFormat:    "json",
		ColorOutput:     true,
		DefaultMethod:   "GET",
		MaxResponseSize: 1024 * 1024, // 1MB
	}
}

func (store *RequestStore) FindByName(name string) *Request {
	for i, req := range store.Requests {
		if req.Name == name {
			return &store.Requests[i]
		}
	}
	return nil
}

func (store *RequestStore) AddRequest(req Request) error {
	// Check for duplicate names
	if existing := store.FindByName(req.Name); existing != nil {
		return fmt.Errorf("request with name '%s' already exists", req.Name)
	}
	
	store.Requests = append(store.Requests, req)
	return nil
}