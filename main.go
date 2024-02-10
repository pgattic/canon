package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
  "encoding/json"
)

type Priority struct {
  Priority []string `json:"priority"`
}

type Aliases struct {
	Aliases map[string][]string `json:"aliases"`
}

func resolve_book(input_book string, canon_dir string) (canon string, book string) {
  var priority Priority
  // Open the JSON file
  config_f, err := os.Open(canon_dir + "/texts/config.json")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer config_f.Close()

  // Decode the JSON file into the struct
  err = json.NewDecoder(config_f).Decode(&priority)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  for _, canon := range priority.Priority {
    var aliases Aliases

    // Open the JSON file
    file, err := os.Open(canon_dir + "/texts/" + canon + "/config.json")
    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
    }
    defer file.Close()

    // Decode the JSON file into the struct
    err = json.NewDecoder(file).Decode(&aliases)
    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
    }

    for book, aliases := range aliases.Aliases {
      if strings.ToLower(book) == strings.ToLower(input_book) {
        return canon, book
      }
      for _, alias := range aliases {
        if strings.ToLower(alias) == strings.ToLower(input_book) {
          return canon, book
        }
      }
    }   
  }

  fmt.Println("Error: Unresolved book name: "+input_book)
  os.Exit(1)
  return
}

func print_entire_book(canon string, book string) {

}

func print_entire_canon(canon string) {

}

func print_ref(reference string) {
  book := reference
  var rest string
  for strings.Contains("1234567890 :-,;", book[len(book)-1:]) {
    rest = string(book[len(book)-1]) + rest
    book = book[:len(book)-1]
  }
  rest = strings.TrimSpace(rest)

  canon_dir, err := os.UserHomeDir()

  if err != nil {
    panic(err)
  }

  canon_dir += "/.canon"

  canon, book := resolve_book(book, canon_dir)

  if book == "" {
    print_entire_canon(canon)
    os.Exit(0)
  }

  if rest == "" { // no chapters or verses mentioned
    print_entire_book(canon, book)
    os.Exit(0)
  }

  refs := strings.Split(rest, ";")

  for r := 0; r < len(refs); r++ {
    ref := strings.TrimSpace(refs[r])
    if strings.Contains(ref, ":") { // if specific verse(s) referenced
      split := strings.Split(ref, ":")
      if len(split) > 2 {
        fmt.Println("Syntax Error: Too many colons")
        fmt.Println("  in \""+ref+"\"")
        os.Exit(1)
      }
      chapter := split[0]

      fs_ref := canon_dir + "/texts/" + canon + "/" + book + "/" + chapter + ".md"
      dat, err := os.ReadFile(fs_ref)
      if err != nil {
        panic(err)
      }

      verse_ranges := strings.Split(split[1], ",")

      for vr := 0; vr < len(verse_ranges); vr++ {
        verse_range := strings.TrimSpace(verse_ranges[vr])
        if strings.Contains(verse_range, "-") {
          ref_split := strings.Split(verse_range, "-")
          if len(ref_split) > 2 {
            fmt.Println("Syntax Error: Too many hyphens")
            fmt.Println("  in \""+verse_range+"\"")
            os.Exit(1)
          }
          start_verse, err_1 := strconv.Atoi(strings.TrimSpace(ref_split[0]))
          end_verse, err_2 := strconv.Atoi(strings.TrimSpace(ref_split[1]))

          if err_1 != nil || err_2 != nil {
            fmt.Println("Syntax Error: Unresolvable verse identifier")
            fmt.Println("  in \""+verse_range+"\"")
            return
          }

          chap := strings.Split(string(dat), "\n")

          result := strings.Join(chap[start_verse-1:end_verse], "\n")

          fmt.Println(result)

        } else {
          verse, err := strconv.Atoi(strings.TrimSpace(verse_range))

          if err != nil {
            fmt.Println("Syntax Error: Unresolvable verse identifier")
            fmt.Println("  in \""+verse_range+"\"")
            return
          }

          chap := strings.Split(string(dat), "\n")

          result := chap[verse-1]

          fmt.Println(result)
        }
      }

    } else { // no verse referenced, print the entire chapter
      fs_ref := canon_dir + "/texts/" + canon + "/" + book + "/" + ref + ".md"
      dat, err := os.ReadFile(fs_ref)
      if err != nil {
        panic(err)
      }
      fmt.Println(string(dat))
    }
  }
}

func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please specify an output.")
    return
  }

  for i := 1; i < len(args); i++ {
    print_ref(args[i])
  }

}


