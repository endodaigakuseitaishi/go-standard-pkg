package main

import (
	"fmt"
	"strconv"
	// "time"
	// "os"
	// "log"
)

func main() {
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)
}
