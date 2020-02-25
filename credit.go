package easyparcel

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/easyparcel-go/models"
)

func GetCreditBalance() (models.CreditResponse, models.Error) {
  URL += "/?ac=EPCheckCreditBalance"
  requestBody, _ := json.Marshal({"api": APIKEY})

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "credit")
  objString, _ := json.Marshal(obj)
  k := models.CreditResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
