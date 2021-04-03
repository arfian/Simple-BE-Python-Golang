package transformers

type Register struct {
	Phone		string `json:"phone"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}

type ResData struct {
	Data interface{} `json:"data"`
	ErrorMessage string `json:"error-message"`
	IsError bool `json:"is-error"`
}