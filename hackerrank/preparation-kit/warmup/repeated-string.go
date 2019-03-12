package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	chars := make(map[uint8]int64)
	fracChars := make(map[uint8]int64)
	slen := int64(len(s))
	times := n / slen
	leftN := n % slen
	for i := int64(0); i < slen; i++ {
		chars[s[i]]++
		if i < leftN {
			fracChars[s[i]]++
		}
	}
	return (chars[s[0]] * times) + fracChars[s[0]]
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout := os.Stdout
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	//s := readLine(reader)

	//n, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)

	result := repeatedString("aba", 100)

	fmt.Fprintf(writer, "%d\n", result)

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
