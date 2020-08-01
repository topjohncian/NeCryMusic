package command

import (
	"encoding/json"
	"fmt"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/topjohncian/NeCryMusic/pkg/http"
)

type Sentence struct {
	UID  int    `json:"uid"`
	Type int    `json:"type"`
	From string `json:"from"`
	Text string `json:"text"`
}

func GetOne() string {
	res, err := http.HttpGet("http://api.heerdev.top/nemusic/random")
	if err != nil {
		cqp.AddLog(cqp.Info, "NeCryMusic", "请求语录时发生错误："+err.Error())
		return "服务器去火星了，等会儿再试试吧"
	}

	sentence := &Sentence{}
	err = json.Unmarshal(res, sentence)
	if err != nil {
		cqp.AddLog(cqp.Info, "NeCryMusic", "请求语录时发生错误："+err.Error())
	}
	return fmt.Sprintf("\"%s\" —— %s", sentence.Text, sentence.From)
}
