
# Canon Toolset [Draft]

Canon - The extensible book referencer

## Modules

### Canon

This is the heart of the toolset, and it provides the tooling necessary for managing and reading text.

Provides: `canon`

- Responsible for retrieving text given one reference as an input
- Can manage "texts", utilizing Git to keep them up-to-date
  - Stores the text "packages" in ~/.canon/texts/[Package Name]/
- Can also take highlights as an input (using --hl="[Highlight Data]") and render the text according to the highlight. [TODO: Determine a standard for both storing and passing the highlight data]

### Canon-Mark

The annotation manager for Canon.

Provides: `canonmk`

- Stores and retrieves canon highlights, annotations, and links, stored in raw-text form
- Takes a scripture reference as an input, and outputs the same scripture reference along with its relevant highlight data, so as to be piped into `canon`. For example: `canonmk "Matthew 5:14-16"` might output something like `"Matthew 5:14-16" --hl="[Highlight Data]"` [NOTE: It may be better to only output the [Highlight Data]. Think about this?]
- Has some syntax to add or remove highlights, annotations, and links, whereupon it modifies the annotation library as stored in the user's ~/.canon/marks/ directory (in raw text)

### Canon-Study

A TUI (Terminal UI) study suite that manages canon text packages and canonmk data, utilizing the features of canon and canonmk

Provides: `canonstud`

## Standards

### .canon/texts

Installing a text is as simple as running `git clone --depth=1 [git repo]` in the `~/.canon/texts` directory

Here is an example of the structure of a ~/.canon/texts/ directory:

```
~/.canon/texts/
├── config.json
├── New Testament
│   ├── 1 Corinthians
│   ├── 1 John
│   ├── 1 Peter
...
│   ├── Revelation
│   ├── Romans
│   ├── Titus
│   └── config.json
└── Old Testament
    ├── 1 Chronicles
    ├── 1 Kings
    ├── 1 Samuel
...
    ├── Solomon's Song
    ├── Zechariah
    ├── Zephaniah
    └── config.json
```

Each book is a directory as well, containing one ".md" file for each chapter, named after the chapter name; such as `12.md` for chapter 12. The verse numbers are expected to line up with the line numbers of the text documents.

Note that each directory has a config.json file. The type of data stored in this file depends on where it is. [NOTE: In the future, this format may change, as JSON may not be the optimal format for this use case]

As of now, the config.json framework is meant to help `canon` locate the desired book of scripture.

Here is an example top-level config.json:

```
{
  "priority": [
    "Old Testament",
    "New Testament"
  ]
}
```

This specifies the order in which `canon` searches the children directories for matches on the input book (i.e. "Matthew"). 

This configuration is useful for example if the user has more than one translation of the Bible in their library, or there are other naming collisions, so that there is a consistent, discrete output. 

Also, each text is meant to contain its own config.json, which is used to specify alias names of the books contained therein. For example:

```
{
  "aliases": {
    "Matthew": [
      "Mat",
      "Mat."
    ],
    "Mark": [
      "Mar",
      "Mar.",
      "Mrk",
      "Mrk."
    ],
    "Luke": [
      "Luk",
      "Luk."
    ],
    "John": [
      "Jon",
      "Jon.",
      "Jhn",
      "Jhn."
    ],
    "Acts": [
      "Act",
      "Act."
    ],
...
    "3 John": [
      "3Jon",
      "3 Jon",
      "3Jon.",
      "3 Jon.",
      "3Jhn",
      "3 Jhn",
      "3Jhn.",
      "3 Jhn."
    ],
    "Jude": [
      "Jud",
      "Jud."
    ],
    "Revelation": [
      "Rev",
      "Rev."
    ]
  }
}
```

This is meant to help streamline searching for verses from the commandline. With this configuration, commands like `canon "1 John 4:19-20"` can be abbreviated to the much more succinct `canon 1jhn4:19-20`. Note that Canon's use of these aliases is case-insensitive.


