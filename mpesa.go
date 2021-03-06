package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func RequestAuthentication(c *Credetials) string {
	var AuthResp AuthResponse

	client := Client()

	req, err := http.NewRequest("GET", endpoint+"/oauth/v1/generate?grant_type=client_credentials", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(c.APP_KEY, c.APP_SECRET)
	req.Header.Add("cache-control", "no-cache")
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	DecodeResponseBody(response, &AuthResp)
	return AuthResp.Access_token

}

func StkPush(s *StkPushRequest, c *Credetials) (interface{}, error) {

	var i interface{}

	client := Client()

	encoded_data, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint+"/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(encoded_data))

	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+RequestAuthentication(c))
	req.Header.Set("cache-control", "no-cache")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	DecodeResponseBody(response, &i)
	return i, nil
}

func StkPushTransactionStatus(s *StkPushStatusRequest, c *Credetials) (interface{}, error) {

	var i interface{}

	client := Client()

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint+"/mpesa/stkpushquery/v1/query", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("authorization", "Bearer "+RequestAuthentication(c))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Timestamp", s.Timestamp)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	DecodeResponseBody(response, &i)
	return i, nil
}

func DecodeResponseBody(h *http.Response, i interface{}) error {

	body, err := ioutil.ReadAll(h.Body)
	if err != nil {
		return err
	}

	err1 := json.Unmarshal(body, &i)
	if err1 != nil {
		return err1
	}
	return nil
}

func GenerateEncodedPassword(shortcode, passkey string) map[string]string {
	mytime := time.Now()
	formattedTime := mytime.Format("20060102150405")
	msg := shortcode + passkey + formattedTime
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(msg))

	return map[string]string{
		"encoded":       encodedPassword,
		"formattedTime": formattedTime,
	}
}

func Client() http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return *client
}
