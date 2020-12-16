package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pars := r.URL.Query()
		ms := pars.Get("ms")
		if ms == "" {
			ms = "0"
		}
		msi, err := strconv.Atoi(ms)
		if err != nil {
			msi = 0
		}
		time.Sleep(time.Duration(msi) * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Hello Go, visit %s/?ms=%s for delay %s ms.", r.Host, ms, ms)))
	})

	http.ListenAndServe("0.0.0.0:80", nil)
}
