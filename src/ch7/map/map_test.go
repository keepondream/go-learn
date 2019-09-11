package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 4, 65: 99}
	t.Log(m1[1], m1[65])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 11
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[2]; ok {
		t.Log(ok)
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log(ok)
		t.Log("key 3 is not existing.")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 3: 6, 5: 10}
	for i, v := range m1 {
		t.Logf("key => %d,value => %d", i, v)
	}
}
