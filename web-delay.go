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
		pars := r.URL.Query()
		ms := pars.Get("ms")
		msi := 0
		if ms == "" {
			maxi := 1000
			max := pars.Get("max")
			if max != "" {
				tmp, err := strconv.Atoi(max)
				if err != nil {
					tmp = 1
				}
				maxi = tmp
			}
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
		w.Write([]byte(fmt.Sprintf("Hello Go, delayed by %d ms.\n", msi)))
	})

	http.ListenAndServe("0.0.0.0:80", nil)
}
