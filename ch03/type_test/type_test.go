package type_test 

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	//var a int = 1
	var a int32 = 1
	var b int64 
	//b = a //cannot use a (variable of type int) as int64 value in assignment
	b = int64(a)
	t.Log(a,b)

	var c MyInt 
	//c = b //cannot use b (variable of type int64) as MyInt value in assignment
	c = MyInt(b)
	t.Log(b,c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a 
	//aPtr += 1 //invalid operation: aPtr += 1 (mismatched types *int and untyped int)
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string 
	t.Log("*" + s + "*")
	t.Log(len(s))
	// if s == nil { //invalid operation: s == nil (mismatched types string and untyped nil)
	// 	t.Log("s is empty!")
	// }
	if s == "" {
		t.Log("s is empty 111")
	}
}