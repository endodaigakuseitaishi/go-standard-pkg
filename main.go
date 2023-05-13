package main

import (
	"fmt"
	"strings"
	// "strconv"
	// "time"
	// "os"
	// "log"
)

func main() {
	s1 := strings.Index("AXSWEDFR", "R")
	fmt.Println(s1)

	s2 := strings.LastIndex("AXSWEDFR", "R")
	fmt.Println(s2)

	s3 := strings.Contains("AQDCBXNE", "E")
	fmt.Println(s3)

	s4 := strings.Count("AQDCBXNEAA", "A")
	fmt.Println(s4)

	s5 := strings.Replace("AQDCBXNEAA", "A", "B", -1)
	s6 := strings.Replace("AQDCBXNEAA", "A", "B", 1)
  fmt.Println(s5, s6)
}
