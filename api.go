package sqwiggle

import (
  "fmt"
  "net/http"
  "net/url"
  "log"
  "io/ioutil"
)

const (
  baseUrl string = "api.sqwiggle.com/"
)

type Client struct {
  Token string
}

type ErrorResponse struct {
  Status int
  Kind string
  Message string
  Details string
}

type User struct {
  Email string
  Name string
}

func (client* Client) get(path string) string {
  net := &http.Client{}
  u := new(url.URL)
  u.Scheme = "https"
  u.Host = baseUrl
  u.Path = path
  fmt.Printf("asfdaf %v", u.String())
  req, err := http.NewRequest("GET", u.String(), nil)
  req.SetBasicAuth(client.Token, "X")
  resp, err := net.Do(req)
  if err != nil{
    log.Fatal(err)
  }
  bodyText, err := ioutil.ReadAll(resp.Body)
  s := string(bodyText)
  return s
}

func (client *Client) Users() []User {
  result := []User{}
  resp := client.get("users")
  fmt.Printf("resp %v", resp)
  result = append(result, User { "1", "2" })
  return result
}
