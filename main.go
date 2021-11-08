package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	//use WriteString
	fmt.Println("hello")
	f0, err := os.OpenFile("f0.txt", os.O_RDWR, 0644)
	defer f0.Close()
	stringArray2 := []string{"Hello", "World"}
	justString := strings.Join(stringArray2, "")
	_, err = f0.WriteString(justString)
	err = f0.Sync()
	fmt.Println(err)

	//use fmt.Fprintf
	f1, err := os.Create("f1.txt")
	defer f1.Close()
	fmt.Fprintf(f1, justString)

	//use ioutil.WriteFile
	f2 := "f2.txt"
	byteString := []byte("Hello")
	ioutil.WriteFile(f2, byteString, 0644)

	//use bufio.NewWriter()
	f3, err := os.Create("f3.txt")
	w := bufio.NewWriter(f3)
	_, err = w.WriteString(justString)
	w.Flush()

	//use io.WriteString
	f4, err := os.Create("f4.txt")
	n, err := io.WriteString(f4, justString)
	fmt.Printf("wrote %d bytes\n", n)

}
