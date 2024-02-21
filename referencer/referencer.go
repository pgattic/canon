package referencer

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "encoding/json"
)

type Flags struct { // command-line flags
  VerseNumbers bool // -n
  PrintPath bool // -p
}

type Priority struct {
  Priority []string `json:"priority"`
}

type Aliases struct {
	Aliases map[string][]string `json:"aliases"`
}

var execFlags Flags // global since it is referenced all over the place, set in ParseRef()

func resolveBook(input_book string, canon_dir string) (path string) {
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
        return canon + "/" + book
      }
      for _, alias := range aliases {
        if strings.ToLower(alias) == strings.ToLower(input_book) {
          return canon + "/" + book
        }
      }
    }   
  }

  fmt.Println("Error: Unresolved book name: "+input_book)
  os.Exit(1)
  return
}

func printEntireCanon(canon string) {

}

func printEntireBook(path string) {

}

func printChapter(chapter []string) {
  printVerseRange(1, len(chapter), chapter)
}

func printVerseRange(startVerse int, endVerse int, sourceContent []string) {
  for v := startVerse; v <= endVerse; v++ {
    printVerse(v, sourceContent)
  }
}

func printVerse(verse int, sourceContent []string) {
  if execFlags.PrintPath {
    fmt.Print(">")
  }
  if execFlags.VerseNumbers {
    fmt.Println(" " + strconv.Itoa(verse) + " " + sourceContent[verse-1])
  } else {
    fmt.Println(sourceContent[verse-1])
  }
}

func ParseRef(reference string, executionFlags Flags) { // Comments will follow the process of parsing "John 3:5,16-17; 14:15"
  execFlags = executionFlags
  book := reference
  rest := ""
  for strings.Contains("1234567890 :-,;", book[len(book)-1:]) { // "John 3:16-17" -> "John", "3:16-17"
    rest = string(book[len(book)-1]) + rest
    book = book[:len(book)-1]
  }
  rest = strings.TrimSpace(rest)

  canon_dir, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }
  canon_dir += "/.canon"

  bookPath := resolveBook(book, canon_dir) // Locate "John" (its canon is not intrinsic)

  if execFlags.PrintPath {
    fmt.Println("@" + bookPath + "/")
  }

  if rest == "" { // if no chapters or verses mentioned
    printEntireBook(bookPath)
    return
  }

  refs := strings.Split(rest, ";") // "John" "3:5,16-17; 14:15" -> "John" ["3:5,16-17" "14:15"] (NOTE: This feature subject to removal)

  for r := 0; r < len(refs); r++ {
    ref := strings.TrimSpace(refs[r])
    split := strings.Split(ref, ":")
    chapter := strings.TrimSpace(split[0])
    fs_ref := canon_dir + "/texts/" + bookPath + "/" + chapter
    dat, err := os.ReadFile(fs_ref)
    if err != nil {
      fmt.Println("Error: File not found")
      fmt.Println("  "+fs_ref)
      return
    }
    if execFlags.PrintPath {
      fmt.Println("@@" + chapter)
    }
    chap := strings.Split(strings.TrimSpace(string(dat)), "\n")

    if !strings.Contains(ref, ":") {
      printChapter(chap)
      return
    }

    verse_ranges := strings.Split(split[1], ",") // "3" "5,16-17" -> "3" ["5", "16-17"]

    for vr := 0; vr < len(verse_ranges); vr++ {
      verse_range := strings.TrimSpace(verse_ranges[vr])

      if execFlags.PrintPath {
        fmt.Println("@@@" + verse_range)
      }
      if strings.Contains(verse_range, "-") {
        ref_split := strings.Split(verse_range, "-")
        start_verse, err_1 := strconv.Atoi(strings.TrimSpace(ref_split[0]))
        end_verse, err_2 := strconv.Atoi(strings.TrimSpace(ref_split[1]))

        if err_1 != nil || err_2 != nil {
          fmt.Println("Syntax Error: Unresolvable verse identifier")
          fmt.Println("  in \""+verse_range+"\"")
          return
        }
        printVerseRange(start_verse, end_verse, chap)
      } else {
        verse, err := strconv.Atoi(strings.TrimSpace(verse_range))

        if err != nil {
          fmt.Println("Syntax Error: Unresolvable verse identifier")
          fmt.Println("  in \""+verse_range+"\"")
          return
        }
        printVerse(verse, chap)
      }
    }
  }
}

