package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1:11,2:22,3:33}
	t.Log("m1",m1)
	t.Logf("len(m1) = %d", len(m1))
	m2 := map[int]int{} 
	m2[4]=16
	t.Logf("len(m2) = %d", len(m2))
	m3 := make(map[int]int, 10) //capacity
	t.Logf("len(m3) = %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v,flag := m1[1]; flag {
		t.Logf("key 1 is existing, its value is %d", v)
	} else {
		t.Log("key 1 is not existing!")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:11,2:22,3:33,4:44}
	for key,value := range m1 {
		t.Logf("key = %d, value = %d!", key, value)
	}
}