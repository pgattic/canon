package manager // haha get it? lol

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  "github.com/pgattic/canon/config"
)

func cloneRepo(repoURL, destination string) error {
  cmd := exec.Command("git", "clone", repoURL, destination, "--depth", "1")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func updateConfig(repoName string) error {
  configFile := config.LoadConfig()

  // Update priority attribute (put it at the first spot)
  configFile.Priority = append([]string{repoName}, configFile.Priority...)

  config.SaveConfig(configFile)
  return nil
}

func Install(repoURL string, repoDir string) {

  // Clone the Git repository into the texts directory
  if err := cloneRepo(repoURL, config.TextsDir + "/" + repoDir); err != nil {
    fmt.Printf("Error cloning repository: %v\n", err)
    return
  }

  // Update the config.json with the repository name
  if err := updateConfig(repoDir); err != nil {
    fmt.Printf("Error updating config.json: %v\n", err)
    return
  }

  fmt.Println("Repository cloned successfully and config updated.")
}

func Remove(repoDir string) {
  configData := config.LoadConfig()

  var newPriority []string
  for i := 0; i < len(configData.Priority); i++ {
    if configData.Priority[i] != repoDir {
      newPriority = append(newPriority, configData.Priority[i])
    }
  }
  if len(configData.Priority) == len(newPriority) {
    fmt.Println("Package not found in config.json")
  } else {
    os.RemoveAll(filepath.Join(config.TextsDir, repoDir))
  }
  configData.Priority = newPriority
  config.SaveConfig(configData)
}

func List() {
  configList := config.LoadConfig().Priority

  fmt.Print("\033[36;1m") // Pretty blue color!!1!
  fmt.Println(strings.Join(configList[:], "\n"))
  fmt.Print("\033[0m")
}

