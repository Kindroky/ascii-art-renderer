# ASCII Art Renderer

This project is a command-line tool written in Go that generates ASCII art based on a given string and template. It reads a specified banner template file and converts the input string into stylized ASCII art. The supported templates include `standard`, `shadow`, and `thinkertoy`.

## Features

- Converts any string into ASCII art using the selected template.
- Supports multi-line input with `\n` recognized as a line break.
- Includes three predefined templates: `standard`, `shadow`, and `thinkertoy`.
- Efficient handling of file input and rune-to-template mapping.

## Usage

To use this program, you must provide:
1. The string to convert.
2. The name of the template (`standard`, `shadow`, or `thinkertoy`).
To run the program, use the following command in your terminal:

```bash
go run ascii-art.go "[your desired text]" [template: standard, shadow, or thinkertoy]
```

### Example

Run the following command in your terminal:

```bash
go run . "Hello, World!" standard
