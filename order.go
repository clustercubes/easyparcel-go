package easyparcel

import (
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/easyparcel-go/models"
)

func SubmitBulkOrder(data models.BulkOrder) (models.BulkOrderResponse, models.Error) {
  URL += "/?ac=MPSubmitOrderBulk"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "order")
  objString, _ := json.Marshal(obj)
  k := models.BulkOrderResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
