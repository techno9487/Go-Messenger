package main

import (
	"bytes"
	"log"
	"os"
	"mime/multipart"
	"io"
	"encoding/json"
	"net/http"
	"fmt"
	"net/textproto"
)

func SendImage(path string,recipient IdStruct) {

	var buf bytes.Buffer

	mp := multipart.NewWriter(&buf)

	f,err := os.Open(path)

	if err != nil {
		log.Println(err)
		return
	}

	defer f.Close()

	fw,err := CreateImageFile(mp,"nebula.jpg")

	if err != nil {
		log.Println(err)
		return
	}

	_,err = io.Copy(fw,f)
	if err != nil {
		log.Println(err)
		return
	}

	fw,err = mp.CreateFormField("recipient")

	enc := json.NewEncoder(fw)
	err = enc.Encode(&recipient)

	if err != nil {
		log.Println(err)
		return
	}

	fw,err = mp.CreateFormField("message")

	if err != nil {
		log.Println(err)
		return
	}

	enc = json.NewEncoder(fw)
	enc.Encode(&MessageImage{
		Attachment:MessageAttachment{
			Type:"image",
			Payload:map[string]string{},
		},
	})

	mp.Close()

	uri := "https://graph.facebook.com/v2.6/me/messages?access_token="+GlobalConfig.Token
	
	request,err := http.NewRequest("POST",uri,&buf)
	request.Header.Set("Content-Type",mp.FormDataContentType())

	if err != nil {
		log.Fatal(err)
	}
	
	client := &http.Client{}

	resp,err := client.Do(request)

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		log.Println(buf.String())
	}

}

func CreateImageFile(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "filedata", filename))
	h.Set("Content-Type", "image/png")
	return w.CreatePart(h)
}