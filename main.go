package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Modify(str string) string {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
checkError(err)

defer stdout.Close()
writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
str := readLine(reader)
result := ModifyString(str)

fmt.Fprintf(writer, "%s\n", result)
writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
	return ""
	}
	return strings.TrimRight(string(str), "\r\n")
	}
	func checkError(err error) {
	if err != nil {
	panic(err)
	}
	}
