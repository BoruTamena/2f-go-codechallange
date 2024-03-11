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