package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/moosilauke18/plaidApply/job"
	"github.com/sendgrid/sendgrid-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	// The api endpoint
	URL    = "https://contact.plaid.com/jobs"
	EMAIL  = "efarrell18@gmail.com"
	NAME   = "Evan Farrell"
	RESUME = "https://drive.google.com/file/d/0B1wn06WAk2BdYUdwM0laWDJjUE0/view?usp=sharing"
	GITHUB = "https://github.com/moosilauke18"
)

var (
	// Variables to be set from ENV
	sendgridAPI   string
	EMAIL_TO      string
	EMAIL_TO_NAME string
)

// Main func to send request
func main() {
	jobReq := job.New(NAME, EMAIL, RESUME, GITHUB)
	js, err := json.MarshalIndent(jobReq, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(js))
	if err != nil {
		log.Fatal("Please Try again")
	}
	req.Header.Set("User-Agent", "github.com:moosilauke18/plaidApply")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	sendEmail()

}
func sendEmail() {
	from := mail.NewEmail("Evan Farrell", "efarrell18@gmail.com")
	subject := "Software Engineer job"
	to := mail.NewEmail(EMAIL_TO_NAME, EMAIL_TO)
	messageBody := fmt.Sprintf("Hi %s,\n My name is Evan. I saw your post on `Who's Hiring` on Hacker News about jobs and have worked with Plaid in the past. \nI submitted a job via the api on your site, but figured that it would be worthwhile to submit a post here too. \n\nI look forward to hearing from you.\nAlso, to be a bit different I created a small application in Golang to submit the API request and send the email. You can check it out here: https://github.com/moosilauke18/plaidApply\n\nThank you,\n %s", EMAIL_TO_NAME, NAME)
	content := mail.NewContent("text/plain", messageBody)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(sendgridAPI, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.StatusCode)
	log.Println(response.Body)
	log.Println(response.Headers)
}
func init() {
	// Get ENVs
	sendgridAPI = os.Getenv("SENDGRID_API_KEY")
	EMAIL_TO_NAME = os.Getenv("EMAIL_TO_NAME")
	EMAIL_TO = os.Getenv("EMAIL_TO")
}
