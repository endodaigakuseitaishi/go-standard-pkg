package main

import (
	// "fmt"
	"os"
	// "log"
)

func main() {
	// os.Exit(1)
	// fmt.Println(os.Args)

	// _, err := os.Open("test.txt")
	// if err!= nil {
  //   fmt.Println(err)
	// 	log.Fatal(err)
  // }

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Println("length=%d\n", len(os.Args))

	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// f, err := os.Open("test.txt")
	// if err!= nil {
  //   log.Fatal(err)
  // }
	// defer f.Close()

	f, _ := os.Create("test.txt")
	f.Write([]byte("hello world"))
	f.WriteAt([]byte("Go"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("foo")
}
