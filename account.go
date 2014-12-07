package dropletkit

import (
  "io/ioutil"
  "encoding/json"
  "net/http"
)

type Account struct {
  Uuid string `json:"uuid"`
  Email string `json:"email"`
  EmailVerified bool `json:"email_verified"`
  DropletLimit int `json:"droplet_limit"`
}

type AccountWrapper struct {
  Account Account
}

func AccountInfo(options Options) Account {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", options.baseUrl + "/" + options.version + "/account", nil)
  req.Header.Add("Authorization", "Bearer " + options.Token)
  response, _ := client.Do(req)

  body, _ := ioutil.ReadAll(response.Body)
  decoded := AccountWrapper{}
  json.Unmarshal(body, &decoded)

  return decoded.Account
}
