package PriorityQueue

import "testing"

func TestPriorityQueueSimple(t *testing.T) {
	pq := make(PriorityQueue, 0)

	pq.Enqueue("D", 4)
	pq.Enqueue("A", 1)
	pq.Enqueue("C", 3)
	pq.Enqueue("B", 2)

	correct := []string{"A", "B", "C", "D"}

	size := len(pq)
	for i := 0; i < size; i++ {
		v := pq.Dequeue()
		if correct[i] != v {
			t.Fatalf("Got %s expected %s", v, correct[i])
		}
	}
}

func TestPriorityQueueComplex(t *testing.T) {
	pq := make(PriorityQueue, 0)

	pq.Enqueue("D", 4)
	pq.Enqueue("A", 1)
	pq.Enqueue("E", 5)
	pq.Dequeue()
	pq.Enqueue("B", 2)
	pq.Dequeue()
	pq.Enqueue("C", 3)

	correct := []string{"C", "D", "E"}

	size := len(pq)
	for i := 0; i < size; i++ {
		v := pq.Dequeue()
		if correct[i] != v {
			t.Fatalf("Got %s expected %s", v, correct[i])
		}
	}
}
