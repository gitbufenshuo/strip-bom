package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func valid(path string) bool {
	stat, err := os.Stat(path)
	if err == nil {
		return !stat.IsDir()
	}
	return false
}

func main() {
	path := os.Args[1]
	if valid(path) {
		f, _ := os.Open(path)
		defer f.Close()
		b, _ := ioutil.ReadAll(f)
		if bytes.Contains(b, []byte{0xEF, 0xBB, 0xBF}) {
			b = b[3:]
			f.Write(b)
			fmt.Println("Done")
		}
	}
}
