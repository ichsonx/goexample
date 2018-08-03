package main
/*
 golang exec的基本使用，就是调用本地os的命令或者可执行文件的example
*/
import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	simpleOutput()
}

/*
 command CombinedOutput 方法会运行其中的命令，并且获取命令执行结果的返回，这个返回是将“执行成功的结果”和“err”结合在一起一并返回。
如果执行成功，则返回的是“执行成功的结果”，如果失败，则返回的是“err”
*/
func simpleCombinedOutput() {
	cmd := exec.Command("ls", "-la")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

/*
 command Output 方法会运行其中的命令，并且获取命令执行结果的返回，包括err
*/
func simpleOutput() {
	out, err := exec.Command("ls", "-la").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Output is %s\n", out)
}
