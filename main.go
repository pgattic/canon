package main

import (
  "fmt"
  "os"
  "github.com/pgattic/canon/config"
  "github.com/pgattic/canon/manager"
  "github.com/pgattic/canon/referencer"
)

func main() {
  config.EnsureSetup()
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please specify an output.")
    return
  }

  
  switch args[1] {

  /* Package management arguments */
  case "install":
    if len(args) < 4 {
      fmt.Println("Please specify a repo and a dirname (example: \"canon install https://github.com/user/repo Repo\")")
      return
    }
    manager.Install(args[2], args[3])
    return
  case "remove":
    if len(args) < 3 {
      fmt.Println("Please specify a package to remove")
      return
    }
    manager.Remove(args[2])
    return
  case "list":
    manager.List()
    return


  /* reference verse(s) */
  default:
    var execFlags referencer.Flags
    var refIdx int // index of the args that is the verse index (flags could be before or after the verse ref)
    for i := 1; i < len(args); i++ {
      if args[i][0] == '-' {
        if args[i][1] == '-' { // args starting with "--"
          switch args[i] {
          case "--paragraph":
            execFlags.Paragraph = true
          case "--verbose":
            execFlags.Verbose = true
          case "--numbered":
            execFlags.VerseNumbers = true
          }
          continue
        }
        for ch := 1; ch < len(args[i]); ch++ {
          switch args[i][ch] { // Execution flags
          case 'p':
            execFlags.Paragraph = true
          case 'v':
            execFlags.Verbose = true
          case 'n':
            execFlags.VerseNumbers = true
          }
        }
      } else {
        refIdx = i
      }
    }
    referencer.ParseRef(args[refIdx], execFlags)
  }
}


