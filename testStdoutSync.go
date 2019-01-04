package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//for i := 0; i < 50; i++ {
	//	time.Sleep(500 * time.Millisecond)
	//	h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
	//	fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
	//	os.Stdout.Sync()
	//}
	//fmt.Println("\nAll System Go!")

	//fmt.Printf("%.4f \n", float32(34) / 450)
	for i := 34; i < 245; i++ {
		time.Sleep(300 * time.Millisecond)
		rate := float32(i) / float32(244)
		h := strings.Repeat("=", int(rate*100)) + strings.Repeat(" ", 100-int(rate*100))
		fmt.Printf("\r%.2f%%[%s]", rate*100, h)
		os.Stdout.Sync()
		if int(rate) == 100 {
			fmt.Println("finish")
		}
	}
}
