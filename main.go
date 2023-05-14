package main

import (
	// "fmt"
	// "io/ioutil"
	// "log"

	// "sort"
	// "sync"
	// "regexp"
	// "strings"
	// "strconv"
	// "time"
	// "os"
	// "log"
	// "encoding/json"
	// "context"
	// "net/url"
	// "fmt"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "net/url"
)

func main() {
	// URL解析
	// u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	// fmt.Println(u)
	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)
	// fmt.Println(u.RawQuery)
	// fmt.Println(u.Fragment)
	// fmt.Println(u.Query())

	// URL生成
	// url := &url.URL{}
	// url.Scheme = "http:"
	// url.Host = "google.com"
	// q := url.Query()
	// q.Set("q", "golang")
	
	// url.RawQuery = q.Encode()

	// fmt.Println(url)

	// Getメソッド
	// resは構造体で受け取る
	// res, _ := http.Get("http://example.com/")
	// fmt.Println(res.StatusCode)
	// fmt.Println(res.Proto)
	// fmt.Println(res.Header["Content-Type"])
	// fmt.Println(res.Header["Date"])

	// fmt.Println(res.Request.Method)
	// fmt.Println(res.Request.URL)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	// POSTメソッド
	// vs := url.Values{}
	// vs.Add("id", "1")
	// vs.Add("message", "hello world")
	// fmt.Println(vs.Encode())

	// res, err := http.PostForm("http://example.com/", vs)
	// if err!= nil {
  //   log.Fatal(err)
  // }
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	fmt.Println("Hello, World!")
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}
// type MyHandler struct {}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Hello, %q", r.URL.Path)
// }

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temple.html")
	if err!= nil {
		log.Println(err)
  }
	t.Execute(w, "aaaaaaaaaaaa")
}
