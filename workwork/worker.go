package workwork

import (
	"awesomeProject/historydata"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Working(Id int, Url string,  Interval int , hist historydata.Adder) {
	for true {
	time.Sleep(time.Duration(Interval) * time.Second)
	resp, err := http.Get("http://" + Url)
	time_start := time.Now()

	if err != nil {
		fmt.Printf("%s", err)
		hist.Add(historydata.History{
			Response:  "Null",
			Duration:  time.Since(time_start),
			CreatedAt: time.Now(),
			Id:        Id,
		})
	}else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		hist.Add(historydata.History{
			Response:  string(contents),
			Duration:  time.Since(time_start),
			CreatedAt: time.Now(),
			Id:        Id,
		})}
	}







}