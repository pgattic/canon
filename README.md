
# Canon - An Extensible Book Referencer

[Software Demo Video](https://youtu.be/5VAD_pyJUzk)

## Purpose

As a person of religion, I frequently involve myself in the study of written scripture. Being also a proponent of the [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy), I felt inspired to make a program that facilitates my study of the scriptures, and also tries to keep in line with computer software best practices. I also thought it would be a fun exercise into learning Golang!

## Disclaimer

To avoid confusion, I would like to state that though I have my opinions, this program is **not theologically or religiously opinionated**. As a matter of fact, it has potential use cases outside of religion. It comes with no texts pre-installed, and it makes no assumptions about the subject of the text it is querying. It could work with any type of text, as long as it is packaged in the way specified by the SPEC.md document.

Although I provide the packages for the King James Version New Testament, King James Version Old Testament, Book of Mormon, LDS Doctrine and Covenants, and LDS Pearl of Great Price in their own separate repositories, they are not intended to be the only packages used by this program, and they are not required for its functionality. I encourage you to create packages for any books you may use.

## Usage

In order to install this program, you must first have the Go compiler installed.

1. `git clone https://github.com/pgattic/canon`
2. `cd canon`
3. `make`
4. `sudo make install`

Now, use canon's builtin `install` command to install and index a repository package. Pass in the repo URL, followed by a shortname to use for the directory. For example, this is how I would install my King James Version New Testament package:

`canon install https://github.com/pgattic/nt-kjv-canon nt`

References to a paragraph/verse are made with the following syntax: `"[BookName] [Chapter][:[Verse Ranges]]; [Chapter][:[Verse Ranges]]; ..."`, for example, referencing the Book of John, in the 5th, 16th, and 17 verses of chapter 3, and the 15th verse of chapter 14 would be done with `canon "John 3:5,16-17; 14:15"`.

## Development Environment

In the process of developing this software, I used:

- Neovim (Code editor)
- Golang (Programming langauge)

## Useful Websites

### Golang

- [Go Cheatsheet](https://devhints.io/go)
- [Tour of Go](https://go.dev/tour)

### Other

In creating this program, I wanted its usage syntax to be one that the most people would be familiar with, regardless of their background. Although it would be impossible to make it cover all potential shorthands for referencing text, here are some sources that I used for inspiration on the syntax of the references.

- [Bible Citation - Wikipedia](https://en.wikipedia.org/wiki/Bible_citation)
- [How to Cite the Bible - Messiah.edu](https://www.messiah.edu/download/downloads/id/1647/bible_cite.pdf)
- [The Quran: Word by Word](https://corpus.quran.com/wordbyword.jsp)
- [Preach My Gospel: Chapter 3, Lesson 4](https://www.churchofjesuschrist.org/study/manual/preach-my-gospel-2023/04-chapter-3/11-chapter-3-lesson-4?lang=eng)

## Progress

This program is currently in a very functional state; however, it still has some progress to make in terms error handling and user comfort. These are things that I hope to improve upon before I feel comfortable listing this as an AUR package or Copr repo.

