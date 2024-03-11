# Code Challange 

# Logger in Go

Defines a simple logging interface named Logger with three methods: Info, Warning, and Error. It also provides an implementation of this interface using a Log struct, which has a field for an os.File to write logs to.

## Logger Interface

The Logger interface has the following methods:

- `Info(message string)`: Prefixes the message with "[info]".
- `Warning(message string)`: Prefixes the message with "[warning]".
- `Error(message string)`: Prefixes the message with "[error]".

## Log Struct

The Log struct implements the Logger interface and has a field named `file` of type *os.File representing the log file.

## NewLog Function

The package provides a function named `NewLog` that creates a new `Log` instance with the specified filename.

## Solution
See the provided Go program in the file `logger.go.`



# Key-Value Map Populator

Define Go program that reads a text file containing data in the format "key:value" on each line. It parses the file and populates a map with the key-value pairs. It also handles any errors that may occur during file reading or parsing.

## Usage

1. Make sure you have Go installed on your system.
2. Clone this repository or copy the contents of `main.go` into a new Go file.
3. Modify the `filePath` variable in the code to point to your input text file.

 ## solution
See the provided Go program in the file `parser.go.`
