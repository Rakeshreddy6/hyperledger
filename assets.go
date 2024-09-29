package models

type Asset struct {
	DEALERID  string `json:"dealerId"`
	MSISDN    string `json:"msisdn"`
	MPIN      string `json:"mpin"`
	BALANCE   int    `json:"balance"`
	STATUS    string `json:"status"`
	TRANSAMOUNT int    `json:"transAmount"`
	TRANSTYPE  string `json:"transType"`
	REMARKS    string `json:"remarks"`
}