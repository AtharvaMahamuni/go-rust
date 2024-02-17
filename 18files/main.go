package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Let's make some files!")
	content := "This content is written inside the file."
	filePath := "./myfile.txt"

	file, err := os.Create(filePath)
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile(filePath)
}

func readFile(filePath string) {
	dataByte, err := os.ReadFile(filePath)
	checkNilErr(err)

	data := string(dataByte)
	fmt.Println("File Data\n", strings.TrimSpace(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
