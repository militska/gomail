package cmd

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"militska/gomail/internal"
	"net/http"
)

func SendMailHandler(ch chan internal.Email) {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Warningln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			_, err = w.Write([]byte("Content-Type must be application/json"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data internal.Email
		if err = json.Unmarshal(body, &data); err != nil {
			log.Warningln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = data.Check(); err != nil {
			log.Warningln(err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		ch <- data

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}
