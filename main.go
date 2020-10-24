package main

import (
	"fmt"
	twilio "github.com/sfreiberg/gotwilio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type twilioSender struct {
	client *twilio.Twilio
	from   string
	to     string
}

func (s *twilioSender) Send(text string) error {
	_, ex, err := s.client.SendSMS(s.from, s.to, text, "", "")
	if ex != nil {
		return ex
	}
	if err != nil {
		return err
	}
	return nil
}

type handler struct {
	sender *twilioSender
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	problem := func(err error) {
		log.Printf("Error occured: %s", err)
		http.Error(w, "some problem", http.StatusInternalServerError)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		problem(err)
		return
	}
	if err := h.sender.Send(string(body)); err != nil {
		problem(err)
		return
	}
	fmt.Fprint(w, "ok")
}

func main() {
	sid := os.Getenv("TWILIO_SID")
	token := os.Getenv("TWILIO_TOKEN")
	fromBumber := os.Getenv("FROM_NUMBER")
	toNumber := os.Getenv("TO_NUMBER")
	if sid == "" || token == "" || fromBumber == "" || toNumber == "" {
		log.Fatal("Provide TWILIO_SID, TWILIO_TOKEN, FROM_NUMBER and TO_NUMBER env vars")
	}

	port := ":8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = fmt.Sprintf(":%s", portEnv)
	}

	twilio := twilio.NewTwilioClient(sid, token)
	sender := &twilioSender{
		client: twilio,
		from:   fromBumber,
		to:     toNumber,
	}

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, LogMiddleware(&handler{sender})))
}
