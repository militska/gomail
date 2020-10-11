package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func sendMailHandler() {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if r.Header.Get("Content-Type") != "application/json" {
			_, _ = w.Write([]byte("Content-Type must be application/json"))
			w.WriteHeader(400)
			return
		}

		if err != nil {
			w.WriteHeader(400)
			return
		}

		var data SendEmailVo
		if err = json.Unmarshal(body, &data); err != nil {
			w.WriteHeader(400)
			return
		}

		if errDataCheck := data.check(); errDataCheck != nil {
			w.WriteHeader(400)
			_, _ = w.Write([]byte(errDataCheck.Error()))
			return
		}

		go sendMsg(data)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}
