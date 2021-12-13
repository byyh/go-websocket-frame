package types

type TestReq struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type TestResp struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
