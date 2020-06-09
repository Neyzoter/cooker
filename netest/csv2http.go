package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type RuntimeData struct {
	Val1 string `json:"val1"`
	Val2 string `json:"val2"`
	Val3 string `json:"val3"`
	Val4 string `json:"val4"`
	Val5 string `json:"val5"`
	Val6 string `json:"val6"`
	Val7 string `json:"val7"`
	Val8 string `json:"val8"`
	Val9 string `json:"val9"`
	Val10 string `json:"val10"`
	Val11 string `json:"val11"`
	Val12 string `json:"val12"`
	Val13 string `json:"val13"`
	Val14 string `json:"val14"`
	Val15 string `json:"val15"`
	Val16 string `json:"val16"`
	Val17 string `json:"val17"`
	Val18 string `json:"val18"`
	Val19 string `json:"val19"`
	Val20 string `json:"val20"`
	Val21 string `json:"val21"`
	Val22 string `json:"val22"`
	Val23 string `json:"val23"`
	Val24 string `json:"val24"`
	Val25 string `json:"val25"`
	Val26 string `json:"val26"`
	Val27 string `json:"val27"`
	Val28 string `json:"val28"`
	Val29 string `json:"val29"`
	Val30 string `json:"val30"`
}

type Vehicle struct {
	App int64  `json:"app"`
	Vid string `json:"vid"`
	Vtype string `json:"vtype"`
	RtDataMap map[string]RuntimeData `json:"rtDataMap"`
}

type VehicleHttpPack struct {
	Year string `json:"year"`
	Month string `json:"month"`
	Day string `json:"day"`
	Sign string `json:"sign"`
	Vehicle Vehicle `json:"vehicle"`
}

// int transform to string with "
func int2str (v int) string {
	return "\"" + strconv.Itoa(v) + "\""
}

// string transform to string with "
func str2str (s string) string {
	return "\"" + s + "\""
}

// get json format from one-point data
func getJson (row []string, timeMs int64) string{
	// init runtime data by reflect
	//var runtimeDataIf interface{}
	runtimeData := new(RuntimeData)
	runtimeDataV := reflect.ValueOf(runtimeData)
	if runtimeDataV.Kind()==reflect.Ptr {
		elem := runtimeDataV.Elem()
		for i := 1; i <= 30;i++  {
			v := elem.FieldByName("Val" + strconv.Itoa(i))
			if v.Kind()==reflect.String{
				*(*string)(unsafe.Pointer(v.Addr().Pointer())) = row[i-1]
			}
		}

	}

	rtData := make(map[string]RuntimeData)
	rtData[strconv.FormatInt(timeMs, 10)] = *runtimeData

	vehicle := Vehicle{
		App:       1,
		Vid:       "4",
		Vtype:     "mazida",
		RtDataMap: rtData,
	}
	vehicleHttpPack := VehicleHttpPack{
		Year:    "2020",
		Month:   "3",
		Day:     "6",
		Sign:    "123",
		Vehicle: vehicle,
	}
	vehicleHttpPackJsonByte,err := json.Marshal(vehicleHttpPack)
	if err != nil{
		log.Fatalf("Can not trans to Json : %+v", err)
	}
	vehicleHttpPackJsonStr := string(vehicleHttpPackJsonByte)
	return vehicleHttpPackJsonStr
}
// 返回OK的个数
var respOKNum = 0;

// http request
func httpRequest (body string) (r string) {
	url := "http://localhost:8085/iov/api/runtime-data/vehicleHttpPack"
	contentType := "application/json"
	resp, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "ERR"
	}
	response := string(b[:])
	//if strings.Compare(response, "OK") == 0 {
	//	respOKNum ++;
	//}
	return response
	//fmt.Println(string(body))
}

// read csv all one time
func readCsvAll (filename string) [][]string{
	fs1, _ := os.Open(filename)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content
}

func sendTest () {
	//// use 8 kernel
	//runtime.GOMAXPROCS(8)
	fileName := "/home/scc/code/go/cooker/netest/synthetic_data_with_anomaly-s-1-Transpose.csv"
	//fileName := "/home/scc/code/go/cooker/netest/test.csv"

	timeMs := time.Now().UnixNano() / 1e6
	// start running time
	startime := time.Now().UnixNano() / 1e6
	// data points num
	lines := 0

	//// read all lines
	//content := readCsvAll(fileName)
	//for _, row := range content {
	//	body := getJson(row, timeMs)
	//	httpRequest(body)
	//	timeMs += 10
	//	lines ++
	//}
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	for maxCycle := 1000;maxCycle > 0; maxCycle --{
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		lines ++
		if err == io.EOF {
			break
		}
		//fmt.Println(vehicleHttpPackJsonStr)
		//fmt.Println()
		timeMs += 10
		body := getJson(row, timeMs)
		// http request
		httpRequest(body)
		// sleep
		//time.Sleep(time.Duration(10) * time.Millisecond)
	}

	endtime := time.Now().UnixNano() / 1e6

	fmt.Printf("Send %d points in %d ms\n Received %d OK (%.2f %%)",lines, endtime - startime, respOKNum, float32(respOKNum) / float32(lines) * 100.0)
}


