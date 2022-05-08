package channels

import "fmt"

func consumer(cname string, ch chan int) {
	for i := range ch {
		fmt.Println("consumer--", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {

		// fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}
