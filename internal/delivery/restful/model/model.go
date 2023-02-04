package model

type CreateTextModel struct {
	Text string `json:"text"`
}

type GetTextModel struct {
	Code string `json:"code"`
}

type ResponseModel struct {
	Text string `json:"text"`
}
