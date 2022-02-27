package main

import(
	"fmt"
)

func main(){
	ch1 := make(chan int, 10)
	ch2 := make(chan string, 10)

	go func(){
		for i:=0;i<10;i++ {
			fmt.Println("write to channel:",string(i))
			ch2 <- string(i)
		}
		close(ch2)
	}()

	go func(){
		for i:=0;i<10;i++ {
			fmt.Println("write to channel:",i)
			ch1 <- i
		}
		close(ch1)
	}()

	// for v:= range ch{
	// 	fmt.Println("read from channel:",v)
	// }
	
	select{
	case v:= <- ch1:
		fmt.Println("read from channel1:",v)
	case v:= <- ch2:
		fmt.Println("read from channel2:",v)
	default:
		fmt.Println("channel no data")
	}


	// fmt.Println("hello world")
}