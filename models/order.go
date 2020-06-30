package models

type BulkOrder struct {
	AuthKey string  `json:"authentication"`
	ApiKey  string  `json:"api"`
	Data    []Order `json:"bulk"`
}

type Order struct {
	Weight         float64 `json:"weight"`
	Content        string  `json:"content"`
	Value          float64 `json:"value"`
	ServiceId      string  `json:"service_id"`
	PickupName     string  `json:"pick_name"`
	PickupCompany  string  `json:"pick_company"`
	PickupContact  string  `json:"pick_contact"`
	PickupMobile   string  `json:"pick_mobile"`
	PickupAddr1    string  `json:"pick_addr1"`
	PickupAddr2    string  `json:"pick_addr2"`
	PickupCity     string  `json:"pick_city"`
	PickupState    string  `json:"pick_state"`
	PickupCode     string  `json:"pick_code"`
	PickupCountry  string  `json:"pick_country"`
	SendName       string  `json:"send_name"`
	SendContact    string  `json:"send_contact"`
	SendMobile     string  `json:"send_mobile"`
	SendAddr1      string  `json:"send_addr1"`
	SendAddr2      string  `json:"send_addr2"`
	SendCity       string  `json:"send_city"`
	SendState      string  `json:"send_state"`
	SendCode       string  `json:"send_code"`
	SendCountry    string  `json:"send_country"`
	CollectionDate string  `json:"collect_date"`
	Sms            bool    `json:"sms"`
	SendEmail      string  `json:"send_email"`
}

type BulkOrderResponse struct {
	Result      []OrderResponse `json:"result"`
	ApiStatus   string          `json:"api_status"`
	ErrorCode   string          `json:"error_code"`
	ErrorRemark string          `json:"error_remark"`
}

type OrderResponse struct {
	ParcelNumber string `json:"parcel_number"`
	OrderNumber  string `json:"order_number"`
	Price        string `json:"price"`
	Status       string `json:"status"`
	Remarks      string `json:"remarks"`
	Courier      string `json:"courier"`
}

type BulkOrderStatus struct {
	ApiKey string        `json:"api"`
	Data   []OrderStatus `json:"bulk"`
}

type OrderStatus struct {
	OrderNo string `json:"order_no"`
}

type BulkOrderStatusResponse struct {
	Result      []OrderStatusResponse `json:"result"`
	ApiStatus   string                `json:"api_status"`
	ErrorCode   string                `json:"error_code"`
	ErrorRemark string                `json:"error_remark"`
}

type OrderStatusResponse struct {
	Status       string `json:"status"`
	Remarks      string `json:"remarks"`
	OrderNo      string `json:"order_no"`
	OrderStatus  string `json:"order_status"`
	OrderPayable bool   `json:"order_payable"`
}
