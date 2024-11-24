package func_test 

import "testing"
import "math/rand"
import "time"
import "fmt"

func returnMultiValues(a int) (int,int) {
	return a * a, a * a * a
}

func TestRandn(t *testing.T) {
	a := rand.Intn(10)
	t.Logf("a = %d", a)
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues(2)
	t.Log(a,b)
}

func generate_a_number(n int) int {
	res := rand.Intn(n)
	return res 
}

func timeSpent(inner func(op int)int) func(op int)int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret 
	}
}

func SlowFunc(op int) int {
	time.Sleep(time.Second*1)
	return op 
}

func TestFnTime(t *testing.T) {
	//timeSpent(rand.Intn)(100)
	//timeSpent(generate_a_number)(100)
	timeSpent(SlowFunc)(10)
}

func Sum(ops ...int) int {
	res := 0
	for _,x := range ops {
		res += x
	}
	return res 
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1))
	t.Log(Sum(1,2))
	t.Log(Sum(1,2,3))
	t.Log(Sum(1,2,3,4))
}

func Clear() {
	fmt.Println("Clear resources!")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start.")
	//panic("error")
	//fmt.Println("end.")
}