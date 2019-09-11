package map_test

import "testing"

// 集合的初始化值为零值
// 利用if的多条件和 集合的多返回值 判断是否存在
// range遍历数组 index , value ; rang 遍历集合map key , value
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 3, 2: 6, 3: 9}
	t.Log(m1[1])
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
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 5, 5: 19}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
