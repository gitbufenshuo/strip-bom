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
		fr, _ := os.Open(path)
		defer fr.Close()
		b, _ := ioutil.ReadAll(fr)
		if bytes.Contains(b, []byte{0xEF, 0xBB, 0xBF}) {
			b = b[3:]
			fw, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			defer fw.Close()
			n, err := fw.Write(b)
			if err != nil {
				fmt.Println(n, err.Error())
			}

		}
	}
	fmt.Println(path, "Done")
}
