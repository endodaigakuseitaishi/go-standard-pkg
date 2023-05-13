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
	mutex = new(sync.Mutex)
	
	for i := 0; i < 10; i++ {
		go func() {
			for i:=0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	for {

	}
}
