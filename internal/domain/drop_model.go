package domain

type CreateText struct {
	Text string `json:"text"`
}

type GetText struct {
	Code string `json:"code"`
}
