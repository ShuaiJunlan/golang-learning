package linear

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueueConcurrencyManualLock(t *testing.T) {
	megaQueue := NewQueue(false)
	var group sync.WaitGroup
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaQueue.Lock()
				ok := megaQueue.Enqueue(times)
				if ok == false {
					t.Error("insert failed " + string(times))
				}
				megaQueue.Unlock()
				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaQueue.Lock()
				if megaQueue.HasElement() {
					_, ok := megaQueue.Dequeue()
					if ok == false {
						t.Error("lock failed, hasElement() but Pop() failed")
					}
				}
				megaQueue.Unlock()
				time.Sleep(time.Millisecond * 8)
			}
			group.Done()
		}()
	}
	group.Wait()
}

func TestQueueConcurrencyAutoLock(t *testing.T) {
	megaQueue := NewQueue(true)
	var group sync.WaitGroup
	for i := 0; i <= 10; i++ {
		go func() {
			if megaQueue.HasElement() {
				_, ok := megaQueue.Peek()
				if !ok {

				}
			}
			if megaQueue.String() == "" {
				t.Error("String() failed")
			}

			if megaQueue.Len() < 0 {
				t.Error("Len() failed")
			}
			time.Sleep(time.Millisecond * 3)
		}()
	}
	//spam enqueue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				ok := megaQueue.Enqueue(times)
				if ok == false {
					t.Error("insert failed " + string(times))
				}
				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	//spam dequeue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				if megaQueue.HasElement() {
					_, ok := megaQueue.Dequeue()
					if !ok {

					}
				}
				time.Sleep(time.Millisecond * 8)
			}
			group.Done()
		}()
		//group.Done() // dead lock?
	}
	group.Wait()
}

func TestQueueBasicTypes(t *testing.T) {
	for _, r := range fakeTable {
		testQueueFunctionality(t, r)
	}

}
func testQueueFunctionality(t *testing.T, toPush []interface{}) {
	s := NewQueue(false)
	for i, v := range toPush {
		ok := s.Enqueue(v)
		if ok == false {
			t.Error("Enqueue failed")
		}
		if s.Len() != i+1 {
			t.Errorf("len failed, expected %v, got %v, for %v", i, s.Len(), toPush)
		}
		value, ok := s.Peek()
		if ok == false {
			t.Error("peek failed")
		}
		if value != toPush[0] {
			t.Errorf("peek failed, expected %v, got %v", v, value)
		}
	}
	for _, v := range toPush {
		el, ok := s.Dequeue()
		if ok == false {
			t.Error("dequeue failed")
		}
		if el != v {
			t.Errorf("dequeue failed, expected %v, got %v", v, el)
		}
	}
	if s.HasElement() {
		t.Errorf("queue is not empty after all Pop(), size=%v", s.Len())
	}
}
func TestQueueInitPeekIsNil(t *testing.T) {
	s := NewQueue(false)
	peek, ok := s.Peek()
	if ok {
		t.Error("peek should be false when on an empty queue")
	}
	if peek != nil {
		t.Errorf("expected nil, got %v", peek)
	}
}
func TestQueueInitPopIsNil(t *testing.T) {
	s := NewQueue(false)
	var pop, ok = s.Dequeue()
	if ok {
		t.Error("Pop should be false when uses on an empty queue")
	}
	if pop != nil {
		t.Errorf("expected nill, got %v", pop)
	}
}
func TestQueueIsEmpty(t *testing.T) {
	helperInitIsEmpty("queue", NewQueue(false), t)
}

func TestQueueSkippingNewShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("the code did not panic")
		}
	}()
	s := Queue{}
	c := s.Len()
	fmt.Print(c)
}
func TestQueueStringer(t *testing.T) {
	s := NewQueue(false)
	v := s.String()
	if v != "Queue [0]" {
		t.Error("stringer was incorrect for 0 length" + v)
	}
	s.Enqueue(1)
	s.Enqueue(1)
	v = s.String()
	if v != "Queue [2]" {
		t.Error("stringer was incorrect for 2 length" + v)
	}
}
func BenchmarkQueueSync1000(b *testing.B) {
	benchQueueSync(1000, b)
}
func BenchmarkQueueSync100000(b *testing.B) {
	benchQueueSync(100000, b)
}
func BenchmarkQueueSync1000000(b *testing.B) {
	benchQueueSync(10000000, b)
}
func benchQueueSync(count int, b *testing.B)  {
	for i := 0; i < count; i++ {
		q := NewQueue(false)
		for c := 0; c < count; c++ {
			if q.Enqueue("a") == false {
				b.Error("q enqueue failed")
			}
		}
		for c := 0; c < count; c++ {
			if val, ok := q.Dequeue(); ok == false || val != "a" {
				b.Error("q dequeue failed")
			}
		}
	}
}





