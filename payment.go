package easyparcel

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	models "github.com/clustercubes/easyparcel-go/models"
)

func SubmitBulkOrderPayment(data models.BulkOrderPayment) (models.BulkOrderPaymentResponse, models.Error) {
	URL += "/?ac=MPPayOrderBulk"
	requestBody, _ := json.Marshal(data)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-type", "application/json")

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	obj, err1 := response(body, err, "orderPayment")
	objString, _ := json.Marshal(obj)
	k := models.BulkOrderPaymentResponse{}
	json.Unmarshal(objString, &k)

	return k, err1
}
