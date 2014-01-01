package sqwiggle

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type Client struct {
  Token string
}

/* func Get(target string, id uint) string { */
func (client* Client) Get(objectType string) string {
	net := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:3001/%s",objectType), nil)
	req.SetBasicAuth(client.Token, "X")
	resp, err := net.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
