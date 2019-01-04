package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

func main() {
	count, err := sh.Command("curl", "http://t66y.com/thread0806.php?fid=15&search=&page=2").Output()
	if err != nil{
		fmt.Print(err)
	}

	fmt.Println(string(count))
}
