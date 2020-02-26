package models

type BulkOrderPayment struct {
	AuthKey string	`json:"authentication"`
	ApiKey string	`json:"api"`
	Data []Order `json:"bulk"`
}

type OrderPayment struct {
	OrderNo string	`json:"order_no"`
}

type BulkOrderPaymentResponse struct {
	ApiStatus string	`json:"api_status"`
	ErrorCode string	`json:"error_code"`
	ErrorRemark string	`json:"error_remark"`
  Result []OrderPaymentResponse `json:"result"`
}

type OrderPaymentResponse struct {
  OrderNo string	`json:"orderno"`
  Message string	`json:"messagenow"`
  Parcel []ParcelDetail `json:"parcel"`
}

type ParcelDetail struct {
  ParcelNo string `json:"parcelno"`
  Awb string `json:"awb"`
  AwbLink string `json:"awb_id_link"`
  TrackingUrl string `json:"tracking_url"`
}
