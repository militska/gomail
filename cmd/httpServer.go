package cmd

import (
	log "github.com/sirupsen/logrus"
	"militska/gomail/internal"
	"militska/gomail/tools"
	"net/http"
	"os"
	"runtime"
)

func InitHTTPServer() {
	if os.Getenv("ENABLED_API") == tools.ENABLED {
		s := &http.Server{
			Addr:           ":8070",
			ReadTimeout:    tools.DefaultTimeout,
			WriteTimeout:   tools.DefaultTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		ch := make(chan internal.Email, tools.EmailSizeBuffer)
		cnt := runtime.NumCPU()

		SendMailHandler(ch)

		for i := 1; i < cnt; i++ {
			go internalSend(ch)
		}
		log.Warning(s.ListenAndServe())
	}
}

func internalSend(ch chan internal.Email) {
	for {
		data := <-ch
		internal.SendMsg(&data)
	}
}
