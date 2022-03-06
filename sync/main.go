package main

import(
	"fmt"
	"sync"
)

func waitByChannel(){
	ch := make(chan int, 10)
	for i:= 0;i < 10;i++ {
		go func(i int){
			fmt.Println("send to channel:", i)
			ch <- i
		}(i)
	}
	
	for i:= 0;i < 10;i++ {
		fmt.Println("read from channel:", <- ch)
	}
}

func waitByGroup(){
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:= 0;i < 10;i++ {
		go func(i int){
			fmt.Println("println index:", i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("restore operation")
}

func main(){
	waitByChannel()
	waitByGroup()
	// fmt.Println("hello world")
}