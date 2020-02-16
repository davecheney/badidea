package badidea_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/davecheney/badidea"
)

func TestGoroutineID(t *testing.T) {
	got := badidea.GoroutineID()
	if got < 1 {
		t.Fatalf("expected goroutine id > 0, got: %v", got)
	}
	want, got := got, badidea.GoroutineID()
	if got != want {
		t.Fatalf("expected same goroutine id: %v, got: %v", want, got)
	}

	c := make(chan int64)
	go func() {
		c <- badidea.GoroutineID()
	}()
	got = <-c
	if got < want {
		t.Fatalf("expected new goroutine to have higher id than: %v, got %v", want, got)
	}
}

func ExampleGoroutineID() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("goroutine id:", badidea.GoroutineID())
		}()
	}
	wg.Wait()
}
