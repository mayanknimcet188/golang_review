package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// bytes.Buffer type implements both the Reader and Writer interfaces
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf)
	var src_buf bytes.Buffer
	var dst_buf bytes.Buffer
	src_buf.Write([]byte("This is a source buffer"))
	// io package contains two important interfaces
	// Reader and the Writer and Read and Write methods are respectively
	// to be implemented by any type that implements these interfaces
	written, _ := io.Copy(&dst_buf, &src_buf)
	fmt.Println(written)
	fmt.Println(dst_buf.String())
	// Incase we only have to read from strings and not files strings.NewReader is more efficient than bytes.Buffer
	str := "Sample String"
	fmt.Println(strings.NewReader(str))
}
