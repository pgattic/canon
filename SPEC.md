
# Canon Toolset [Draft]

Canon - The extensible book referencer

NOTE: A lot of this is just brainstorming. While a functioning implementation of `canon` has been made to retrieve references, it is not yet feature-complete; most of what is specified here has no implementation.

## Modules

### Canon (This codebase)

This is the heart of the toolset, and it provides the tooling necessary for managing and reading text.

Provides: `canon`

- Responsible for retrieving text given one reference as an input, i.e. `canon "Matthew 5:14-16"`.
- Can manage "texts", utilizing Git to keep them up-to-date
  - Stores the text "packages" in `~/.canon/texts/[Package Name]/` (see the [.canon/texts](#canontexts) section for more info)
- Can output the file location where it found the chapter first (relative to canon's data directory), followed by the verses in comma-separated fashion, so tools like Canon-Mark can identify the source material for the output text and perform more manipulations on it.

### Canon-Mark (No working implementation)

The annotation manager for Canon.

Provides: `canonmk`

- Stores and retrieves the locations of canon highlights, annotations, and links
- Has some syntax to add or remove highlights, annotations, and links, whereupon it modifies the annotation library as stored in the user's `~/.canon/marks/` directory (in raw text)

### Canon-Study (No working implementation)

A TUI (Terminal UI) study suite that manages canon text packages and canonmk data, utilizing the features of canon and canonmk

Provides: `canonstud`

## Standards

### .canon/texts

Installing a text is to be as simple as running `git clone --depth=1 [git repo]` in the `~/.canon/texts` directory

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

Each book is a directory as well, containing one ".md" file for each chapter, named after the chapter name; such as `12.md` for chapter 12. Canon's referencing expects the verse numbers to line up with the line numbers of the text documents. 

(NOTE TO SELF: Equating line numbers to verse numbers may be a bad idea for things that aren't verse-based. Consider other potential avenues. Also, calling them "Markdown" files may be inaccurate, as `canon` may never support most features of Markdown syntax.)

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

This is meant to help streamline searching for verses from the commandline. With this configuration, commands like `canon "1 John 4:19-20"` can be abbreviated to the much more succinct `canon 1jhn4:19-20`. Note that Canon's use of these aliases is case-insensitive. For the study tools, this also provides the order in which the books are meant to appear.

Creating a canon text is as simple as organizing it in the structure defined here, and hosting it as a public Git repo. Canon is expected to be able to organize the installation and removal of texts by utilizing git, and assigning them a location in the "priority" attribute of the top-level "config.json".

