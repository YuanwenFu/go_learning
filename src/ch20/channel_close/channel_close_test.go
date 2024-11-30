package channel_close 

import (
	"testing"
	"fmt"
	//"time"
	"sync"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
		//ch <- 11 //panic: send on closed channel
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		data := <- ch 
	// 		fmt.Println(data) 
	// 	}
	// 	wg.Done()
	// }() 

	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break 
			}
		}
		wg.Done()
	}() 
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup  
	ch := make(chan int) 
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}