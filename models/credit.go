package models

type CreditResponse struct {
	Currency string	`json:"currency"`
	Result float64	`json:"result"`
	ApiStatus string	`json:"api_status"`
	ErrorCode string	`json:"error_code"`
  ErrorRemark string	`json:"error_remark"`
}
