package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())

	var buffer1 bytes.Buffer
	buffer1.WriteString("bytes.Buffer example: bytes.Buffer.WriteString\n")
	fmt.Println(buffer1.String())

	// var buffer2 bytes.Buffer
	// // cannot use buffer2 (type bytes.Buffer) as type io.Writer in argument to io.WriteString:
	// //	bytes.Buffer does not implement io.Writer (Write method has pointer receiver)
	// io.WriteString(buffer2, "bytes.Buffer example: io.WriteString\n")
	// fmt.Println(buffer2.String())
}
