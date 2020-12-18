package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		ms := query.Get("ms")
		msi := 0
		if ms == "" {
			maxi := 1000
			max := query.Get("max")
			if max != "" {
				tmp, err := strconv.Atoi(max)
				if err != nil {
					tmp = 1
				}
				maxi = tmp
			}
			rand.Seed(time.Now().UnixNano())
			msi = rand.Intn(maxi)
		} else {
			tmp, err := strconv.Atoi(ms)
			if err != nil {
				tmp = 0
			}
			msi = tmp
		}
		if msi > 0 {
			time.Sleep(time.Duration(msi) * time.Millisecond)
		}
		text := query.Get("text")
		if text == "" {
			text = fmt.Sprintf("Hello Go, delayed by %d ms.\n", msi)
		}
		w.Write([]byte(text))
	})

	http.ListenAndServe("0.0.0.0:80", nil)
}
