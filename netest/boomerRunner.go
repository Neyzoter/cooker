package main

import (
	"encoding/csv"
	"github.com/myzhan/boomer"
	"log"
	"os"
	"strings"
	"time"
)

func sendBoom () {
	start := time.Now()
	elapsed := time.Since(start)
	body := getJson(row, timeMs)
	timeMs ++
	// http request
	// return OK : success
	resp := httpRequest(body)
	if strings.Compare(resp, "OK") == 0 {
		boomer.RecordSuccess("http", "send", elapsed.Nanoseconds() / int64(time.Millisecond), int64(10))
	} else {
		boomer.RecordFailure("http", "send", elapsed.Nanoseconds() / int64(time.Millisecond), "http error")
	}
}

func boomerRunner() {

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

	boomer.Run(task)

}
