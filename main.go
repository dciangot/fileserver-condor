package main

import (
    "net/http"
)

func main(){
    http.Handle("/", http.FileServer(http.Dir("/data")))
    http.ListenAndServe(":31231", nil)
}
