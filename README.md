## Decode Library

The Decode Library is a Go library that provides a collection of functions for converting different types of codes to text and vice versa. It supports conversion between Morse Code, Alphabet, Binary, Numbers, and more, making it easier to communicate and manipulate data in different formats.

## Installation

To use the Decode Code Conversion Library in your Go project, follow the steps below:

1. Make sure you have Go installed on your system. For installation instructions, refer to the [official Go website](https://golang.org/doc/install).

2. Create a new directory for your project.

3. In the project directory, run the following command to initialize a new Go module:

   ```shell
   go mod init module-name
   ```

4. Download the decode library using the `go get` command:

   ```shell
   go get github.com/emanuelvsz/decode
   ```

## Usage

After installation, you can import the library package into your Go project and use the available conversion functions.

```go
import "github.com/emanuelvsz/decode"
```

### Morse Code to Alphabet Conversion

```go
text := decode.MorseToAlphabet("... --- ...")
fmt.Println(text) // Output: "SOS"
```

### Alphabet to Morse Code Conversion

```go
morse := decode.AlphabetToMorse("HELLO")
fmt.Println(morse) // Output: ".... . .-.. .-.. ---"
```

### Binary to Numbers Conversion

```go
number := decode.BinaryToNumbers("101010")
fmt.Println(number) // Output: 42
```

### Numbers to Binary Conversion

```go
binary := decode.NumbersToBinary(42)
fmt.Println(binary) // Output: "101010"
```

## Adding New Functionality

The Decode Code Conversion Library is designed to be easily extensible, allowing the addition of new conversion functionalities. To add a new functionality, follow the steps below:

1. Open the `decode_library.go` file in a text editor.

2. Define a new function for the desired conversion, following the same pattern as the existing functions in the library. Make sure to provide a proper description, input parameters, and return type.

3. Save the file with the new function added.

You can now import the library package again and use the new functionality in your project.

## Final Remarks

The Decode Code Conversion Library in Go provides a convenient way to convert different types of codes to text and vice versa. The detailed documentation helps understand how to use the existing functions and also allows for the expansion of the project with new functionalities. Feel free to contribute to the library by adding new conversions and improvements.

If you need further information or additional support, please don't hesitate to reach out to the development team.

Feel free to modify, edit, or delete this documentation as you see fit.
