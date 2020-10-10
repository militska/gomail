package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func sendMailHandler() {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var data SendEmailDTO
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}

		sendMsg(data)
	})
}
