package model

type JcApiKeyValuePairReq struct {
	Key  string      `json:"key" v:"required#Redis Set Key can not be empty"`
	Data interface{} `json:"data" v:""`
}

type JcApiHSetValueReq struct {
	Key   string      `json:"key" v:"required#Redis HSet Key can not be empty"`
	Field string      `json:"field" v:"required#Redis HSet Field can not be empty"`
	Data  interface{} `json:"data" v:""`
}

type JcApiOnlyKeyFieldReq struct {
	Key   string `json:"key" v:"required#Redis Key can not be empty"`
	Field string `json:"field" v:""`
}

type JcApiInfoFieldReq struct {
	Field string `json:"field" v:""`
}
