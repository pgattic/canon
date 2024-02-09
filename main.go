package main

import (
	"fmt"
	"os"

	//"regexp"
	"strings"
  "strconv"
)

func parse_reference(reference string) (book string, chapter string, start_verse int, end_verse int) {
  splat := strings.Split(reference, ":")
  book = splat[0]
  ref_str := splat[1]

  for strings.Contains("1234567890", book[len(book)-1:]) { // Separate the chapter number from book name
    chapter = string(book[len(book)-1]) + chapter
    book = book[:len(book)-1]
  }
  book = strings.TrimSpace(book)

  var end_ref_str string
  if strings.Contains(ref_str, "-") {
    split := strings.Split(ref_str, "-")
    ref_str = split[0]
    end_ref_str = split[1]
  }

  start_verse, err_1 := strconv.Atoi(ref_str)

  if err_1 != nil {
    fmt.Println("Dang it")
    return
  }

  var err_2 error

  if end_ref_str != "" {
    end_verse, err_2 = strconv.Atoi(end_ref_str)
    if err_2 != nil {
      fmt.Println("Dang it")
    }
  } else {
    end_verse = start_verse
  }
  return book, chapter, start_verse, end_verse
}

func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please specify an output.")
    return
  }

  var reference string = args[1]

  book, chapter, start_verse, end_verse := parse_reference(reference)

  fs_ref := "texts/Book of Mormon/" + book + "/" + chapter + ".md"

  dat, err := os.ReadFile(fs_ref)

  if err != nil {
    panic(err)
  }

  chap := strings.Split(string(dat), "\n")

  result := strings.Join(chap[start_verse-1:end_verse], "\n")

  fmt.Println(result)
}


