package main

import (
  "fmt"
  "os"
  "github.com/pgattic/canon/config"
  "github.com/pgattic/canon/manager"
  "github.com/pgattic/canon/referencer"
)

func displayHelp() {
  fmt.Println("Usage:", os.Args[0], "[COMMAND/REFERENCE] [OPTS]")
  fmt.Println()
  fmt.Println("Canon Book Referencer")
  fmt.Println()
  fmt.Println("Commands:")
  fmt.Println("  help                  Display this page")
  fmt.Println("  install [Repo URL]    Install canon package from repository                   ")
  fmt.Println("  list                  List installed Canon packages")
  fmt.Println("  remove [Package]      Remove package by shortname")
  fmt.Println()
  fmt.Println("If anything besides these commands is given, the input is assumed to be a book")
  fmt.Println("reference, followed or preceded by any combination of these reference options:")
  fmt.Println()
  fmt.Println("Reference Options")
  fmt.Println("  -n, --numbered        Print verse/paragraph numbers before each line.")
  fmt.Println("  -p, --paragraph       Print the lines of text with an extra space in between")
  fmt.Println("                        them, as paragraphs.")
  fmt.Println("  -v, --verbose         Print extra information, like where the book was found.")
  fmt.Println("                        Useful for supplementary tools like canonmk.")
  fmt.Println()
  fmt.Println("Canon made by Preston Corless (pgattic), free under the MIT License.")
  fmt.Println("More information can be found at https://github.com/pgattic/canon")
}

func main() {
  config.EnsureSetup()
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Usage:", args[0], "[COMMAND/REFERENCE] [OPTS]")
    fmt.Println("Try \"" + args[0] + " help\" for more information.")
    return
  }

  
  switch args[1] {

  /* Package management arguments */
  case "install":
    if len(args) < 3 {
      fmt.Println("Please specify a canon-formatted repository (example: \"canon install https://github.com/user/repo\")")
      return
    }
    manager.Install(args[2])
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
  case "help":
    displayHelp()
    return


  /* reference verse(s) */
  default:
    var execFlags referencer.Flags
    var refIdx int // index of the args that is the verse index (flags could be before or after the verse ref)
    for i := len(args)-1; i >= 1; i-- {
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


