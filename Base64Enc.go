package Base64Enc

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

//Returns file type object given an array of bytes of a file.
func GetFileType(ByteArray []byte) types.Type {
	FileType, _ := filetype.Match(ByteArray)
	return FileType
}

//Loading Byte Slices into a Buffer and Return an array of bytes.
func FileToBuffer(file *os.File) []byte {
	FileReader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(FileReader) //Array of Byte Slices
	return content
}

//Encode File to normal Base64
func FileToBase64(file *os.File) string {
	ByteArray := FileToBuffer(file)
	Base64 := base64.StdEncoding.EncodeToString(ByteArray)
	return Base64
}

//Encode File to Base64url
func FileToBase64Url(file *os.File) string {
	ByteArray := FileToBuffer(file)
	Base64 := base64.StdEncoding.EncodeToString(ByteArray)
	FileType := GetFileType(ByteArray)
	Mime := FileType.MIME.Value
	Url := "data:" + Mime + ";base64," + Base64
	return Url
}

//Convert Base64/Base64Url Encoded string back to array of bytes.
func Base64ToByte(code string) []byte {
	SplitStringArray := strings.Split(code, ",")                                               //If Base64Url, [data, base64], if Base64. [base64]
	ByteArray, _ := base64.StdEncoding.DecodeString(SplitStringArray[len(SplitStringArray)-1]) //Get the last element of the array.
	return ByteArray
}
