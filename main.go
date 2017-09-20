package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	
	LoadConfig()

	http.HandleFunc("/webhook",func(w http.ResponseWriter,r *http.Request){
		r.ParseForm()

		if(r.Method == "POST") {
			event := WebhookEvent{}
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&event)
			if err != nil {
				log.Fatal(err)
			}

			msg := event.Entry[0].Messaging[0]

			SendImage("nebula.jpg",msg.Sender)

		}

		w.Write([]byte(r.FormValue("hub.challenge")))
	})
	
	http.ListenAndServe(":5000",nil)
}