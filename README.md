# BError - Error Utilities Implementation in Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bearaujus/berror/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bearaujus/berror)](https://goreportcard.com/report/github.com/bearaujus/berror)

This package introduce for advanced error handling, allowing developers to create structured, 
formatted, and traceable errors with minimal effort. It supports features like custom error definitions, 
stack traces, and flexible formatting to make error management clean and robust.

## Installation

To install BError, you can run the following command:

```shell
go get github.com/bearaujus/berror
```

## Import

```go
import "github.com/bearaujus/berror"
```

## Error Wrapper

### 1. Basic Workflow
> Create a new `Error Definition` -> Create `Wrapped Error` *error interface 

### 2. Example Basic `Error Definition`

```go
var (
	ErrStartProcess = berror.NewErrDefinition("Failed to start process: %v")
	ErrReadFile     = berror.NewErrDefinition("Failed to read file: %v")
	ErrOpenFile     = berror.NewErrDefinition("Failed to open file: %v")
)
```

### 3. Example `Error Definitions` With `Options`

```go
var (
    ErrWithCode = berror.NewErrDefinition("Error with code: %v",
        berror.OptionErrDefinitionWithErrCode("E1001"))
    
    ErrNoStackTrace = berror.NewErrDefinition("No stack trace: %v",
        berror.OptionErrDefinitionWithErrCode("E1002"),
        berror.OptionErrDefinitionWithDisabledStackTrace())
    
    ErrCustomFormat = berror.NewErrDefinition("Custom format: %v",
        berror.OptionErrDefinitionWithErrCode("E1003"),
        berror.OptionErrDefinitionWithCustomFormater(func(err, code, stack string) string {
            return fmt.Sprintf("Custom: %v | Code: %v | Trace: %v", err, code, stack)
        }))
    
    ErrJSONFormat = berror.NewErrDefinition("JSON format: %v",
        berror.OptionErrDefinitionWithErrCode("E1004"),
        berror.OptionErrDefinitionWithCustomFormater(berror.ErrWrapperFormatterJSON))
)
```

### 4. Example Usage

```go
func main() {
	// Example 1: Basic error creation
	fmt.Println("Basic Errors:")
	fmt.Println(ErrStartProcess.New("Invalid input"))
	fmt.Println(ErrStartProcess.New(os.ErrClosed))
	fmt.Println()

	// Example 2: Error definitions with options
	fmt.Println("Error Definitions with Options:")
	fmt.Println(ErrWithCode.New("Resource not found"))
	fmt.Println(ErrNoStackTrace.New("Operation timed out"))
	fmt.Println(ErrCustomFormat.New("Custom formatting example"))
	fmt.Println(ErrJSONFormat.New("JSON formatted error"))
	fmt.Println()

	// Example 3: Nested errors
	err := StartProcess("invalid-path")
	fmt.Println("Nested Errors:")
	fmt.Println(err)
	fmt.Println()

	// Example 4: Error comparison
	fmt.Println("Error Comparisons:")
	fmt.Printf("Matches os.ErrNotExist: %v\n", errors.Is(err, os.ErrNotExist))
	fmt.Printf("Matches os.ErrClosed: %v\n", errors.Is(err, os.ErrClosed))
	fmt.Printf("Matches ErrStartProcess: %v\n", ErrStartProcess.Is(err))
	fmt.Printf("Matches ErrOpenFile: %v\n", ErrOpenFile.Is(err))
	fmt.Printf("Matches ErrWithCode: %v\n", ErrWithCode.Is(err))
}

// StartProcess simulates a process that involves reading a file.
func StartProcess(path string) error {
	data, err := ReadFile(path)
	if err != nil {
		return ErrStartProcess.New(err)
	}
	_ = data // Do something with the data
	return nil
}

// ReadFile simulates reading a file and propagates errors from OpenFile.
func ReadFile(path string) ([]byte, error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, ErrReadFile.New(err)
	}
	defer file.Close()
	return io.ReadAll(file)
}

// OpenFile simulates opening a file and wraps any errors encountered.
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, ErrOpenFile.New(err)
	}
	return file, nil
}
```
````shell
Basic Errors:
Failed to start process: Invalid input (./main.go:40)
Failed to start process: file already closed (./main.go:41)

Error Definitions with Options:
[E1001] Error with code: Resource not found (./main.go:46)
[E1002] No stack trace: Operation timed out
Custom: Custom format: Custom formatting example | Code: E1003 | Trace: ./main.go:48
{"code":"E1004","err":"JSON format: JSON formatted error","stack":"./main.go:49"}

Nested Errors:
Failed to start process: Failed to read file: Failed to open file: open invalid-path: The system cannot find the file specified. (./main.go:91) (./main.go:81) (./main.go:71)

Error Comparisons:
Matches os.ErrNotExist: true
Matches os.ErrClosed: false
Matches ErrStartProcess: true
Matches ErrOpenFile: true
Matches ErrWithCode: false
````

### 5. Additional Notes
When you run your Go programs, the full file path of the source code is often included in the error stack traces. 
For example, without using any additional flags, you might see something like this:

```shell
go run .\examples\err_wrapper\main.go

Error: Resource not found (D:/Documents/GitHub/berror/examples/err_wrapper/main.go:46)
```

This output shows the full local path (D:/Documents/GitHub/berror/examples/err_wrapper/main.go) where the error occurred. 
While this can be useful for debugging during development, it might not be desirable for production logs, 
as it can reveal the structure of your file system or clutter the output.

The `-trimpath` flag in Go removes local file paths from the compiled output, 
leaving only the relative path in error stack traces. This makes the stack trace cleaner and more suitable for 
sharing or production use. Here's how you can use it:

```shell
go run -trimpath .\examples\err_wrapper\main.go

Error: Resource not found (./main.go:46)
```

Or for go build:

```shell
go build -trimpath -o myapp .\examples\err_wrapper\main.go
```

Now, instead of the full path (D:/Documents/GitHub/berror/examples/err_wrapper/main.go),
you only see a relative path (./main.go). This makes the error message cleaner and easier to read.

## License

This project is licensed under the MIT License - see
the [LICENSE](https://github.com/bearaujus/berror/blob/master/LICENSE) file for details.
