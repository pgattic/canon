package main

import (
	"fmt"
	"io/fs"
	"os"

	//"regexp"
	"strings"
  "strconv"
)


func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please specify an output.")
    return
  }

  var reference string = args[1]

  splat := strings.Split(reference, ":")
  book := splat[0]
  ref_str := splat[1]
  var chapter string

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

//  fmt.Println(book)
//  fmt.Println(chapter)
//  fmt.Println(ref_str)
//  fmt.Println(end_ref_str)

  fs_ref := "texts/Book of Mormon/" + book + "/" + chapter + ".md"

  fmt.Println(fs_ref)

  if fs.ValidPath(fs_ref) { // this is useless and does nothing
    fmt.Println("yes")
  }

  dat, err := os.ReadFile(fs_ref)

  if err != nil {
    panic(err)
  }

  chap := strings.Split(string(dat), "\n")


  ref, err_1 := strconv.Atoi(ref_str)

  if err_1 != nil {
    fmt.Println("Dang it")
    return
  }

  var result string

  if end_ref_str != "" {
    fmt.Println("more than one verse")
    end_ref, err_2 := strconv.Atoi(end_ref_str)
    if err_2 != nil {
      fmt.Println("Dang it")
      return
    }
    result = strings.Join(chap[ref-1:end_ref], "\n")
  } else {
    result = chap[ref-1]
  }

  fmt.Println(result)

//  pattern := `^([^:]+)(\d+):(\d+)(?:-(\d+))?$`
//
//  re := regexp.MustCompile(pattern)
//  matches := re.FindStringSubmatch(input)
//
//  if len(matches) == 0 {
//    fmt.Println("No match found")
//    return
//  }
//
//  name := matches[1]
//  number1 := matches[2]
//  number2 := matches[3]
//  numberRange := matches[4]
//
//  fmt.Println("Name:", name)
//  fmt.Println("Number 1:", number1)
//  fmt.Println("Number 2:", number2)
//  if numberRange != "" {
//    fmt.Println("Number Range:", numberRange)
//  } else {
//    fmt.Println("No number range specified")
//  }
}


