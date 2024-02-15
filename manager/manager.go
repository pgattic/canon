package manager // haha get it? lol

import (
  "fmt"
  "os"
  "os/exec"
  "encoding/json"
  "path/filepath"
)

type Config struct {
  Priority []string `json:"priority"`
}

type Manager struct {
  
}

func GitClone(repoURL string, repoDir string) {
  // Define the path to the texts directory
  textsDir := filepath.Join(os.Getenv("HOME"), ".canon", "texts")

  // Ensure the texts directory exists
  //    if err := os.MkdirAll(textsDir, 0755); err != nil {
  //        fmt.Printf("Error creating texts directory: %v\n", err)
  //        return
  //    }

  // Clone the Git repository into the texts directory
  if err := cloneRepo(repoURL, textsDir + "/" + repoDir); err != nil {
    fmt.Printf("Error cloning repository: %v\n", err)
    return
  }

  // Update the config.json with the repository name
  configPath := filepath.Join(textsDir, "config.json")
  if err := updateConfig(configPath, repoDir); err != nil {
    fmt.Printf("Error updating config.json: %v\n", err)
    return
  }

  fmt.Println("Repository cloned successfully and config updated.")
}

func cloneRepo(repoURL, destination string) error {
  cmd := exec.Command("git", "clone", repoURL, destination)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func updateConfig(configPath, repoName string) error {
  // Read existing config file
  file, err := os.ReadFile(configPath)
  if err != nil {
    return err
  }

  // Unmarshal JSON into Config struct
  var config Config
  if err := json.Unmarshal(file, &config); err != nil {
    return err
  }

  // Update priority attribute
  config.Priority = append(config.Priority, repoName)

  // Marshal the updated Config struct back to JSON
  updatedConfig, err := json.MarshalIndent(config, "", "    ")
  if err != nil {
    return err
  }

  // Write the updated JSON to the config file
  if err := os.WriteFile(configPath, updatedConfig, 0644); err != nil {
    return err
  }

  return nil
}



