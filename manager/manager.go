package manager // haha get it? lol

import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
)

type Config struct {
  Priority []string `json:"priority"`
}

func loadConfig(configPath string) (Config, error) {
  var config Config

  file, err := os.ReadFile(configPath)
  if err != nil {
    return config, err
  }

  // Unmarshal JSON into Config struct
  if err := json.Unmarshal(file, &config); err != nil {
    return config, err
  }
  return config, err
}

func saveConfig(config Config, configPath string) error {
  // Marshal the updated Config struct back to JSON
  updatedConfig, err := json.MarshalIndent(config, "", "    ")
  if err != nil {
    return err
  }

  if err := os.WriteFile(configPath, updatedConfig, 0644); err != nil {
    return err
  }
  return nil
}

func cloneRepo(repoURL, destination string) error {
  cmd := exec.Command("git", "clone", repoURL, destination, "--depth", "1")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func updateConfig(configPath, repoName string) error {
  config, err := loadConfig(configPath)
  if err != nil {
    return err
  }

  // Update priority attribute (put it at the first spot)
  config.Priority = append([]string{repoName}, config.Priority...)

  if err := saveConfig(config, configPath); err != nil {
    return err
  }
  return nil
}

func Install(repoURL string, repoDir string) {
  textsDir := filepath.Join(os.Getenv("HOME"), ".canon", "texts")

  // Ensure the texts directory exists
  if err := os.MkdirAll(textsDir, 0755); err != nil {
    fmt.Printf("Error creating texts directory: %v\n", err)
    return
  }

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

func Remove(repoDir string) {
  textsDir := filepath.Join(os.Getenv("HOME"), ".canon", "texts")

  // Ensure the texts directory exists
  if err := os.MkdirAll(textsDir, 0755); err != nil {
    fmt.Printf("Error creating texts directory: %v\n", err)
    return
  }

  os.RemoveAll(filepath.Join(textsDir, repoDir))
  configPath := filepath.Join(textsDir, "config.json")

  config, err := loadConfig(configPath)
  if err != nil {
    fmt.Println("Failed to update config.json.")
  }
  var newPriority []string
  for i := 0; i < len(config.Priority); i++ {
    if config.Priority[i] != repoDir {
      newPriority = append(newPriority, config.Priority[i])
    }
  }
  if len(config.Priority) == len(newPriority) {
    fmt.Println("Package not found in config.json")
  }
  config.Priority = newPriority
  if err := saveConfig(config, configPath); err != nil {
    panic(err)
  }
}

func List() {
  textsDir := filepath.Join(os.Getenv("HOME"), ".canon", "texts")

  // Ensure the texts directory exists
  if err := os.MkdirAll(textsDir, 0755); err != nil {
    fmt.Printf("Error creating texts directory: %v\n", err)
    return
  }

  config, err := loadConfig(filepath.Join(textsDir, "config.json"))
  if err != nil {
    fmt.Println("Failed to read config.json.")
  }
  fmt.Print("\033[36;1m") // Pretty blue color!!1!
  for i := 0; i < len(config.Priority); i++ {
    fmt.Println(config.Priority[i])
  }
}

