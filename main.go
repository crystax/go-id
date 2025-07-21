package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id1 := goid1()
			id2 := goid2()
			if id1 != id2 {
				panic("goid mismatch")
			}
		}()
	}

	wg.Wait()

	n := 100000
	d1 := measure(func() {
		for range n {
			goid1()
		}
	})

	d2 := measure(func() {
		for range n {
			goid2()
		}
	})

	fmt.Printf("d1: %.09fs\n", d1/1000000000)
	fmt.Printf("d2: %.09fs\n", d2/1000000000)
	fmt.Printf("diff: %.02f%%\n", 100*d2/d1)
}

func measure(f func()) float64 {
	start := time.Now()
	f()
	return float64(time.Since(start))
}

func goid1() uint64 {
	return getg().goid()
}

func goid2() uint64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	if n <= 0 {
		return 0
	}

	stack := buf[:n]

	m := regexp.MustCompile(`^goroutine\s+(\d+)`).FindSubmatch(stack)
	if len(m) != 2 {
		fmt.Fprintf(os.Stderr, "bad call stack: %s\n", string(stack))
		return 0
	}

	sid := string(m[1])

	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad id in stack %q: %q\n", string(stack), sid)
		return 0
	}

	return id
}

