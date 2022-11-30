package models

type Request struct{
	URL string `json:"url"`
	Expiry int `json:"expiry"`
}

type Response struct{
	URL string `json:"url"`
	Hash string `json:"hash"`
}