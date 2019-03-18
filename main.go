package main

import (
	"fmt"
	babble "github.com/tdcox/test110/babble"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	babbler := NewBabbler()
	babbler.Count = 3
	babbler.Separator = " "

	fmt.Fprintf(w, babbler.Babble()+"\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
