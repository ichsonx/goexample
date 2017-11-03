package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("m.go")
	checck(err)
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	file2, err := os.Open("Gopkg.toml")
	checck(err)
	readline2(file2)

}

func readline2(file *os.File)  {
	scaner := bufio.NewScanner(file)
	for scaner.Scan(){
		fmt.Println(scaner.Text())
	}
}

func checck(e error)  {
	if e != nil {
		panic(e)
	}
}
