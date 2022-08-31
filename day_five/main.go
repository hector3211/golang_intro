package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("file.txt")
	Writer := io.Writer(file)
	n, err := Writer.Write([]byte("hello world"))
	fmt.Println(n, err)
	fmt.Println("%s", n)
	file.Close()
}
