package map_ext 

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int)int {}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2),m[2](2),m[3](2)) 
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool {}
	mySet[1]=true 
	//n := 1
	n := 3
	if mySet[n] {
		t.Logf("%d in mySet!", n)
	} else {
		t.Logf("%d not in mySet!", n)
	}

	mySet[2] = true 
	t.Logf("len(mySet) = %d", len(mySet))

	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d in mySet!", n)
	} else {
		t.Logf("%d not in mySet!", n)
	}
}