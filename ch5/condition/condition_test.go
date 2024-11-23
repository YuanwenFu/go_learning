package condition_test 

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i:=0;i<5;i++ {
		switch i {
		case 0,2:
			t.Log(i,"i is 0 or 2!")
		case 1,3:
			t.Log(i,"i is 1 or 3!")
		default:
			t.Log(i,"i not in (0,1,2,3)!")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i:=0;i<5;i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "is even!")
		case i%2 == 1:
			t.Log(i, "is odd!")
		default:
			t.Log("unkown!")
		}
	}
}