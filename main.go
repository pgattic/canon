package main

import (
	"fmt"
	"os"
	"strings"
  "strconv"
)


/*

Note: Commas only separate verses within a chapter, semicolons are used for inter-chapter, inter-book, or inter-canon references.
Note: In theory, the hyphen could be used to span across chapters, right? Or even to say "1 Nephi 2-3", right? (I won't worry about that now.)
Note: There shall be no references made across the gap from one book to another. Distinct book references shall be proceeding command-line args.
The only exception will be when querying an entire canon of scripture, or a group of canons.

*/

type ChapReference struct {
  chapter string
  verses []int
}

type BookReference struct {
  book string
  chapters []ChapReference
}

type CanonReference struct {
  canon string
  books []BookReference
}

type Reference struct {
  canons []CanonReference
}

func locate_book(input string) (canon string, book string, locate_err bool) { // TODO: Actually search the filesystem according to book resolving specs
  canon = "Book of Mormon"
  book = input
  locate_err = false
  return
}

func resolve_book(segment string) (canon string, book string, rest string, book_res_err bool) {
  for ch := len(segment)-1; ch >= 0; ch-- { // NOTE: This loop can probably be rewritten as a regex
    if strings.Contains("1234567890-,[]: ", string(segment[ch])) {
      rest = string(segment[ch]) + rest
      segment = segment[:ch] + segment[ch+1:]
    } else {
      break
    }
  }
  segment = strings.TrimSpace(segment) // Probably optional
  rest = strings.TrimSpace(rest)
  // At this point, "1 Nephi 3:7-9" has become "1 Nephi" and "3:7-9" in vars segment and rest
  var locate_err bool
  canon, book, locate_err = locate_book(segment)
  book_res_err = locate_err
  return
}

func (r Reference) Parse(reference_str string) {
  subrefs := strings.Split(reference_str, ";") // on Gospel Library, these subrefs would each show up as separate hyperlinks
  reference := Reference{}
  var ctx_canon string
  var ctx_book string
  for i := 0; i < len(subrefs); i++ {
    canon, book, rest, book_res_err := resolve_book(subrefs[i])
    if book == "" {
      book = ctx_book;
    } else {
      ctx_book = book;
    }
    if canon == "" {
      canon = ctx_canon;
    } else {
      ctx_canon = canon;
    }
    if book_res_err {
      fmt.Println("Error: Unable to locate reference")
      return
    }
    reference.canons = append(reference.canons, )
  }
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


