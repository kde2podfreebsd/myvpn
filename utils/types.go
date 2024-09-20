package utils


type ResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      int    `json:"port"`
	Method    string `json:"method"`
	AccessUrl string `json:"accessUrl"`
}

type AccessKeyList struct {
	AccessKeys []ResponseData `json:"accessKeys"`
}

type Limit struct {
	Bytes int `json:"bytes"`
}

type CreateAccessKeyPayload struct {
	Method   string `json:"method"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Limit    Limit  `json:"limit"`
}