package main

import (
	"fmt"
	// "sync"
	// "regexp"
	// "strings"
	// "strconv"
	"time"
	// "os"
	// "log"
	"encoding/json"
)

type A struct{}

type User struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created"`
	A *A `json:"A,omitempty"`
}

func (u User) MarshalJson() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr. " + u.Name,
	})
	return v, err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "hzdkv@example.com"
	u.CreatedAt = time.Now()

	bs, err := json.Marshal(u)
	if err!= nil {
    panic(err)
  }
	fmt.Println(string(bs))
}
