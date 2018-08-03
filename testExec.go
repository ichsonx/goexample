package main
/*
 golang exec的基本使用，就是调用本地os的命令或者可执行文件的example
 存在“notice”
*/
import (
	"fmt"
	"log"
	"os/exec"
	)

func main() {
	notice()
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
		return
	}
	fmt.Printf("The Output is %s\n", out)
}

/*
 注意：对于以下使用nc（netcat）命令必须使用CombinedOutput方法，才能获取“执行成功返回的结果信息”。
使用Output是无法获取“执行成功返回的结果信息”的，但它们都可以获取失败的返回消息

同样测试过ls -la、docker、docker images 等命令，通过Output方法都可以获取“执行成功返回的结果信息”。
目前只发现nc这个命令需要CombinedOutput方法获取“执行成功返回的结果信息”。
*/
func notice() {
	out, err := exec.Command("nc", "-n", "-v", "-z", "172.17.50.10", "445").CombinedOutput()
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	fmt.Printf("The Output is %v\n", string(out))
}

