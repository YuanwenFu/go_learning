package cannel_by_context 

import (
	"fmt"
	"testing"
	"time"
	"context"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <- ctx.Done():
		return true 
	default:
		return false 
	}
}

func TestCannel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break 
				}
				time.Sleep(time.Millisecond * 5) 
			}
			fmt.Println(i, "Cancelled") 
		}(i, ctx) 
	}
	cancel()
	time.Sleep(time.Second * 1)
}