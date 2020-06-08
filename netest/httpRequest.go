package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	url := "http://localhost:8085/iov/api/runtime-data/vehicleHttpPack"
	fmt.Println(url)
	contentType := "application/json"
	for cycle := 1; cycle > 0 ; cycle --  {
		json := "{\"year\":\"2019\",\"month\":\"9\"," +
			"\"day\":\"10\",\"sign\":\"123\",\"vehicle\":{\"app\":\"10\",\"vid\":\"5\",\"vtype\":\"mazida\"," +
			"\"rtDataMap\":{\"REPLACE_TIME_1\":{\"speed\":\"REPLACE_SPEED_1\",\"ecuMaxTemp\":\"REPLACE_ECU_MAX_TEMP_1\"},\"REPLACE_TIME_2\":" +
			"{\"speed\":\"REPLACE_SPEED_2\",\"ecuMaxTemp\":\"REPLACE_ECU_MAX_TEMP_2\"},\"REPLACE_TIME_3\":{\"speed\":\"REPLACE_SPEED_3\",\"ecuMaxTemp\":\"REPLACE_ECU_MAX_TEMP_3\"}," +
			"\"REPLACE_TIME_4\":{\"speed\":\"REPLACE_SPEED_4\",\"ecuMaxTemp\":\"REPLACE_ECU_MAX_TEMP_4\"}}}}"
		for i := 1; i < 5 ; i++  {
			json = strings.Replace(json,"REPLACE_TIME_" + strconv.Itoa(i),strconv.FormatInt(time.Now().UnixNano() / 1e6,10)  , -1)
			json = strings.Replace(json,"REPLACE_SPEED_"+ strconv.Itoa(i), strconv.Itoa(rand.Intn(100)), -1)
			json = strings.Replace(json,"REPLACE_ECU_MAX_TEMP_" + strconv.Itoa(i), strconv.Itoa(rand.Intn(100)), -1)
		}
		fmt.Println(json)
		data := strings.NewReader(json)
		resp, err := http.Post(url, contentType, data)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		time.Sleep(time.Duration(10) * time.Millisecond)
	}

}
