package main

import (
	"fmt"
	"sort"
	// "sync"
	// "regexp"
	// "strings"
	// "strconv"
	// "time"
	// "os"
	// "log"
	// "encoding/json"
)

type Entry struct {
	Name string
	Value int
}


func main() {
	i := []int{3, 6, 4, 2, 8, 5, 9, 1, 7}
	s := []string{"x", "b", "c", "d", "e", "d", "a"}

	fmt.Println(i, s)
	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)

	el := []Entry {
		{"w", 1023},
		{"b", 323},
    {"v", 3435},
    {"q", 5},
    {"u", 544},
    {"c", 686},
	}
	// fmt.Println(el)
	fmt.Printf("%T", el)

	// sort.Slice(el, func(i, j int) bool {return el[i].Value < el[j].Value})
	// fmt.Println("--------------------------------------------------------")
	// fmt.Println(el)
	// fmt.Println("--------------------------------------------------------")
}
