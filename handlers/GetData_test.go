package handlers

import (
	"awesomeProject/mocked_http"
	"awesomeProject/urldata"
	"fmt"
	"net/http"
	"testing"
)

func TestGet(t *testing.T){
	data := urldata.New()
	data.Add(urldata.Item{Id:1,Url:"test" ,Interval:2})

	handler := GetData(data)

	w := &mocked_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()
	fmt.Println(result[0]["ID"])
	if result[0]["ID"] != 1{
		fmt.Println("???")
	}
	if len(result) != 1 {
		t.Errorf("Item was not added")
	}

	if result[0]["url"] != "test" {
		t.Errorf("Not propperlly set")
	}

}