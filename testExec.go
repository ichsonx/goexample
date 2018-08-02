package main

import (
	"os/exec"
			"io/ioutil"
	"os"
	"syscall"
	"fmt"
)

func main() {
	_, out, err := simple()

	if err != nil {
		fmt.Printf("error \n")
	}

	fmt.Println(out)
}

//func simple()  {
//	//cmd := exec.Command("touch", "aaaaa")
//	cmd := exec.Command("ls", "./")
//	outpip, err := cmd.StdoutPipe()
//
//	err = cmd.Run()
//	if err != nil {
//		fmt.Printf("error: %v \n" , err)
//		return
//	}
//	out, err := ioutil.ReadAll(outpip)
//	fmt.Println(out)
//	//fmt.Println(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
//	fmt.Println("command finish...")
//}

func simple() (int, string, error) {
	cmd := exec.Command(os.Getenv("SHELL"), "ls /")

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	outpip, err := cmd.StdoutPipe()
	if err != nil {
		return 0, "", err
	}

	err = cmd.Start()
	if err != nil {
		return 0, "", err
	}

	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return 0, "", err
	}

	return cmd.Process.Pid, string(out), nil
}
