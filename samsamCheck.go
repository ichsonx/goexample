package main

import (
	"os/exec"
	"fmt"
	"sync"
	"strconv"
	"strings"
	)

var (
	ips = make([][]string, 0)
	
)

func init() {
	for i := 1; i < 255; i++{
		ips = append(ips, []string{"-n", "-v", "-z", fmt.Sprintf("172.17.50.%s", strconv.Itoa(i)), "445"})
		ips = append(ips, []string{"-n", "-v", "-z", fmt.Sprintf("172.17.50.%s", strconv.Itoa(i)), "3389"})
	}
}

func main() {
	wg := sync.WaitGroup{}
	for _, ip := range ips {
		wg.Add(1)
		go execCmd(&wg, "nc", ip...)
	}
	wg.Wait()
}

func execCmd(wg *sync.WaitGroup, name string, arg ...string) {
	defer wg.Done()
	out, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		fmt.Printf("err exec command : %s %s\n", name, arg)
	}else if strings.Contains(string(out), "succee"){

	}
}
