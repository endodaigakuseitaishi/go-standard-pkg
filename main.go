package main

import (
	"fmt"
	"sync"
	// "regexp"
	// "strings"
	// "strconv"
	"time"
	// "os"
	// "log"
)

var st struct{A, B, C int}

var mutex *sync.Mutex

func UpdateAndPrint(n int) {

	mutex.Lock()

	st.A = n
	time.Sleep(time.Millisecond)
	st.B = n
	time.Sleep(time.Millisecond)
	st.C = n
	time.Sleep(time.Millisecond)
	fmt.Println(st)

	mutex.Unlock()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
      fmt.Println(1)
    }
		wg.Done()
	}()
	go func() {
    for i := 0; i < 100; i++ {
      fmt.Println(2)
    }
		wg.Done()
  }()
	go func() {
    for i := 0; i < 100; i++ {
      fmt.Println(3)
    }
		wg.Done()
  }()

	wg.Wait()
}
