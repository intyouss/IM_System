package models

type Message struct {
	Code int         `json:"code" default:"0"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
