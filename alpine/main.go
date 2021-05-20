package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
                fmt.Fprint(w, "Hello! Simple server for MIF. Your request was processed.\n")
        },
        )
        http.ListenAndServe(":8000", nil)
}

