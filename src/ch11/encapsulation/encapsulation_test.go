package encapsulation_test 

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string 
	Name string 
	Age int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0","Bob",20}
	e1 := Employee{Name:"Mike",Age:30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22 
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2) 
	fmt.Println("end!")
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }


func TestStructOperations(t *testing.T) {
	e := Employee{"0","Bob",20}
	//e := &Employee{"0","Bob",20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}