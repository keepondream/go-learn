package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for index := 0; index < 10; index++ {
		s = append(s, 1)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]

	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(summer)

}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// 数组可以用 == 比较
	// slice 只能用 nil 比较
	// 数组声明 [...|n]int{}
	// slice 声明  []int{}  区别是不指定长度
	if a != nil {
		t.Log("equal")
	}
	t.Log(a, b)
}
