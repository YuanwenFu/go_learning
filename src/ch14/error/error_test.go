package err_test 

import "testing"
import "errors"
import "fmt"

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100!")

func GetFibonacci(n int) ([]int,error) {
	// if n < 2 || n > 100 {
	// 	return nil, errors.New("n should be in [2,100]")
	// }
	if n < 2{
		return nil, LessThanTwoError
	} else if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1,1}
	for i:=2;i<n;i++ {
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

func TestGetFibonacci(t *testing.T) {
	//t.Log(GetFibonacci(10))
	//t.Log(GetFibonacci(-10))
	if v,err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

