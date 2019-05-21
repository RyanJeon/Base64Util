package Base64Enc

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

func GetFileType(ByteArray []byte) types.Type {
	kind, _ := filetype.Match(ByteArray)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
	}
	return kind
}

func FileToBase64(file *os.File) string {
	ByteArray := FileToBuffer(file)
	Base64 := base64.StdEncoding.EncodeToString(ByteArray)
	return Base64
}

func FileToBase64Url(file *os.File) string {
	ByteArray := FileToBuffer(file)
	Base64 := base64.StdEncoding.EncodeToString(ByteArray)
	FileType := GetFileType(ByteArray)
	Mime := FileType.MIME.Value
	Url := "data:" + Mime + ";base64," + Base64
	return Url
}

//Loading Byte Slices into a Buffer and Return an array of bytes.
func FileToBuffer(file *os.File) []byte {
	FileReader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(FileReader) //Array of Byte Slices
	return content
}
