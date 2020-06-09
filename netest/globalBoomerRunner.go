package main

import (
	"encoding/csv"
	"github.com/myzhan/boomer"
	"log"
	"os"
	"strings"
	"time"
)

var timeMs = time.Now().UnixNano() / 1e6
var row []string
func send () {
	start := time.Now()
	elapsed := time.Since(start)
	body := getJson(row, timeMs)
	timeMs ++
	// http request
	resp := httpRequest(body)
	if strings.Compare(resp, "OK") == 0 {
		globalBoomer.RecordSuccess("http", "send", elapsed.Nanoseconds() / int64(time.Millisecond), int64(10))
	} else {
		globalBoomer.RecordFailure("http", "send", elapsed.Nanoseconds() / int64(time.Millisecond), "http error")
	}

}

var globalBoomer *boomer.Boomer

func globalBoomerRunner() {

	fileName := "/home/scc/code/go/cooker/netest/synthetic_data_with_anomaly-s-1-Transpose.csv"

	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)

	row, err = r.Read()
	task := &boomer.Task{
		Name: "send",
		// The weight is used to distribute goroutines over multiple tasks.
		Weight: 10,
		Fn: send,
	}

	numClients := 8
	// 无限循环运行之前，计算了停止的时间（sleep_time = 1.0 / self.hatch_rate），也就是说利用了sleep来达到每秒运行多少用户的效果。
	hatchRate := 10.0
	globalBoomer := boomer.NewStandaloneBoomer(numClients, hatchRate)
	globalBoomer.Run(task)

}
