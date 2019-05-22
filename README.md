# Base64Util
This is a GoLang package to assist base64 file encoding/decoding. 

## Installation
```
go get github.com/ryanjeon/Base64Util
```
# Example
    package main

    import (
        "fmt"
        "os"
        Base64Enc "github.com/ryanjeon/Base64Util"
    )

    func main() {
        file, err := os.Open("clown.png")
        if err != nil {
            fmt.Println("error")
        }

        Base64 := Base64Enc.FileToBase64Url(file)
        fmt.Println(Base64) \\Prints a Base64Url String 
    }

# Base64Enc Methods
## Encoding
- **FileToBase64(file *os.File)**: takes a pointer to a File 
and returns a base64 encoded string
- **FileToBase64Url(file *os.File)**: takes a pointer to a File and returns a base64url encoded string

## Decoding
- **Base64ToByte(code string)**: takes a base64/base64url encoded string, and returns a corresponding byte array.