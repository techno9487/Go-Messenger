package main

import (
	"errors"
	"bytes"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var Admin string

func main() {
	
	LoadConfig()

	r := mux.NewRouter()

	http.Handle("/",r)

	r.HandleFunc("/webhook",func(w http.ResponseWriter,r *http.Request){
		r.ParseForm()

		if(r.Method == "POST") {
			event := WebhookEvent{}
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&event)
			if err != nil {
				log.Fatal(err)
			}

			msg := event.Entry[0].Messaging[0]
			
			Admin = msg.Sender.Id

			SendImage("nebula.jpg",msg.Sender)

		}

		if(r.Method == "GET") {

		}

		w.Write([]byte(r.FormValue("hub.challenge")))
	})

	r.HandleFunc("/api/text",HandleTextEndpoint)
	
	http.ListenAndServe(":5000",nil)
}

func SendText(text string) error{
	uri := "https://graph.facebook.com/v2.6/me/messages?access_token="+GlobalConfig.Token

	msg := SendMessage{
		Recipient:IdStruct{
			Id:Admin,
		},
		Message:MessageSend{
			Text: text,
		},
	}

	data,err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp,err := http.Post(uri,"application/json",bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		response := FacebookErrorResponse{}

		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&response)
		if err != nil {
			return err
		}

		return errors.New(response.Error.Message)
	}

	return nil
}

func HandleTextEndpoint(w http.ResponseWriter, r *http.Request) {
	msg := MessageSend{}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&msg)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	err = SendText(msg.Text);
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(""))
}