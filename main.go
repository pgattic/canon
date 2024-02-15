package main

import (
  "fmt"
  "os"
  "github.com/pgattic/canon/manager"
  "github.com/pgattic/canon/referencer"
)

func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please specify an output.")
    return
  }

  if args[1] == "--install" || args[1] == "-i" {
    if len(args) < 4 {
      fmt.Println("Please specify a repo and a dirname (example: \"canon install https://github.com/user/repo Repo\")")
    }
    manager.GitClone(args[2], args[3])
  } else {
    for i := 1; i < len(args); i++ {
      referencer.ParseRef(args[i])
    }
  }
}


