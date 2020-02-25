package models

type BulkRateCheck struct {
	AuthKey string	`json:"authentication"`
	ApiKey string	`json:"api"`
	Data []RateCheck `json:"bulk"`
}

type RateCheck struct {
	PickupCode string	`json:"pick_code"`
	PickupState string	`json:"pick_state"`
	PickupCountry string	`json:"pick_country"`
	SendCode string	`json:"send_code"`
	SendState string	`json:"send_state"`
	SendCountry string	`json:"send_country"`
	Weight float64	`json:"weight"`
	Width float64	`json:"width"`
	Height float64	`json:"height"`
	CollectionDate string	`json:"date_coll"`
}

type BulkRateCheckResponse struct {
	ApiStatus string	`json:"api_status"`
	ErrorCode string	`json:"error_code"`
	ErrorRemark string	`json:"error_remark"`
	Result []RateCheckResponse	`json:"result"`
}

type RateCheckResponse struct {
	Status string	`json:"status"`
	Remarks string	`json:"remarks"`
	Rates []RateResponse	`json:"rates"`
}

type RateResponse struct {
	RateId string `json:"rate_id"`
	ServiceDetail string `json:"service_detail"`
	ServiceId string `json:"service_id"`
	ServiceType string `json:"service_type"`
	CourierId string `json:"courier_id"`
	CourierLogo string `json:"courier_logo"`
	PickupDate string `json:"pickup_date"`
	Delivery string `json:"delivery"`
	Price string `json:"price"`
	AddonPrice string `json:"addon_price"`
	ShipmentPrice string `json:"shipment_price"`
	RequireMinOrder int `json:"require_min_order"`
	ServiceName string `json:"service_name"`
	CourierName string `json:"courier_name"`
	DropoffPoint []DropoffPointResponse	`json:"dropoff_point"`
	PickupPoint []PickupPointResponse	`json:"pickup_point"`
}

type DropoffPointResponse struct {
	PointId string `json:"point_id"`
	PointName string `json:"point_name"`
	PointContact string `json:"point_contact"`
	PointAddr1 string `json:"point_addr1"`
	PointAddr2 string `json:"point_addr2"`
	PointAddr3 string `json:"point_addr3"`
	PointAddr4 string `json:"point_addr4"`
	PointPostcode string `json:"point_postcode"`
	PointCity string `json:"point_city"`
	PointState string `json:"point_state"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Price string `json:"price"`
}

type PickupPointResponse struct {
	PointId string `json:"point_id"`
	Company string `json:"company"`
	PointName string `json:"point_name"`
	PointContact string `json:"point_contact"`
	PointLat string `json:"point_lat"`
	PointLon string `json:"point_lon"`
	PointAddr1 string `json:"point_addr1"`
	PointAddr2 string `json:"point_addr2"`
	PointAddr3 string `json:"point_addr3"`
	PointAddr4 string `json:"point_addr4"`
	PointCity string `json:"point_city"`
	PointState string `json:"point_state"`
	PointPostcode string `json:"point_postcode"`
	Price string `json:"price"`
}
