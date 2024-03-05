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
  cmd := exec.Command("git", "clone", repoURL, destination, "--depth", "1") // --depth 1 because the commit history is not useful here
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func updateConfig(repoName string) {
  configFile := config.LoadConfig()

  // Update priority attribute (put it at the first spot)
  configFile.Priority = append([]string{repoName}, configFile.Priority...)

  config.SaveConfig(configFile)
}

func getRepoName(url string) string {
  slashParts := strings.Split(url, "/")
  return strings.TrimSuffix(slashParts[len(slashParts)-1], ".git")
}

func Install(repoURL string) {
  repoName := getRepoName(repoURL)

  // Clone the Git repository into the texts directory
  if err := cloneRepo(repoURL, config.TextsDir + "/" + repoName); err != nil {
    fmt.Printf("Error cloning repository: %v\n", err)
    return
  }

  // Update the config.json with the repository name
  updateConfig(repoName)

  fmt.Println("Package installed successfully and config updated.")
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
  for i := 0; i < len(configList); i++ {
    fmt.Println(configList[i])
  }
  fmt.Print("\033[0m")
}

