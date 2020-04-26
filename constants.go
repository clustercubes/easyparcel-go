package easyparcel

import (
  "encoding/json"
  models "github.com/helmiruza/easyparcel-go/models"
)

var ENVIRONMENT = ""
var APIKEY = ""
var AUTHKEY = ""
var URL = ""

const (
  ProductionUrl = "https://connect.easyparcel.my"
	StagingUrl = "https://demo.connect.easyparcel.my"
)

func Init(e string, f string, g string) {
	ENVIRONMENT = e
  APIKEY = f
  AUTHKEY = g

  if ENVIRONMENT == "production" {
    URL = ProductionUrl
  }

  if ENVIRONMENT == "staging" {
    URL = StagingUrl
  }
}

func objectFactory(objectType string) (interface{}) {
  switch (objectType) {
    case "creditBalance":
      return models.CreditBalanceResponse{}
    case "rateCheck":
      return models.BulkRateCheckResponse{}
    case "order":
      return models.BulkOrderResponse{}
    case "orderPayment":
      return models.BulkOrderPaymentResponse{}
  }
  return nil
}

func response(body []byte, err error, objectType string) (interface{}, models.Error) {
  errorMessage := models.Error{}
  response := objectFactory(objectType)

  if err != nil {
    errorMessage.Error = models.ErrorDetail { Type: "unknown", Message: "unknown" }
    return response, errorMessage
  }

  json.Unmarshal(body, &errorMessage)

  if len(errorMessage.Error.Type) > 0 {
    return response, errorMessage
  }

  json.Unmarshal(body, &response)

  return response, errorMessage
}
