package referencer

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strconv"
  "strings"
  "github.com/pgattic/canon/config"
)

type Flags struct { // command-line flags
  Paragraph bool // -p
  Verbose bool // -v
  VerseNumbers bool // -n
}

type Aliases struct {
  Aliases map[string][]string `json:"aliases"`
}

var execFlags Flags // global since it is referenced all over the place, set in ParseRef()

func resolveBook(input_book string) (path string) {
  priority := config.LoadConfig()

  for _, canon := range priority.Priority {
    var aliases Aliases

    // Open the canon's config file
    file, err := os.ReadFile(filepath.Join(config.TextsDir, canon, "config.json"))
    if err != nil {
      panic(err)
    }

    // Unmarshal data into Aliases struct
    err_1 := json.Unmarshal(file, &aliases)
    if err_1 != nil {
      panic(err_1)
    }

    for book, aliases := range aliases.Aliases {
      if strings.ToLower(book) == strings.ToLower(input_book) {
        return filepath.Join(canon, book)
      }
      for _, alias := range aliases {
        if strings.ToLower(alias) == strings.ToLower(input_book) {
          return filepath.Join(canon, book)
        }
      }
    }   
  }

  fmt.Println("Error: Unresolved command or book name:", input_book)
  os.Exit(1)
  return
}

func printEntireCanon(canon string) {
  // TODO
}

func printEntireBook(path string) {
  // TODO
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
  if execFlags.Verbose {
    fmt.Print("@@@" + strconv.Itoa(verse) + " ")
  }
  if execFlags.VerseNumbers {
    fmt.Println("", verse, sourceContent[verse-1])
  } else {
    fmt.Println(sourceContent[verse-1])
  }
  if execFlags.Paragraph {
    fmt.Println()
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

  bookPath := resolveBook(book) // Locate what book "John" is in

  if execFlags.Verbose {
    fmt.Println("@" + bookPath + "/")
  }

  if rest == "" { // if no chapters or verses mentioned
    printEntireBook(bookPath)
    return
  }

  refs := strings.Split(rest, ";") // "John" "3:5,16-17; 14:15" -> "John" ["3:5,16-17" "14:15"]

  for r := 0; r < len(refs); r++ {
    ref := strings.TrimSpace(refs[r])
    split := strings.Split(ref, ":")
    chapter := strings.TrimSpace(split[0])
    fs_ref := filepath.Join(config.TextsDir, bookPath, chapter)
    dat, err := os.ReadFile(fs_ref)
    if err != nil {
      fmt.Println("Error: Chapter not found")
      fmt.Println("  "+fs_ref)
      return
    }
    if execFlags.Verbose {
      fmt.Println("@@" + chapter)
    }
    chap := strings.Split(strings.TrimSpace(string(dat)), "\n") // Entire chapter as a slice of strings

    if !strings.Contains(ref, ":") {
      printChapter(chap)
      return
    }

    verse_ranges := strings.Split(split[1], ",") // "3" "5,16-17" -> "3" ["5", "16-17"]

    for vr := 0; vr < len(verse_ranges); vr++ {
      verse_range := strings.TrimSpace(verse_ranges[vr])

      if strings.Contains(verse_range, "-") {
        ref_split := strings.Split(verse_range, "-")
        start_verse, err_1 := strconv.Atoi(strings.TrimSpace(ref_split[0]))
        end_verse, err_2 := strconv.Atoi(strings.TrimSpace(ref_split[1]))

        if start_verse < 1 {start_verse = 1}
        if end_verse > len(chap) {end_verse = len(chap)}

        if err_1 != nil || err_2 != nil {
          fmt.Println("Syntax Error: Unresolvable verse identifier")
          fmt.Println("  in \""+verse_range+"\"")
          os.Exit(1)
        }
        printVerseRange(start_verse, end_verse, chap)
      } else {
        verse, err := strconv.Atoi(strings.TrimSpace(verse_range))

        if err != nil {
          fmt.Println("Syntax Error: Unresolvable verse identifier")
          fmt.Println("  in \""+verse_range+"\"")
          os.Exit(1)
        }
        if verse >= 1 && verse <= len(chap) {
          printVerse(verse, chap)
        }
      }
    }
  }
}

