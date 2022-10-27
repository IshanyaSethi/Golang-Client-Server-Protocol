package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// checks for error
func ErrorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	file_name := "land.png"
	file, er := os.Open("/home/asterix/spatial.guide/Golang_begin/" + file_name)

	//here we check if there was an error
	ErrorChecker(er)

	read := bufio.NewReader(file)

	data, err := ioutil.ReadAll(read)
	ErrorChecker(err)

	file_name = strings.Replace(file_name, ".txt", ".gz", -1)

	file, err = os.Create("/home/asterix/spatial.guide/Golang_begin/" + file_name)
	ErrorChecker(err)

	w := gzip.NewWriter(file)
	w.Write(data)
	// gives a notification when file compression is done
	fmt.Println("File compressed successfully")

	w.Close()
}
