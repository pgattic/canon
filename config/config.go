package config

import (
  "encoding/json"
  "fmt"
  "os" 
  "path/filepath"
)

var TextsDir string = filepath.Join(getUserHomeDir(), ".canon", "texts")
var configPath string = filepath.Join(TextsDir, "config.json")

type Config struct {
  Priority []string `json:"priority"`
}

func LoadConfig() Config {

  file, err := os.ReadFile(configPath)
  if err != nil {
    panic(err)
  }

  // Unmarshal JSON into Config struct
  var config Config
  err_1 := json.Unmarshal(file, &config)
  if err_1 != nil {
    panic(err_1)
  }
  return config
}

func SaveConfig(config Config) {
  // Marshal the updated Config struct back to JSON
  updatedConfig, err := json.MarshalIndent(config, "", "    ")
  if err != nil {
    panic(err)
  }

  if err_1 := os.WriteFile(configPath, updatedConfig, 0644); err != nil {
    panic(err_1)
  }
}

func EnsureSetup() {
  if _, err := os.Stat(TextsDir); os.IsNotExist(err) {
    // Directory does not exist, create it
    err := os.MkdirAll(TextsDir, 0755) // 0755 is the directory permissions
    if err != nil {
      fmt.Println("Error creating directory:", err)
      return
    }
  } else if err != nil {
    // Some other error occurred
    fmt.Println("Error checking directory:", err)
    return
  }

  // Check if the config file exists
  if _, err := os.Stat(configPath); os.IsNotExist(err) {
    // Create the directory path if it doesn't exist
    if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
      fmt.Println("Error creating directory:", err)
      return
    }

    // Create a new config with "priority" as an empty list
    newConfig := Config{
      Priority: []string{},
    }

    // Marshal the new config into JSON format
    configData, err := json.MarshalIndent(newConfig, "", "  ")
    if err != nil {
      fmt.Println("Error encoding config to JSON:", err)
      return
    }

    // Write the JSON data to the config file
    err = os.WriteFile(configPath, configData, 0644)
    if err != nil {
      fmt.Println("Error writing config file:", err)
      return
    }
  } else if err != nil {
    fmt.Println("Error checking config file:", err)
    return
  }
}

func getUserHomeDir() string {
  if homeDir, err := os.UserHomeDir(); err == nil {
    return homeDir
  }
  // Fallback option if getting user's home directory fails
  return os.Getenv("HOME")
}

