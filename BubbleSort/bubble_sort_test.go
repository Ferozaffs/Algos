package BubbleSort

import "testing"

func TestBubbleSortSimple(t *testing.T) {

	list := []int{4, 2, 3, 1}
	cycles := Sort(&list)

	t.Logf("%v", list)

	if cycles != 3 {
		t.Fatalf("Cycles doesn't match, expected 3 got %d", cycles)
	}

	size := len(list)
	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("List not sorted, %d is larger than %d", list[i], list[i+1])
		}
	}
}

func TestBubbleSortEarlyReturn(t *testing.T) {

	list := []int{1, 2, 3, 4}
	cycles := Sort(&list)

	t.Logf("%v", list)

	if cycles != 1 {
		t.Fatalf("Cycles doesn't match, expected 1 got %d", cycles)
	}

	size := len(list)
	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("List not sorted, %d is larger than %d", list[i], list[i+1])
		}
	}
}

func TestBubbleSortComplex(t *testing.T) {

	list := []int{23423, -1231, 4567, -12331, 3211, 2, 0, 312}
	cycles := Sort(&list)

	t.Logf("%v", list)

	if cycles != 5 {
		t.Fatalf("Cycles doesn't match, expected 5 got %d", cycles)
	}

	size := len(list)
	for i := 0; i < size-1; i++ {
		if list[i] > list[i+1] {
			t.Fatalf("List not sorted, %d is larger than %d", list[i], list[i+1])
		}
	}
}
