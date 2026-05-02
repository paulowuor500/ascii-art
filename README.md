# ASCII-Art

## Overview

This project is a **Go-based ASCII art generator** that converts text input into stylized banners using predefined font files. It supports multiple banner styles and handles escape sequences like `\n` for multi-line output.

---

## Features

* Convert plain text into ASCII art
* Multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Handles newline characters (`\n`)
* Clean modular structure with internal packages
* Unit tests included

---

## Example

### Input

go run . "Hello\n" | cat -e
```

### Output


 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

---

## Project Structure

```
ascii-art/
├── main.go
├── go.mod
├── README.md
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── internal/
    ├── parser.go
    ├── printer.go
    └── printer_test.go
```

## How It Works

### 1. Parsing Input

The parser (`internal/parser.go`) processes:

* Raw string input
* Escape sequences like `\n`
* Splits input into lines and characters

### 2. Banner Files

Each banner file contains ASCII representations of characters using a fixed height format. The program maps each input character to its corresponding ASCII block.

### 3. Printing Engine

The printer (`internal/printer.go`):

* Aligns characters line by line
* Combines them into full words
* Outputs formatted ASCII art to stdout

---

### Clone the Repository

git clone https://learn.zone01kisumu.ke/git/paowuor/ascii-art
cd ascii-art
```

---

## Usage

### Basic Usage

go run . "Hello"
```

### With New Lines

go run . "Hello\nWorld"
```

### Specify Banner Style

go run . "Hello" standard
go run . "Hello" shadow
go run . "Hello" thinkertoy
```
---

## Testing

Run unit tests:

go test ./...
```

---

## Error Handling

The program handles:

* Invalid input characters
* Missing or corrupted banner files
* Incorrect arguments

---

## Customization

You can add new fonts by:

1. Creating a new `.txt` file in the `banner/` directory
2. Following the same character height/format convention
3. Updating the program to recognize the new banner

---

## Contributing

Contributions are welcome:

* Improve performance
* Add new banner styles
* Enhance test coverage
