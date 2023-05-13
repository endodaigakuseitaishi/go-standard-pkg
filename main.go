package main

import (
	"fmt"
	"time"
	// "os"
	// "log"
)

func main() {
	t2 := time.Date(2020, 5, 10, 12, 0, 0, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(time.Hour)

	t3 := time.Now()
	fmt.Println(t3)
	d2 := t3.Sub(t2)
	fmt.Println(d2)

	fmt.Println(t2.After(t3))
	fmt.Println(t2.Before(t3))
	fmt.Println(t3.Before(t2))
	fmt.Println(t3.After(t2))

	time.Sleep(5*time.Second)
	fmt.Println("aaaaaaa")
}
