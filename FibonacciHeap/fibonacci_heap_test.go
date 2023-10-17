package FibonacciHeap

import "testing"

func TestFibonacciHeapSimple(t *testing.T) {
	fh := FibonacciHeap{}
	fh.Insert(1)
	fh.Insert(3)
	fh.Insert(5)
	fh.Insert(2)
	fh.Insert(4)

	if fh.GetMin() != 1 {
		t.Fatalf("GetMin() failed, expected 1 got %d", fh.GetMin())
	}
	if fh.min.GetDegrees() != 1 {
		t.Fatalf("GetDegrees() failed, expected 1 got %d", fh.min.GetDegrees())
	}

	fh.ExtractMin()
	fh.ExtractMin()

	if fh.GetMin() != 3 {
		t.Fatalf("GetMin() failed, expected 3 got %d", fh.GetMin())
	}
	if fh.min.GetDegrees() != 2 {
		t.Fatalf("GetDegrees() failed, expected 2 got %d", fh.min.GetDegrees())
	}
}
