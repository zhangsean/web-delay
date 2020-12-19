package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

// Request 处理中的请求
type Request struct {
	// 请求ID
	id int
	// 休眠多少毫秒
	delay int
	// 开始时间
	begin time.Time
	// 结束时间
	end time.Time
}

const cacheKey string = "reqList"

func main() {
	c := cache.New(10*time.Minute, 10*time.Minute)

	r := mux.NewRouter()
	r.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		var requests []Request
		if obj, found := c.Get(cacheKey); found {
			requests = obj.([]Request)
		}
		html := "<table>"
		html += "<tr><th>ID</th><th>Delay(ms)</th><th>Begin time</th><th>Done time</th></tr>"
		done := r.URL.Query().Get("done")
		var aCount, dCount, pCount int
		for _, t := range requests {
			aCount++
			if t.end.IsZero() {
				pCount++
			} else {
				dCount++
			}
			if done == "1" && t.end.IsZero() {
				continue
			} else if done == "0" && !t.end.IsZero() {
				continue
			}
			html += "<tr>"
			html += fmt.Sprintf("<td>%v</td>", t.id)
			html += fmt.Sprintf("<td>%v</td>", t.delay)
			html += fmt.Sprintf("<td>%v</td>", t.begin.Format("2006-01-02 15:04:05.000000"))
			if !t.end.IsZero() {
				html += fmt.Sprintf("<td>%v</td>", t.end.Format("2006-01-02 15:04:05.000000"))
			}
			html += "</tr>"
		}
		html += "</table>"
		html = fmt.Sprintf("<a href='?done=0'>Processing requests (%d)</a>", pCount) + html
		html = fmt.Sprintf("<a href='?done=1'>Done requests (%d)</a>&nbsp;&nbsp;&nbsp;", dCount) + html
		html = fmt.Sprintf("<a href='?'>All requests (%d)</a>&nbsp;&nbsp;&nbsp;", aCount) + html
		w.Write([]byte(html))
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
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

		// 缓存记录处理中的请求列表
		var requests []Request
		if obj, found := c.Get(cacheKey); found {
			requests = obj.([]Request)
		}
		id := len(requests) + 1
		curTask := Request{
			id,
			msi,
			begin,
			time.Time{},
		}
		requests = append(requests, curTask)
		c.SetDefault(cacheKey, requests)
		log(id, "Received:", len(requests))
		var processing int
		for _, task := range requests {
			if task.end.IsZero() {
				processing++
			}
		}
		log(id, "Processing:", processing)

		if msi > 0 {
			// 休眠以模拟延迟
			log(id, "Sleep(ms):", msi)
			time.Sleep(time.Duration(msi) * time.Millisecond)
		}

		// 响应文本
		text := query.Get("text")
		if text == "" {
			text = fmt.Sprintf("Hello Go, delayed by %d ms.\n", msi)
		}
		w.Write([]byte(text))
		if obj, found := c.Get("processing"); found {
			processing = obj.(int)
		}

		// 标记任务完成
		if obj, found := c.Get(cacheKey); found {
			requests = obj.([]Request)
		}
		processing = 0
		for i, task := range requests {
			if task.id == id {
				requests[i].end = time.Now()
				curTask = requests[i]
			}
			if requests[i].end.IsZero() {
				processing++
			}
		}
		c.SetDefault(cacheKey, requests)
		log(id, "Remain:", processing)
		log(id, "NumGoroutine:", runtime.NumGoroutine())
		log(id, "Cost(ms):", time.Now().Sub(begin).Milliseconds())
		log(id, "Request info:", curTask.id, curTask.delay, curTask.begin.Format("2006-01-02 15:04:05.000000"), curTask.end.Format("2006-01-02 15:04:05.000000"))
	})

	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:808", nil)
}

func log(id int, msg ...interface{}) {
	fmt.Println(id, time.Now().Format("2006-01-02 15:04:05.000000"), msg)
}
