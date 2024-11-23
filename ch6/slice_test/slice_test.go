package slice_test 

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int 
	t.Log(len(s0),cap(s0))
	s0 = append(s0, 11)
	t.Log(len(s0),cap(s0))

	s1 := []int {11,12,13,14}
	t.Log(len(s1),cap(s1))

	s2 := make([]int, 3, 5) //type,length,capacity
	t.Log(len(s2),cap(s2))
	t.Log(s2[0],s2[1],s2[2])
	s2 = append(s2,3)
	t.Log(s2[0],s2[1],s2[2],s2[3])
	t.Log(len(s2),cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s1 := []int{}
	for i:=0;i<10;i++ {
		s1 = append(s1,i)
		t.Log(len(s1),cap(s1))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string {"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log("year", year, "Q2", Q2)
}

func TestSliceCompare(t *testing.T) {
	s1 := []int {1,2,3,4}
	s2 := []int {1,2,3,4}
	// if s1 == s2 { //invalid operation: s1 == s2 (slice can only be compared to nil)
	// 	t.Log("equal")
	// }
}