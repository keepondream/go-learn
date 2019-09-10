package condition

import "testing"

func TestSwitchMultiCase(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for index := 0; index < 5; index++ {
		switch {
		case index%2 == 0:
			t.Log("Even")
		case index%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
