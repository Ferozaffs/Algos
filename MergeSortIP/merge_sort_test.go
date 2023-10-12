package MergeSortIP

import "testing"

func TestMergeSortIPSimple(t *testing.T) {

	list := []int{4, 2, 3, 1}
	Sort(&list)

	t.Logf("%v", list)

	size := len(list)
	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("List not sorted, %d is larger than %d", list[i], list[i+1])
		}
	}
}

func TestMergeSortIPComplex(t *testing.T) {

	list := []int{23423, -1231, 4567, -12331, 3211, 2, 0, 312, 52}
	Sort(&list)

	t.Logf("%v", list)

	size := len(list)
	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("List not sorted, %d is larger than %d", list[i], list[i+1])
		}
	}
}
