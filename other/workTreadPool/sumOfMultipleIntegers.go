package main

import "fmt"

// NUMBER 工作池的goroutine数目
const (
	NUMBER = 10
)

// 工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

// 任务处理：计算begin到end的和 , 执行结果写入结果chan result
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

// InitTask 初始化待处理task chan
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		k := 10*j + 1
		m := 10 * (j + 1)
		tas := task{
			begin:  k,
			end:    m,
			result: r,
		}
		taskChan <- tas
	}

	if mod != 0 {
		tas := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tas
	}
	close(taskChan)
}

// DistributeTask 读取taskChan并且分发到worker goroutine处理， 总的数量是workers
func DistributeTask(taskChan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskChan, done)
	}
}

// ProcessTask 工作goroutine处理具体工作,并将处理结果发送到结果chan
func ProcessTask(taskChan <-chan task, done chan struct{}) {
	for t := range taskChan {
		t.do()
	}
	done <- struct{}{}
}

// CloseResult 通过done channel同步等待所有goroutine的结束，然后关闭结果chan
func CloseResult(done chan struct{}, resultChan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultChan)
}

// ProcessResult 读取结果通道，汇总结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}

func main() {
	workers := NUMBER

	//工作通道
	taskChan := make(chan task, 10)

	//结果通道
	resultChan := make(chan int, 10)

	//worker信号通道
	done := make(chan struct{}, 10)

	//初始化task的goroutine,计算100个自然数的和
	go InitTask(taskChan, resultChan, 100)

	//分发任务到NUMBER个goroutine池
	DistributeTask(taskChan, workers, done)

	//获取各个goroutine处理完任务的通知，并关闭通道
	go CloseResult(done, resultChan, workers)

	//通过结果通道获取结果并且汇总
	sum := ProcessResult(resultChan)

	fmt.Println(sum)
}
