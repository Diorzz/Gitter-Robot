package main

import (
	"encoding/json"
)

type QueryInfo struct {
	ReqType  int         `json:"reqType"`
	Per      *Perception `json:"perception"`
	UserInfo *UserInfo   `json:"userInfo"`
}

type Perception struct {
	InputText *InputText `json:"inputText"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

type InputText struct {
	Text string `json:"text"`
}

type TuringRoboat struct {
	postMsg PostMsg
}

type PosMsgImg struct {
	baseUrl string
}

func (pmi *PosMsgImg) GetUrl() string {
	return pmi.baseUrl
}

func (pmi *PosMsgImg) GetHeader() string {
	return ""
}

func (pmi *PosMsgImg) GetBody(qes string) string {
	inputText := &InputText{
		Text: qes,
	}
	per := &Perception{
		InputText: inputText,
	}

	uInfo := &UserInfo{
		ApiKey: "",
		UserId: "",
	}

	qInfo := &QueryInfo{
		ReqType:  0,
		Per:      per,
		UserInfo: uInfo,
	}

	b, err := json.Marshal(qInfo)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (tr *TuringRoboat) Ask(question string) (string, error) {
	url := tr.postMsg.GetUrl()
	header := tr.postMsg.GetHeader()
	body := tr.postMsg.GetBody(question)

	res, err := httpPost(url, header, body)
	if err != nil {
		return "", err
	}

	return res, err
}

func (tu *TuringRoboat) AskQingYun(question string) (string, error) {

	res, err := httpGet("http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + question)
	if err != nil {
		return "", err
	}

	return res, err
}

func NewTuringRoboat() *TuringRoboat {
	pmi := &PosMsgImg{
		baseUrl: "http://openapi.tuling123.com/openapi/api/v2",
	}
	tu := &TuringRoboat{
		postMsg: pmi,
	}
	return tu
}

type Result struct {
	Result []*Chat `json:"results"`
}

type Chat struct {
	GroupType  int    `json:"groupType"`
	ResultType string `json:"resultType"`
	Values     *Value `json:"values"`
}

type Value struct {
	Text string `json:"text"`
}

type QingYunResult struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}
