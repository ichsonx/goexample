package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readByBufio()
	readByBufio_buf()
	//readwriteByIoutil()
}

/*
	首先获取 *os.File 的这样的文件句柄（只读模式打开），然后通过bufio.NewReader来获取一个读取器变量。
	使用ReadString('\n')来逐行读取字符串，其中的 '\n' 不管是在windows还是linux，都可以这样用，无论中文、英文golang会自动处理好。
	当 bufio 的读取器读到最后没有的时候，会返回一个error，等于 io.EOF 的错误，提醒已经读完了。
*/
func readByBufio() {
	inputFile, inputError := os.Open("./examplefiles/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

/*
	这是上一个方法，但带缓冲 buf []byte 的做法，不会一次性读取整个文件。
	但有个问题就是，如果遇到中文、英文混合，那就容易出现乱码的情况，需要自己解决。
	例子中就是出现乱码的反例测试。
*/
func readByBufio_buf() {
	inputFile, inputError := os.Open("./examplefiles/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 16)

	for {
		n, _ := inputReader.Read(buf)
		if n == 0 {
			break
		}
		fmt.Printf("The input was: %s\n", string(buf))
	}
}

/*
	这里是使用 bufio 的 writer.WriteString 来写数据到文件。读文件的时候可以忽略权限，但在写文件的时候无论是windows还是linux的写权限都需要是0666.
	并且打开输出的文件时候需要指定可选操作例如 os.O_WRONLY(只写)|os.O_CREATE（如果文件不存在就创建），还有其他等等。
*/
func writeByBufio() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

/*
	整个文件读取可以使用 outil.ReadFile ，参数为文件路径。读取后返回的是 []byte 类型。
	方法 ioutil.WriteFile 写入文件，是对整个文件内容的覆盖，而非追加。
	ioutil 写文件时如果文件不存在则自动创建。
	网上资料显示（非正式），使用 ioutil 的读写速度都是较快的。
	明显，使用 ioutil.Write 的方式来写文件，只能全文覆盖，并没有追加的属性。
*/
func readwriteByIoutil() {
	inputFile := "./examplefiles/input.dat"
	outputFile := "./examplefiles/input_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

/*
	直接使用文件句柄来写入文件。
*/
func writeByFile() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}

/*
	只需要给出文件路径，使用 os 包提供的 Remove 方法即可完成删除。
*/
func deleteFile() {
	outputFile := "./examplefiles/input_copy.txt"
	os.Remove(outputFile)
}

/*
	只需要给出目录路径，使用 os 包提供的 Remove 方法即可完成删除。
*/
func deleteDir() {
	outputFile := "./examplefiles/input_copy.txt"
	os.RemoveAll(outputFile)
}

/*
	文件拷贝直接使用 io 包的 Copy 方法即可。
*/
func copyFile() (written int64, err error) {
	srcName := "./examplefiles/input.dat"
	dstName := "./examplefiles/input_copy.txt"
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
