package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Parser(w http.ResponseWriter, req *http.Request, input interface{}) {
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(jsonDataFromHttp))
	err = json.Unmarshal(jsonDataFromHttp, input)
	defer req.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
}
