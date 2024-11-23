package constant_test 

import "testing"

// const (
// 	Monday = 1
// 	Tuesday = 2
// 	Wednesday = 3 
// 	Thursday = 4 
// 	Friday = 5
// 	Saturday = 6
// 	Sunday = 7
// )

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota  
	Writable 
	Executable
)

func TestConstantTry (t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func TestConstantTry1 (t *testing.T) {
	//a := 7 //0111
	a := 1 //0001
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Executable==Executable)
}