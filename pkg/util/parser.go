package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Parser(w http.ResponseWriter, req *http.Request, input interface{}) {
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(jsonDataFromHttp))
	json.Unmarshal(jsonDataFromHttp, input)
	defer req.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
}
