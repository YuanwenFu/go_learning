package select_mine 

import (
	"testing"
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result.")
		retCh <- ret 
		fmt.Println("service exited.")
	}()
	return retCh 
}

func TestSelect(t *testing.T) {
	select {
	case ret :=  <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out") 
	}
}
