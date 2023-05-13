package main

import (
	"fmt"
	"regexp"
	// "strings"
	// "strconv"
	// "time"
	// "os"
	// "log"

)

func main() {
	match, _ := regexp.MatchString("A", "AVSSS")
	fmt.Println(match)

	re1, _ := regexp.Compile("A")
	match = re1.MatchString("AVSSS")
	fmt.Println(match)

	re2 := regexp.MustCompile("A")
	match = re2.MatchString("VSSS")
	fmt.Println(match)

	
}
