package gorestcoin

type GorestError struct {
	Code int64               `json:"code"`
	Data []GorestFieldsError `json:"data"`
}

type GorestFieldsError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
