package array_test 

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[1], arr[2]) 

	arr1 := [4]int {1,2,3,4}
	arr2 := [...]int {1,2,3,4,5}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr1 := [4]int {11,12,13,14}
	for i:=0;i<len(arr1);i++ {
		t.Log(arr1[i])
	}
	for i,x := range arr1 {
		t.Log(i,x)
	}
	for _,x := range arr1 {
		t.Log(x)
	}
}

func TestArraySection(t *testing.T) {
	arr1 := [4]int {11,12,13,14}
	sec1 := arr1[:3]
	t.Log(sec1)
	sec2 := arr1[3:]
	t.Log(sec2)
}