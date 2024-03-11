package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type file struct {
	file_ptr *os.File
}

func open_file(path string) file {

	ptr, err := os.Open(path)

	fil := file{
		file_ptr: ptr,
	}
	if err != nil {

		log.Fatal(err)
		return file{}
	}
	fmt.Println(fil)

	return fil
}

func (f *file) parse_file() map[string]string {

	mp := make(map[string]string)

	sc := bufio.NewScanner(f.file_ptr)

	for sc.Scan() {
		txt := sc.Text()
		txt_slice := strings.Split(txt, ":") // processing the text
		fmt.Println(txt_slice)
		k := txt_slice[0]
		v := txt_slice[1]
		mp[k] = v
	}
	return mp

}

func main() {

	f1 := open_file("logger_file.txt") //  i use the logger file create bu logger go program
	mp := f1.parse_file()              // opening file to read
	fmt.Printf("%#v", mp)

}
