package main

import (
	"fmt"
	"sync"
	"runtime"
)

func worker(id int, result chan<- int, jobs <-chan int, group *sync.WaitGroup) {
	defer group.Done()
	for j := range jobs {
		fmt.Printf(" #%d processing job %d \n", id, j)
		result <- j * 2
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	group := &sync.WaitGroup{}
	resgroup := sync.WaitGroup{}
	jobs := make(chan int, 30)
	result := make(chan int, 30)
	worknum := 30
	group.Add(worknum)
	resgroup.Add(1)

	//建立线程池
	for i := 0; i < worknum; i++ {
		go worker(i, result, jobs, group)
	}

	//接收并打印结果。这里的顺序很重，不然可能会死锁。
	go func() {
		defer resgroup.Done()
		for res := range result {
			fmt.Printf("finish job %d \n", res)
		}
	}()

	//发送job
	for i := 0; i < 10000; i++ {
		jobs <- i
	}

	//推荐的权威实践是，使用sync.waitgroup来结束多线程
	//下面关闭jobs通道后，worker的group才好判断是否再有数据，从而group.done()
	//同理，worker全部处理完后，就关闭result的通道，result才好知道是否处理完，从而resgroup.down()
	//因为最后以result结束完才算完全结束完，所以最后的wait必须是result的resgroup。否则会出现result未处理完，程序就结束了。
	close(jobs)
	group.Wait()
	close(result)
	resgroup.Wait()

	fmt.Printf(" all finish !!!")
}
