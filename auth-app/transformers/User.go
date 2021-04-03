package transformers

type Register struct {
	Phone		string `json:"phone"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}

type Login struct {
	Phone string `json:"phone"`
	Name string `json:"name"`
	Username string `json:"username"`
	Role string `json:"role"`
	Timestamp string `json:"timestamp"`
	BarerToken string `json:"barer-token-jwt"`
}

type ResData struct {
	Data interface{} `json:"data"`
	ErrorMessage string `json:"error-message"`
	IsError bool `json:"is-error"`
}